package main

import (
	"fmt"
	"math/rand"

	"github.com/Rhymond/go-money"
)

type trade [2]int

func main() {
	tradesTaken := 200
	trades := make([]trade, tradesTaken)
	var wins, losses int
	principal := 3000000
	investment := principal

	profit := 0.1
	loss := 0.05
	min, max := 0, 1

	for i := range trades {
		profitDecider := rand.Intn(max - min + 1)
		switch profitDecider {
		case 1:
			principal += int(float64(principal) * profit)
			trades[i][0] = principal
			trades[i][1] = 10
			wins++
		case 0:
			principal -= int(float64(principal) * loss)
			trades[i][0] = principal
			trades[i][1] = -5
			losses++
		}
		investment += 1000
		principal += 1000
	}

	for i, trade := range trades {
		fmt.Println(i, " ", trade[0], " @ ", trade[1])
	}

	principalMoney := money.New(int64(principal*100), "INR")
	principalFormatted := principalMoney.Display()
	fmt.Println("Investment : ", investment, "\nReturns : ", principalFormatted)
	fmt.Println("Win % : ", (float64(wins) / float64(tradesTaken) * 100))
	fmt.Println("Loss % : ", (float64(losses) / float64(tradesTaken) * 100))
}
