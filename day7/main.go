package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type node struct {
	name     string
	children map[string]*node
}

var allBags = make(map[string]*node)

func getOrCreateBag(name string) (bag *node) {
	bag, exists := allBags[name]
	if !exists {
		bag = new(node)
		bag.name = name
		bag.children = make(map[string]*node)
		allBags[name] = bag
	}
	return bag
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	bagRegexp := regexp.MustCompile(`(?:(?P<num>\d+)\s+)?(?P<color>[a-z]+\s+[a-z]+)\s+bags?`)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		matches := bagRegexp.FindAllStringSubmatch(line, -1)
		if len(matches) <= 1 {
			continue // contains no other bag, does not need to be checked
		}

		containerBag := getOrCreateBag(matches[0][2])
		for _, match := range matches[1:] {
			containedBag := getOrCreateBag(match[2])
			containerBag.children[containedBag.name] = containedBag
		}
	}

	bagsEventuallyContainingShinyGold := 0
	for _, bag := range allBags {
		if couldContain(bag, "shiny gold") {
			bagsEventuallyContainingShinyGold++
		}
	}

	fmt.Println("Part One:", bagsEventuallyContainingShinyGold, "bag colors can eventually contain shiny gold bags")
}

func couldContain(bag *node, bagName string) bool {
	for _, child := range bag.children {
		if child.name == bagName || couldContain(child, bagName) {
			return true
		}
	}

	return false
}
