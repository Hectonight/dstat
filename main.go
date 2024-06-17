package main

import (
	"dstat/internal/funcs"
	"dstat/internal/platform"
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
	--json
*/

/*
	Add error messages for:

	strconv.ParseFloat: parsing "...": invalid syntax
	strconv.ParseFloat: parsing "...": value out of range

	+Inf, -Inf, and NaN are valid values

	Values my hit +Inf and -Inf

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
	var whitespace *string = flag.StringP("whitespace", "w", "",
		"Treat characters in this string as whitespace.")
	var ignoreFlag *string = flag.StringP("ignore", "i", "",
		"Ignore characters in this string.")
	var sum *bool = flag.Bool("sum", false, "Find the sum of the data")

	//var summary *bool = flag.BoolP("summary", "s", false,
	//	"Give a 5 number summary of the data")

	flag.Parse()

	if flag.NArg()|flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	separators := []rune(*whitespace)
	ignore := []rune(*ignoreFlag)

	var data []float64
	var err error

	if flag.NArg() == 0 {
		fmt.Printf("Input Data (%v to end):\n", platform.EOFKey)
		data, err = funcs.ReadFile(os.Stdin, separators, ignore)
	} else {
		data, err = funcs.ReadFiles(flag.Args(), separators, ignore)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "dstat: %s\n", err.Error())
		os.Exit(1)
	}

	slices.Sort(data)

	if *count {
		fmt.Printf("Count: %v\n", len(data))
	}

	if *sum {
		fmt.Printf("Sum: %v\n", funcs.Sum(data))
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
