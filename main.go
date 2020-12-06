package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	answers := myIo()
	answersPart2 := myIoPart2()

	total := 0
	for _, v := range answers {
		result := answeredYes(v)
		total += result
	}

	totalPart2 := 0
	for _, v := range answersPart2 {
		result := answeredYesPart2(v)
		totalPart2 += result
		fmt.Println("String", v, "Result:", result)
	}

	println("Total score Part 1:", total, "Total score Part 2:", totalPart2)
}

func answeredYes(parsed string) int {
	a := "abcdefghijklmnopqrstuvwxyz"

	score := 0
	for _, v := range a {
		if strings.Contains(parsed, string(v)) {
			score++
		}
	}

	return score
}

func answeredYesPart2(parsed string) int {
	a := "abcdefghijklmnopqrstuvwxyz"

	split := strings.Split(parsed, ";")

	score := 0
	for _, v := range a {
		temp := true
		for _, vv := range split {
			if len(vv) < 1 {
				continue
			}
			if !strings.Contains(vv, string(v)) {
				temp = false
			}
		}
		if temp {
			score++
		}
	}

	return score
}

func myIoPart2() []string {
	f, err := os.Open("files\\day6")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var array []string
	temp := ""

	for scanner.Scan() {
		temp += scanner.Text()
		temp += ";"
		if len(scanner.Text()) <= 0 {
			array = append(array, temp)
			temp = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

func myIo() []string {
	f, err := os.Open("files\\day6")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var array []string
	temp := ""

	for scanner.Scan() {
		temp += scanner.Text()
		if len(scanner.Text()) <= 0 {
			array = append(array, temp)
			temp = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}
