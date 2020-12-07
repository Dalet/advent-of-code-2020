package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type bag struct {
	name          string
	children      map[string]childBag
	totalChildren int
}

type childBag struct {
	bagCount int
	bag      *bag
}

var allBags = make(map[string]*bag)

func getOrCreateBag(name string) (b *bag) {
	b, exists := allBags[name]
	if !exists {
		b = new(bag)
		b.name = name
		b.children = make(map[string]childBag)
		b.totalChildren = -1 // calculate later
		allBags[name] = b
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	bagRegexp := regexp.MustCompile(`(?:(\d+)\s+)?([a-z]+\s+[a-z]+)\s+bags?`)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		matches := bagRegexp.FindAllStringSubmatch(line, -1)
		if matches[1][1] == "" {
			continue // contains no other bag, does not need to be checked
		}

		parentBag := getOrCreateBag(matches[0][2])
		for _, match := range matches[1:] {
			child := getOrCreateBag(match[2])
			childBagCount, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			parentBag.children[child.name] = childBag{childBagCount, child}
		}
	}

	bagsEventuallyContainingShinyGold := 0
	for _, bag := range allBags {
		if couldContain(bag, "shiny gold") {
			bagsEventuallyContainingShinyGold++
		}
	}

	fmt.Println("Part One:", bagsEventuallyContainingShinyGold, "bag colors can eventually contain shiny gold bags")

	shinyGold, found := allBags["shiny gold"]
	if !found {
		panic(nil)
	}

	fmt.Println("Part Two:", countChildren(shinyGold), "bags required inside", shinyGold.name, "bags")
}

func couldContain(pBag *bag, bagName string) bool {
	for _, child := range pBag.children {
		if child.bag.name == bagName || couldContain(child.bag, bagName) {
			return true
		}
	}

	return false
}

func countChildren(pBag *bag) (count int) {
	if pBag.totalChildren != -1 {
		return pBag.totalChildren
	}

	count = 0
	for _, child := range pBag.children {
		count += child.bagCount + child.bagCount*countChildren(child.bag)
	}

	pBag.totalChildren = count
	return
}
