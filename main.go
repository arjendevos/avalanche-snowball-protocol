package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	totalNodes = 1000
	threshold  = 3
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	preferences := []string{"orange", "blue", "green"}

	// Generate nodes with preference
	nodes := generateNodesWithPref(preferences)

	// Every node queries 10 other nodes
	start := time.Now()
	for ni := 0; ni < totalNodes; ni++ {
		success := 0
		totalRound := 0
		startTime := time.Now()

		for ri := 0; success < 20; ri++ {
			totalRound++
			chosenPrefs := make(map[string]int, 10)
			for i := 0; i <= 20; i++ {
				diffNodePref := nodes[rand.Intn(totalNodes)]
				_, exists := chosenPrefs[diffNodePref]
				if !exists {
					chosenPrefs[diffNodePref] = 1
				} else {
					chosenPrefs[diffNodePref]++
				}
			}

			// fmt.Println(chosenPrefs)
			mostChosenPref, highestChosen := getMostChosenPref(nodes[ni], preferences, chosenPrefs)

			// fmt.Println(nodes[ni])

			if highestChosen >= threshold {
				if mostChosenPref == nodes[ni] {
					success++
				} else {
					nodes[ni] = mostChosenPref
					success = 1
				}
			} else {
				success = 0
			}
		}

		since := time.Since(startTime)

		fmt.Printf("\n node: %v, round: %v, chosen: %v, dur: %s, succ: %v\n", ni, totalRound, nodes[ni], since, success)

		// fmt.Println("OWN NEW: ", node)
	}

	since := time.Since(start)
	fmt.Printf("\n%s\n", since)

}

func generateNodesWithPref(preferences []string) []string {
	nodes := make([]string, totalNodes)
	for i := 0; i < totalNodes; i++ {
		nodes[i] = preferences[rand.Intn(len(preferences))]
	}

	return nodes
}

func getMostChosenPref(initialPref string, preferences []string, chosenPrefs map[string]int) (string, int) {
	newPref := initialPref
	highestChosen := 0
	for _, preference := range preferences {
		if chosenPrefs[preference] > highestChosen {
			highestChosen = chosenPrefs[preference]
			newPref = preference
		}
	}

	return newPref, highestChosen
}
