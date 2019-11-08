package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type portfolio struct {
	buyVal  int
	sellVal int
	share   int
}

var portfolios []portfolio

func stockHistory(initialInvestment int, monthlyContribution int, stockPrices [][]int) {
	if len(stockPrices) == 1 {
		return
	}
	lsp := stockPrices[len(stockPrices)-1]
	var selectedI, selectedJ int
	var selectedRatio float64
	for i, sp := range stockPrices {
		for j, p := range sp {
			ratio := float64(lsp[j]) / float64(p)
			if lsp[j] > p && ratio > selectedRatio {
				selectedI = i
				selectedJ = j
				selectedRatio = ratio
			}
		}
	}
	if selectedRatio == 0.0 {
		return
	}
	share := initialInvestment + selectedI*monthlyContribution
	if share > 0 {
		pf := portfolio{buyVal: stockPrices[selectedI][selectedJ], sellVal: lsp[selectedJ], share: share}
		portfolios = append(portfolios, pf)
	}
	stockHistory(monthlyContribution, monthlyContribution, stockPrices[selectedI+1:])
	return
}

func main() {
	var initialInvestment, monthlyContribution int
	var stockPrices [][]int
	fmt.Scanln(&initialInvestment)
	fmt.Scanln(&monthlyContribution)
	for {
		var l string
		var sp []int
		_, err := fmt.Scanln(&l)
		if err == io.EOF {
			break
		}
		ss := strings.Split(l, ",")
		for _, s := range ss {
			si, _ := strconv.Atoi(s)
			sp = append(sp, si)
		}
		stockPrices = append(stockPrices, sp)
	}
	stockHistory(initialInvestment, monthlyContribution, stockPrices)

	// fmt.Println(portfolios)
	sumVal := 0.0
	sumShares := 0
	for _, p := range portfolios {
		sumVal += float64(p.sellVal) / float64(p.buyVal) * float64(p.share)
		sumShares += p.share
	}
	fmt.Println(int(math.Round(sumVal)) - sumShares)
}
