package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
	"strings"
)

func solutionPart01(lines []string) {
	rules, updates := readRulesAndUpdates(lines)

	validUpdates := make([][]string, 0)
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			validUpdates = append(validUpdates, update)
		}
	}

	sum := sumOfMiddles(validUpdates)
	fmt.Printf("%d", sum)
}

func solutionPart02(lines []string) {
	rules, updates := readRulesAndUpdates(lines)

	invalidUpdates := make([][]string, 0)
	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fixedUpdates := make([][]string, 0)
	for _, invalidUpdate := range invalidUpdates {
		newUpdate := tryToFixUpdate(rules, invalidUpdate)
		if isValidUpdate(newUpdate, rules) {
			fixedUpdates = append(fixedUpdates, newUpdate)
		}
	}

	sum := sumOfMiddles(fixedUpdates)
	fmt.Printf("%d", sum)
}

func readRulesAndUpdates(lines []string) (map[string]string, [][]string) {
	rules := make(map[string]string)
	updates := make([][]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			rules[line] = line
		}
		if strings.Contains(line, ",") {
			update := Split(line, ",")
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func hasRule(item string, rules map[string]string) bool {
	_, ok := rules[item]
	return ok
}

func isValidUpdate(update []string, rules map[string]string) bool {
	valid := true
	for j := range update {
		if !isValidAt(j, update, rules) {
			valid = false
			break
		}
	}
	return valid
}

func isValidAt(index int, update []string, rules map[string]string) bool {
	// from index to right
	item := update[index]
	for i := index + 1; i < len(update); i++ {
		pair := item + "|" + update[i]
		if !hasRule(pair, rules) {
			return false
		}
	}
	// from index to left
	for i := index - 1; i >= 0; i-- {
		pair := update[i] + "|" + item
		if !hasRule(pair, rules) {
			return false
		}
	}
	return true
}

func getMiddleAsInt(update []string) int {
	middle := update[(len(update)-1)/2]
	return ParseInt(middle)
}

func tryToFixUpdate(rules map[string]string, update []string) []string {
	newUpdate := make([]string, len(update))
	copy(newUpdate, update)

	sort.SliceStable(newUpdate, func(i, j int) bool {
		key := newUpdate[i] + "|" + newUpdate[j]
		return hasRule(key, rules)
	})
	return newUpdate
}

func sumOfMiddles(updates [][]string) int {
	sum := 0
	for _, update := range updates {
		sum += getMiddleAsInt(update)
	}
	return sum
}

// https://adventofcode.com/2024/day/5
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
	//RunAdventOfCodeWithFile(solutionPart01, "day_05/testcases/input-part-01.txt")

	// part 02: using string or input file
	//RunAdventOfCodeWithString(solutionPart02, "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
	RunAdventOfCodeWithFile(solutionPart02, "day_05/testcases/input-part-02.txt")
}
