package funcs

import "math"

type mode byte

const (
	SAMPLE mode = iota
	POPULATION
)

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

func Variance(data []float64, m mode) float64 {
	if m == POPULATION {
		return varianceSum(data) / float64(len(data))
	} else if m == SAMPLE {
		if len(data) == 0 {
			return 0
		}
		return varianceSum(data) / float64(len(data)-1)
	}
	panic("Not a valid mode")
}

func Stdev(data []float64, m mode) float64 {
	return math.Sqrt(Variance(data, m))
}

func ZScore(value float64, data []float64, m mode) float64 {
	return (value - Mean(data)) / Stdev(data, m)
}
