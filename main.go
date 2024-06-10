package main

import (
	"bufio"
	"dstat/internal/funcs"
	"dstat/internal/help"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	mode := funcs.POPULATION
	var data []float64

Loop:
	for {
		fmt.Print("dstat> ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		fields := strings.Fields(input)

		if len(fields) == 0 {
			continue
		}

		switch fields[0] {

		case "data":
			tmp := funcs.ScanLineData(fields[1:])

			if len(tmp) == 0 {
				fmt.Println(help.Data)
				continue
			}
			data = tmp
			slices.Sort(data)

		case "avg", "mean":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(funcs.Mean(data))

		case "med", "median":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(funcs.Median(data))

		case "min":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(data[0])

		case "max":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(data[len(data)-1])

		case "sort":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			for _, num := range data {
				fmt.Print(num, " ")
			}
			fmt.Print("\n")

		case "population":
			mode = funcs.POPULATION

		case "sample":
			mode = funcs.SAMPLE

		case "variance":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(funcs.Variance(data, mode))

		case "stdev":
			if len(data) == 0 {
				funcs.NoData()
				continue
			}
			fmt.Println(funcs.StandardDeviation(data, mode))

		//case "zscore":
		//	if len(data) == 0 {
		//		funcs.NoData()
		//		continue
		//	}
		//	//fmt.Println(funcs.ZScore(data, mode))

		case "help":
			fmt.Println("NOT IMPLEMENTED")

		case "exit":
			break Loop

		default:
			_, err := fmt.Fprintf(os.Stderr, "no such command: %s\n", fields[0])
			if err != nil {
				panic(err)
			}
		}
	}
}
