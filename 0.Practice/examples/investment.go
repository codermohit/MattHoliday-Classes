package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// a single trade, which has the amount after the trade and the return on that trade
type trade [2]int

func main() {
	var fundsWithdrawn int
	//principal amount to trade
	principal := 20000
	investment := 20000
	//Number of trades that will be done
	trades := make([]trade, 60)

	//minimum and maximum returns possible on a trade
	min := 7
	max := 15

	for i := range trades {
		//returns is the sum of the principal and the returns on a trade
		var returns float64

		//interest to be different on each trade
		interest := float64(rand.Intn(max-min+1) + min)

		returns = (float64(principal) * float64(interest) / 100) + float64(principal)

		//Putting the information of the trade in the trades slice
		trades[i] = trade{int(returns), int(interest)}
		principal = int(returns) + 20000
		investment += 20000

		if principal > 60000000 {
			principal /= 2
			investment = principal
			fundsWithdrawn += 30000000
		}
	}

	for i, val := range trades {
		fmt.Printf("%d : %d @ %d\n", i, val[0], val[1])
	}

	fmt.Printf("investment : %d\n, principal : %d\n", investment, principal)
	fmt.Println("Funds withdrawn : ", fundsWithdrawn)

	if fundsWithdrawn == 0 {
		totalReturns := ((principal - investment) / investment) * 100
		fmt.Println("Total Returns : ", totalReturns, "%")
	}
	/* Code to Generate a graph for the trade-principal */
	p := plot.New()

	p.Title.Text = "Trade Data"
	p.X.Label.Text = "Trades"
	p.Y.Label.Text = "Amount"

	points := make(plotter.XYs, len(trades))

	for i, val := range trades {
		points[i].X = float64(i)
		points[i].Y = float64(val[0])
	}
	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}

	p.Add(s)

	fileName := "plot.png"
	// Save the plot to a PNG file
	if err := p.Save(10*vg.Inch, 8*vg.Inch, fileName); err != nil {
		panic(err)
	}
	cmd := exec.Command("kitty", "+kitten", "icat", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error displaying image: %v\n", err)
	}
}
