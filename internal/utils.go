package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const (
	ansiFg = "\033[38;5;%dm%s\033[0m"
)

var (
	re_explicit *regexp.Regexp
	idx_r       int
	idx_g       int
	idx_b       int
	idx_re      int
	darkColors  []int
	lightColors []int
)

// Compute the DJB2 hash of a byte array
// See: http://www.cse.yorku.ca/~oz/hash.html
func djb2Hash(s []byte) uint64 {
	var hash uint64 = 5381

	for _, ch := range s {
		hash = hash + (hash << 5) + uint64(ch) // hash*33 + uint64(ch)
	}
	return hash
}

// Convert RGB values to an int for use as ANSI escape code
func rgbCode(r, g, b int) int {
	return (r * 36) + (g * 6) + b + 16
}

// Write a string to a *File
func writeStr(out *os.File, s string) {
	if _, err := out.WriteString(s); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Wrap regexp matches with ansi shell escape codes
func wrapColor(re *regexp.Regexp, line []byte, colors []int, explicit map[string]int) []byte {
	// thank god for the closures!
	replacer := func(token []byte) []byte {
		var buf bytes.Buffer
		var color int
		if c, ok := explicit[string(token)]; ok {
			// if an explicit color is provided, use that
			color = c
		} else {
			cid := djb2Hash(token) % uint64(len(colors))
			color = colors[cid]
		}
		fmt.Fprintf(&buf, ansiFg, color, token)
		return buf.Bytes()
	}
	return re.ReplaceAllFunc(line, replacer)
}

// Make ANSI colors
func makeColors(pred func(int, int, int) bool) []int {
	cs := make([]int, 0)

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 6; k++ {
				if pred(i, j, k) {
					rc := rgbCode(i, j, k)
					cs = append(cs, rc)
				} else {
					continue
				}
			}
		}
	}
	return cs
}

// Parse explicit colors
func ParseExplicit(opt []string) (data map[string]int) {
	data = make(map[string]int)
	for _, o := range opt {
		g := re_explicit.FindSubmatch([]byte(o))
		if g != nil {
			mr, _ := strconv.Atoi(string(g[idx_r]))
			mg, _ := strconv.Atoi(string(g[idx_g]))
			mb, _ := strconv.Atoi(string(g[idx_b]))
			data[string(g[idx_re])] = rgbCode(mr, mg, mb)
		}
	}
	// return nil if we have no matches
	if len(data) > 0 {
		return data
	} else {
		return nil
	}
}

// The core engine that does all the work
func Highlight(re *regexp.Regexp, rdr *bufio.Reader, out *os.File, explicit map[string]int, isLight bool) {
	var colors []int
	if isLight {
		colors = lightColors
	} else {
		colors = darkColors
	}

	for {
		switch line, err := rdr.ReadString('\n'); err {
		case nil:
			writeStr(out, string(wrapColor(re, []byte(line), colors, explicit)))

		// exit gracefully if we have EOF
		case io.EOF:
			os.Exit(0)

		// unhandled error
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
}

func init() {
	re_explicit = regexp.MustCompile(`(?P<r>[0-5]),(?P<g>[0-5]),(?P<b>[0-5]):(?P<re>.+)`)
	idx_r = re_explicit.SubexpIndex("r")
	idx_g = re_explicit.SubexpIndex("g")
	idx_b = re_explicit.SubexpIndex("b")
	idx_re = re_explicit.SubexpIndex("re")

	darkColors = makeColors(func(i, j, k int) bool {
		return (i + j + k) > 3
	})

	lightColors = makeColors(func(i, j, k int) bool {
		return (i + j + k) < 11
	})
}
