package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println(calculateComission(50, 99))
	fmt.Println(calculateComission(1, 99))
}

func task1() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, p int
		fmt.Fscan(in, &n, &p)

		// calculate commission
		var sum float64
		var a int
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			_, float := math.Modf(calculateComission(a, p))
			sum += float
		}

		fmt.Fprintln(out, sum)
	}

}

func calculateComission(price int, precent int) float64 {
	return roundWithPrecise(float64(price) * float64(precent) / 100)
}

func roundWithPrecise(num float64) float64 {
	return math.Floor(num*100) / 100
}
