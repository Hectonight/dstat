package main

import (
	"dstat/internal/funcs"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"slices"
)

/*
	Flag ideas

	-r --round [int]
	-q --quartile [1-3]
	-s --summary
	-sc --scientific
	-v --version
*/

/*
	Add error messages for:

	strconv.ParseFloat: parsing "l": invalid syntax
	open TEST: no such file or directory

	On Windows
	open testdata/no_access.txt: Access is denied.

	On WSL
	open testdata/no_access.txt: permission denied

*/

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s: %s [OPTIONS]... [FILES]...\nIf FILES is not set read from standard input.\n",
			os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var help *bool = flag.BoolP("help", "h", false, "Display this message.")
	var mean *bool = flag.Bool("mean", false, "Find the mean of the data.")
	var minFlag *bool = flag.Bool("min", false, "Find the minimum of the data.")
	var maxFlag *bool = flag.Bool("max", false, "Find the maximum of the data.")
	var count *bool = flag.BoolP("count", "n", false, "Size of the data set.")
	var median *bool = flag.Bool("median", false, "Find the median of the data.")
	var stdev *bool = flag.Bool("stdev", false,
		"Find the standard deviation of the data while treating it as a sample.")
	var variance *bool = flag.Bool("var", false,
		"Find the variance of the data while treating it as a sample.")
	var stdevp *bool = flag.Bool("stdevp", false,
		"Find the standard deviation of the data while treating it as a population.")
	var variancep *bool = flag.Bool("varp", false,
		"Find the variance of the data while treating it as a population.")

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

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	slices.Sort(data)

	if *count {
		fmt.Printf("Count: %v\n", len(data))
	}

	if *mean {
		fmt.Printf("Mean: %v\n", funcs.Mean(data))
	}

	if *minFlag {
		fmt.Printf("Min: %v\n", data[0])
	}

	if *median {
		fmt.Printf("Median: %v\n", funcs.Median(data))
	}

	if *maxFlag {
		fmt.Printf("Max: %v\n", data[len(data)-1])
	}

	if *variance {
		fmt.Printf("Stdev: %v\n", funcs.Variance(data, funcs.SAMPLE))
	}

	if *stdev {
		fmt.Printf("Stdev: %v\n", funcs.Stdev(data, funcs.SAMPLE))
	}

	if *variancep {
		fmt.Printf("Stdev: %v\n", funcs.Variance(data, funcs.POPULATION))
	}

	if *stdevp {
		fmt.Printf("Stdev: %v\n", funcs.Stdev(data, funcs.SAMPLE))
	}

}
