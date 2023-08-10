// xpo - A CLI tool to highlight regexp matches
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	i "github.com/ghoseb/xpo/internal"

	flag "github.com/spf13/pflag"
)

func main() {
	var regex string
	var explicit []string
	var isLight bool
	var ex map[string]int

	fs := flag.NewFlagSet("xpo", flag.ContinueOnError)
	fs.StringVarP(&regex, "regex", "r", "", "[required] regular expression to highlight")
	fs.StringArrayVarP(&explicit, "explicit", "e", []string{}, "[optional] explicit color for matches")
	fs.BoolVarP(&isLight, "light-color", "l", false, "[optional] use colors suitable for a light background")

	// custom help printer
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of xpo\n------------\n")
		fmt.Fprintf(os.Stderr, "%s\n", fs.FlagUsagesWrapped(115))
		fmt.Fprintf(os.Stderr, "Check out github.com/ghoseb/xpo/ for examples.")
	}

	switch err := fs.Parse(os.Args[1:]); err {
	case nil:
		break
	case flag.ErrHelp:
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "error in parsing flags: %v\n", err)
		os.Exit(1)
	}

	if regex == "" {
		fmt.Fprintf(os.Stderr, "Please supply a regexp using the -r flag. Use -h for help.")
		os.Exit(1)
	}

	re := regexp.MustCompile(regex)

	if len(explicit) != 0 {
		ex = i.ParseExplicit(explicit)
	}

	rdr := bufio.NewReader(os.Stdin)
	out := os.Stdout

	i.Highlight(re, rdr, out, ex, isLight)
}
