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
	ca