package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	Numbers []float64
	Mean    float64
	Median  float64
	Mode    []float64
	StdDev  float64
}

type page struct {
	Title   string
	Stats   statistics
	Message string
}

var templates = template.Must(template.ParseGlob("*.html"))

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("http server failed to start", err)
	}
}

func index(response http.ResponseWriter, request *http.Request) {
	log.Println(*request)

	if err := request.ParseForm(); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var page page
	page.Title = "Chapter 2.4 - Statistics"

	if numbers, msg, ok := parseRequest(request); ok {
		page.Stats = getStatistics(numbers)
	} else {
		page.Message = msg
	}

	renderTemplate(response, "index.html", &page)
}

func parseRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if data, found := request.Form["numbers"]; found && len(data) > 0 {
		text := strings.Replace(data[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if n, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is an invalid number", false
			} else {
				numbers = append(numbers, n)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func renderTemplate(response http.ResponseWriter, tmpl string, page *page) {
	if err := templates.ExecuteTemplate(response, tmpl, page); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func getStatistics(numbers []float64) (stats statistics) {
	stats.Numbers = numbers
	sort.Float64s(stats.Numbers)
	stats.Mean = stats.calculateMean()
	stats.Median = stats.calculateMedian()
	stats.Mode = stats.calculateMode()
	stats.StdDev = stats.calculateStdDev()
	return stats
}

func (stats *statistics) SumOfNumbers() (sum float64) {
	for _, number := range stats.Numbers {
		sum += number
	}
	return sum
}

func (stats *statistics) NumberListing() (output string) {
	for index, number := range stats.Numbers {
		if index != 0 {
			output += ", "
		}
		output += fmt.Sprintf("%v", number)
	}
	if output == "" {
		output = "-"
	}
	return output
}

func (stats *statistics) LenOfNumbers() float64 {
	return float64(len(stats.Numbers))
}

func (stats *statistics) calculateMean() float64 {
	return stats.SumOfNumbers() / stats.LenOfNumbers()
}

// expects numbers to be sorted
func (stats *statistics) calculateMedian() float64 {
	length := len(stats.Numbers)
	middle := length / 2
	result := stats.Numbers[middle]

	if length%2 == 0 {
		result = (result + stats.Numbers[middle-1]) / 2
	}

	return result
}

func (stats *statistics) calculateMode() (modes []float64) {
	frequencies := make(map[float64]int, len(stats.Numbers))
	highestFrequency := 0
	for _, x := range stats.Numbers {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(frequencies) {
		modes = modes[:0]
	}
	sort.Float64s(modes)
	return modes
}

func (stats *statistics) calculateStdDev() float64 {
	total := 0.0
	for _, number := range stats.Numbers {
		total += math.Pow(number-stats.Mean, 2)
	}
	variance := total / float64(len(stats.Numbers)-1)
	return math.Sqrt(variance)
}
