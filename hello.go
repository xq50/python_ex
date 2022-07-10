package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func formatBigNum(s string) string {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	exp := int(math.Log10(num) / 3)
	base := num / math.Pow10(3*exp)
	suffix := ""

	switch exp {
	case 2:
		suffix = "M"
	case 3:
		suffix = "B"
	case 4:
		suffix = "T"
	}
	return fmt.Sprintf("%.2f%s", base, suffix)
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	ass, err := GetAssets(10)
	if err != nil {
		log.Fatal(err)
	}

	headers := []string{"Name", "Price", "Market Cap", "24h %"}
	rows := [][]string{headers}
	for _, a := range ass.Data {
		rows = append(rows, []string{
			fmt.Sprintf("%s (%s)", a.Name, a.Symbol),
			a.PriceUsd[:7],
			formatBigNum(a.MarketCapUsd),
			a.ChangePercent24Hr[:strings.Index(a.ChangePercent24Hr, ".")+3] + "%",
		})
	}

	table3 := widgets.NewTable()
	table3.Rows = rows
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	tab