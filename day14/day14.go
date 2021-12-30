package day14

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CountMap map[string]int

func (c CountMap) PrintResult(lastChar string) {
	maxRecord := 0
	minRecord := int(^uint(0) >> 1)
	charCountMap := make(CountMap)
	for k, v := range c {
		ch1 := string(k[0])
		charCountMap[ch1] += v
	}
	charCountMap[lastChar] += 1
	for _, v := range charCountMap {
		if v > maxRecord {
			maxRecord = v
		}
		if v < minRecord {
			minRecord = v
		}

	}
	fmt.Println(maxRecord - minRecord)
}

func ApplyRules(countMap CountMap, rules RulesMap) CountMap {
	newCountMap := make(CountMap)
	for pair, count := range countMap {
		rulePairs, ok := rules[pair]
		if !ok {
			newCountMap[pair] += count
			continue
		}
		newCountMap[rulePairs[0]] += count
		newCountMap[rulePairs[1]] += count
	}
	return newCountMap
}

type RulesMap map[string][2]string

func Solution() {
	file, _ := os.Open("day14/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	rulesMap := make(RulesMap)
	template := ""
	readingRules := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRules = true
			continue
		}
		if !readingRules {
			template = line
			continue
		}
		fields := strings.Fields(line)
		rule := fields[0]
		value := fields[2]
		rulesMap[rule] = [2]string{string(rule[0]) + value, value + string(rule[1])}
	}
	countMap := make(CountMap)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		countMap[pair] += 1
	}
	Part1(countMap, rulesMap, string(template[len(template)-1]))
	Part2(countMap, rulesMap, string(template[len(template)-1]))
}

func Part1(countMap CountMap, rulesMap RulesMap, lastChar string) {
	curr := countMap
	for i := 0; i < 10; i++ {
		curr = ApplyRules(curr, rulesMap)
	}
	curr.PrintResult(lastChar)
}

func Part2(countMap CountMap, rulesMap RulesMap, lastChar string) {
	curr := countMap
	for i := 0; i < 40; i++ {
		fmt.Println("step", i)
		curr = ApplyRules(curr, rulesMap)
	}
	curr.PrintResult(lastChar)
}
