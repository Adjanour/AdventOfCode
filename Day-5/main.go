package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	prefix int
	suffix int
}

var updates [][]int
var valid_updates [][]int

var rules []rule

func main() {
	getOrderAndRules()
	fmt.Println(updates)
	fmt.Println(rules)
	for _, update := range updates {
		if isValid(update, rules) {
			fmt.Println("Valid update:", update)
			valid_updates = append(valid_updates, update)
		}
	}
	total := 0
	for _, item := range valid_updates {
		mid := middle(item)
		total += mid
	}
	fmt.Println(total)
}

func getOrderAndRules() {
	// Read lines from file and parse
	// stop on reaching empty line
	// for each line read separate by | put into datastructure
	f, _ := os.Open("input.txt")
	var bf = bufio.NewScanner(f)
	defer f.Close()

	for bf.Scan() {
		// end when you meet an empty new line
		if bf.Text() == "" {
			break
		}
		line := strings.Split(bf.Text(), "|")
		mut := []int{}
		for _, item := range line {
			val, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			mut = append(mut, val)
		}
		rules = append(rules, rule{mut[0], mut[1]})
	}

	for bf.Scan() {
		line := strings.Split(bf.Text(), ",")
		mut := []int{}
		for _, item := range line {
			val, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			mut = append(mut, val)
		}
		updates = append(updates, mut)
	}
}

func isValid(update []int, rules []rule) bool {
	positions := make(map[int]int)
	for i, page := range update {
		positions[page] = i
	}
	for _, rules := range rules {
		if posx, ok := positions[rules.prefix]; ok {
			if posy, ok := positions[rules.suffix]; ok {
				if posx >= posy {
					return false
				}
			}
		}
	}
	return true
}

func middle(update []int) int {
	if len(update) == 0 {
		return -1
	}
	return update[len(update)/2]
}
