package main

import (
	"dstat/internal/funcs"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

/*
	Flag ideas

	-r --round [int]
	--mean
	--stdev
	-var --variance
	--max
	--min
	-q --quartile [1-3]
	-med --median
	--stdevp
	-varp --variance-population
	--mode
	-i --ignore string
	-n --count
	-s --summary
	-sc --scientific
	-v --version
*/

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s: %s [OPTIONS]... [FILES]...\nIf FILES is not set read from standard input.\n",
			os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	// TODO: Add help messages
	var help *bool = flag.BoolP("help", "h", false, "Display this message")

	flag.Parse()

	if flag.NArg()|flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	separators := []rune{}
	ignore := []rune{}

	var data []float64
	var err error
	if flag.NArg() == 0 {

	} else {
		data, err = funcs.ReadFiles(flag.Args(), separators, ignore)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	fmt.Println(funcs.Mean(data))

}
