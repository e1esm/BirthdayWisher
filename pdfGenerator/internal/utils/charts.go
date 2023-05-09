package utils

import (
	"bytes"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"os/exec"
	"pdfGenerator/internal/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

var arr []models.User

func GenerateAllImages(chatID int64, users []models.User, wg *sync.WaitGroup) {
	arr = users

	innerWg := &sync.WaitGroup{}
	innerWg.Add(3)
	go createYearsPieChart(chatID, innerWg)
	go createMonthsPieChart(chatID, innerWg)
	go createAgesPieChart(chatID, innerWg)
	innerWg.Wait()
	wg.Done()

}

func generateFiles(chart *charts.Pie, chatId int64, pieType string, wg *sync.WaitGroup) {
	nameHtml := fmt.Sprintf("pie-%s-%d.html", pieType, chatId)
	namePng := fmt.Sprintf("pie-%s-%d.png", pieType, chatId)
	buf := new(bytes.Buffer)
	_ = chart.Render(buf)

	strContent := buf.String()
	strContent = strings.ReplaceAll(strContent, "let", "")
	strContent = strings.ReplaceAll(strContent, "\"use strict\";", "")
	f, _ := os.Create(nameHtml)
	_, err := f.WriteString(strContent)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("wkhtmltoimage", nameHtml, namePng)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func generateYearsPieData() []opts.PieData {
	chartPeople := make([]opts.PieData, 0, 10)
	yearsCount := countYears()
	for k, v := range yearsCount {
		chartPeople = append(chartPeople, opts.PieData{
			Name:  k,
			Value: v,
		})
	}

	return chartPeople
}

func generateMonthsPieData() []opts.PieData {
	chartPeople := make([]opts.PieData, 0, 10)
	monthsCount := countMonths()
	for k, v := range monthsCount {
		chartPeople = append(chartPeople, opts.PieData{
			Name:  k.String(),
			Value: v,
		})
	}
	return chartPeople
}

func generateAgesPieData() []opts.PieData {
	chartPeople := make([]opts.PieData, 0, 10)
	agesCount := countAges()
	for k, v := range agesCount {
		chartPeople = append(chartPeople, opts.PieData{
			Name:  k,
			Value: v,
		})
	}

	return chartPeople
}

func createYearsPieChart(chatID int64, wg *sync.WaitGroup) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    fmt.Sprintf("Chat %d years pie", chatID),
				Subtitle: time.Now().Format(time.DateOnly)},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Years", generateYearsPieData()).SetSeriesOptions(
		charts.WithPieChartOpts(
			opts.PieChart{
				Radius: 200,
			},
		),
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b} : {c}",
			},
		),
	)
	generateFiles(pie, chatID, "years", wg)
}

func createMonthsPieChart(chatID int64, wg *sync.WaitGroup) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    fmt.Sprintf("Chat %d months pie", chatID),
				Subtitle: time.Now().Format(time.DateOnly),
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Months", generateMonthsPieData()).SetSeriesOptions(
		charts.WithPieChartOpts(
			opts.PieChart{
				Radius: 200,
			},
		),
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b} - {c}",
			},
		),
	)

	generateFiles(pie, chatID, "months", wg)
}

func createAgesPieChart(chatID int64, wg *sync.WaitGroup) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    fmt.Sprintf("Chat %d ages chart", chatID),
				Subtitle: time.Now().Format(time.DateOnly),
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Ages", generateAgesPieData()).SetSeriesOptions(
		charts.WithPieChartOpts(
			opts.PieChart{
				Radius: 200,
			},
		),
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b} - {c}",
			},
		),
	)
	generateFiles(pie, chatID, "ages", wg)
}

func getCategory(age int) string {
	var category string
	switch {
	case age >= 10 && age < 20:
		category = "10-19"
	case age >= 20 && age < 30:
		category = "20-29"
	case age >= 30 && age < 40:
		category = "30-39"
	case age >= 40 && age < 60:
		category = "40-59"
	case age >= 60 && age < 80:
		category = "60-79"
	case age >= 80 && age < 100:
		category = "80-100"
	case age >= 100:
		category = "100-∞"
	}
	return category
}

func countAges() map[string]int {
	agesCount := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		temp := arr[i].Date[:4]
		intTemp, err := strconv.Atoi(temp)
		if err != nil {
			panic(err)
		}
		year, _, _ := time.Now().Date()
		age := year - intTemp
		category := getCategory(age)

		v, ok := agesCount[category]
		if ok {
			agesCount[category] = v + 1
		} else {
			agesCount[category] = 1
		}

	}
	return agesCount
}

func countYears() map[string]int {
	yearsCount := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		temp := arr[i].Date[:4]
		v, ok := yearsCount[temp]
		if ok {
			yearsCount[temp] = v + 1
		} else {
			yearsCount[temp] = 1
		}
	}
	return yearsCount
}

func countMonths() map[time.Month]int {
	monthsCount := make(map[time.Month]int)
	for i := 0; i < len(arr); i++ {
		temp := arr[i].Date[5:7]
		num, err := strconv.Atoi(temp)
		if err != nil {
			panic(err)
		}
		v, ok := monthsCount[time.Month(num)]
		if ok {
			monthsCount[time.Month(num)] = v + 1
		} else {
			monthsCount[time.Month(num)] = 1
		}
	}
	return monthsCount
}
