// xpo - A CLI tool to highlight regexp matches
package main

import (
	"bufio"
	"os"
	"regexp"

	"fmt"

	i "github.com/ghoseb/xpo/internal"
	flag "github.com/spf13/pflag"
)

func main() {
	var regex string
	var explicit []string
	var isLight bool
	var ex map[string]int

	fs := flag.NewFlagSet("xpo", flag.ContinueOnError)
	fs.StringVarP(&regex, "regex", "r", "", "regular expression to highlight")
	fs.StringArrayVarP(&explicit, "explicit", "e", []string{}, "explicit color for matches")
	fs.BoolVarP(&isLight, "light-color", "l", false, "use light colors")

	// fs.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of xpo:\n")

	// 	fmt.Fprintf(os.Stderr, "%s", fs.FlagUsagesWrapped(80))
	// }

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
		fmt.Fprintf(os.Stderr, "Need a regexp to search\n")
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
