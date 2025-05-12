package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Do we need fmt?

var shakespeare = [11]int{115, 104, 97, 107, 101, 115, 112, 101, 97, 114, 101} //Will store correct values for the word shakespeare
var max int = 123
var min int = 97    // Random Num generation makes range [min,max). max is excluded so we make 122+1
var attempt [11]int //Could keep this a global? We overide it everytime anyways right?

var callcount int = 0

func generate() [11]int {
	for i := 0; i < 11; i++ {
		attempt[i] = rand.Intn(max-min) + min
	}
	return attempt
}

// Comparing strings takes longer than integers due to lookups

func evaluate(attempt_ [11]int) int {
	for i := 0; i < 11; i++ {
		//if it matchs, continue, else break and return the score (i)
		if attempt_[i] != shakespeare[i] {
			return i
		} // Will stop early and not waste time on the loop
	}
	return 11
}

func arrayToString(a [11]int) string {
	var answer string = ""
	for i := 0; i < 11; i++ {
		answer = answer + string(a[i])
	}

	return answer
}

func main() {
	var elapsed int64 = 0

	var best_score int = 0
	var best_attempt [11]int

	t := time.Now()

	for time.Since(t) < 300000000000 {
		start := time.Now()

		var attempt = generate()
		var score = evaluate(attempt)

		if score > best_score {
			best_score = score
			best_attempt = attempt
		}
		callcount++
		elapsed += time.Since(start).Nanoseconds()
	}

	fmt.Println(arrayToString(best_attempt))
	fmt.Println(best_score)
	fmt.Println("Total words generated:", callcount)
	fmt.Println("Average Processing Time (Nanoseconds) :", float64(elapsed)/float64(callcount))

}
