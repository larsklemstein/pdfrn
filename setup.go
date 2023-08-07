package main

import (
	"fmt"
	"os"
)

func printUsageAndExit() {
	progname := "pdfrn"
	rc := 2

	usage_header := "\npdfrn == pdfrename\n\n" +
		"Simple command to let you rename a bunch of PDF files\n\n" +
		"Usage:\n"

	author_line := "bugs and hints: lrsklemstein@gmail.com\n"

	usage := fmt.Sprintf("  %s get_fields PATTERN\n", progname) +
		fmt.Sprintf("  %s rename CFG\n\n", progname)

	complete_usage := usage_header + usage + author_line

	fmt.Fprint(os.Stderr, "->"+complete_usage+"<-")
	os.Exit(rc)
}
