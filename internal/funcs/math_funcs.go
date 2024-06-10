package funcs

import (
	"errors"
	"math"
)

const SAMPLE = 0
const POPULATION = 1

func Sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func Mean(data []float64) float64 {
	return Sum(data) / float64(len(data))
}

func Median(data []float64) float64 {
	length := len(data)

	if length%2 == 1 {
		return data[length/2]
	} else {
		mid := length / 2
		return (data[mid] + data[mid/2-1]) / 2.0
	}
}

func varianceSum(data []float64) float64 {
	sum := 0.0
	avg := Mean(data)
	for _, v := range data {
		sum += math.Pow(v-avg, 2)
	}

	return sum
}

func Variance(data []float64, mode int) (float64, error) {
	if mode == POPULATION {
		return variancePopulation(data), nil
	} else if mode == SAMPLE {
		return varianceSample(data), nil
	}

	return 0, errors.New("invalid mode")
}

func varianceSample(data []float64) float64 {
	if len(data) == 1 {
		return 0
	}
	return varianceSum(data) / float64(len(data)-1)
}

func variancePopulation(data []float64) float64 {
	return varianceSum(data) / float64(len(data))
}

func StandardDeviation(data []float64, mode int) (float64, error) {
	v, err := Variance(data, mode)
	return math.Sqrt(v), err
}

func ZScore(value float64, data []float64, mode int) (float64, error) {
	stdev, err := StandardDeviation(data, mode)
	return (value - Mean(data)) / stdev, err
}
