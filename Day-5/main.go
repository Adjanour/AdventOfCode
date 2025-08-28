package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	prefix int
	suffix int
}

func main() {
	rules, updates, err := getRulesAndUpdates("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validUpdates := [][]int{}
	for _, update := range updates {
		if len(update) == 0 {
			continue
		}
		if isValid(update, rules) {
			fmt.Println("Valid update:", update)
			validUpdates = append(validUpdates, update)
		}
	}

	total := 0
	for _, update := range validUpdates {
		total += middle(update)
	}

	fmt.Println("Total of middle values:", total)
}

// Reads rules and updates from file
func getRulesAndUpdates(filename string) ([]Rule, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	rules := []Rule{}
	updates := [][]int{}
	scanningRules := true
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			scanningRules = false
			continue
		}

		if scanningRules {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("invalid rule format: %s", line)
			}
			before, err1 := strconv.Atoi(parts[0])
			after, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return nil, nil, fmt.Errorf("invalid numbers in rule: %s", line)
			}
			rules = append(rules, Rule{prefix: before, suffix: after})
		} else {
			update := parseLine(line)
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return rules, updates, nil
}

// Converts a comma-separated line into a slice of ints
func parseLine(line string) []int {
	parts := strings.Split(line, ",")
	values := make([]int, 0, len(parts))
	for _, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Printf("Warning: skipping invalid number '%s'\n", p)
			continue
		}
		values = append(values, val)
	}
	return values
}

// Checks if an update satisfies all rules
func isValid(update []int, rules []Rule) bool {
	positions := make(map[int]int)
	for i, num := range update {
		positions[num] = i
	}

	for _, r := range rules {
		posPrefix, okPrefix := positions[r.prefix]
		posSuffix, okSuffix := positions[r.suffix]
		if okPrefix && okSuffix && posPrefix >= posSuffix {
			return false
		}
	}
	return true
}

// Returns the middle element of a slice
func middle(update []int) int {
	if len(update) == 0 {
		return -1
	}
	return update[len(update)/2]
}
