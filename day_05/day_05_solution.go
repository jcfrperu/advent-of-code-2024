package main

import (
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"strings"
)

func solutionPart01(lines []string) {
	rules := make(map[string]string)
	updates := make([][]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			rules[line] = line
		}
		if strings.Contains(line, ",") {
			data := Split(line, ",") //[]string
			updates = append(updates, data)
		}
	}

	validLines := make([][]string, 0)
	for _, update := range updates {
		valid := true
		for j, _ := range update {
			if !isValid(rules, update, j) {
				valid = false
				break
			}
		}
		if valid {
			validLines = append(validLines, update)
		}
	}

	sum := 0
	for _, line := range validLines {
		sum += getMiddleAsInt(line)
	}

	fmt.Printf("%d", sum)
}

func hasRule(item string, rules map[string]string) bool {
	_, ok := rules[item]
	return ok
}

func isValid(rules map[string]string, list []string, index int) bool {
	// from index to right
	item := list[index]
	//fmt.Printf("pivot: %s\n", item)
	for i := index + 1; i < len(list); i++ {
		pair := item + "|" + list[i]
		//fmt.Printf("-> %s\n", pair)
		if !hasRule(pair, rules) {
			return false
		}
	}

	// from index to left
	for i := index - 1; i >= 0; i-- {
		pair := list[i] + "|" + item
		//fmt.Printf("<- %s\n", pair)
		if !hasRule(pair, rules) {
			return false
		}
	}

	return true
}

func getMiddleAsInt(lines []string) int {
	size := len(lines)
	middle := lines[(size-1)/2]
	return ParseInt(middle)
}

func solutionPart02(lines []string) {
	count := 0

	fmt.Printf("%d", count)
}

// https://adventofcode.com/2024/day/5
func main() {
	// part 01: using string or input file
	//RunAdventOfCodeWithString(solutionPart01, "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
	//RunAdventOfCodeWithFile(solutionPart01, "day_05/testcases/input-part-01.txt")

	// part 02: using string or input file
	RunAdventOfCodeWithString(solutionPart02, "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")
	//RunAdventOfCodeWithFile(solutionPart02, "day_05/testcases/input-part-02.txt")
}
