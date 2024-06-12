package funcs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

const capacity = 30

func NoData() {
	_, err := fmt.Fprintln(os.Stderr, "no data given")
	if err != nil {
		panic(err)
	}
}

func ConvertFloats(fields []string) ([]float64, error) {
	data := make([]float64, len(fields))
	for i, field := range fields {
		value, err := strconv.ParseFloat(field, 64)

		if err != nil {
			return nil, err
		}
		data[i] = value
	}
	return data, nil
}

func ReadFile(r io.Reader, seperators []rune, ignore []rune) ([]float64, error) {
	data := make([]float64, 0, capacity)
	scanner := bufio.NewScanner(r)

	// Maybe come back and improve with scanner.Split
	for scanner.Scan() {
		text := strings.Map(func(r rune) rune {
			switch {
			case slices.Contains(ignore, r):
				return -1
			case slices.Contains(seperators, r):
				return ' '
			}
			return r
		}, scanner.Text())
		dat, err := ConvertFloats(strings.Fields(text))
		if err != nil {
			return nil, err
		}
		data = slices.Concat(data, dat)
	}
	return data, nil
}

func ReadFiles(files []string, seperators []rune, ignore []rune) ([]float64, error) {
	data := make([]float64, 0, capacity*len(files))

	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}

		dat, err := ReadFile(file, seperators, ignore)
		if err != nil {
			return nil, err
		}

		data = slices.Concat(data, dat)

		file.Close()
	}

	return data, nil
}
