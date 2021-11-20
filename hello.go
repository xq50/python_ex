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
		log.Fatalf("failed to initialize termui: %v", e