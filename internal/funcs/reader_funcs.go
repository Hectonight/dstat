package funcs

import (
	"fmt"
	"os"
	"strconv"
)

//func ScanFile(scn *bufio.Scanner) []float64 {
//	return make([]float64, 1)
//}

func NoData() {
	_, err := fmt.Fprintln(os.Stderr, "no data given")
	if err != nil {
		panic(err)
	}
}

func ScanLineData(fields []string) []float64 {
	data := make([]float64, len(fields))
	for i, field := range fields {
		value, err := strconv.ParseFloat(field, 64)
		if err != nil {
			panic(err)
		}
		data[i] = value
	}
	return data
}
