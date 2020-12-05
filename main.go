package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	passes := myIo()

	highest := 0
	for _, v := range passes {
		result := calcSeat(v)
		if result > highest {
			highest = result
		}
	}
	fmt.Println("Highest seat ID: ", highest)
}

func calcSeat(pass string) int {
	rowPart := pass[0:7]

	rowMax := 127
	rowMin := 0

	for _, v := range rowPart {
		if string(v) == "F" {
			rowMax = (rowMin + rowMax) / 2
		} else if string(v) == "B" {
			rowMin = (rowMin + rowMax + 1) / 2
		}

		fmt.Println("Code:", string(v), "min:", rowMin, "max:", rowMax)
	}

	colMax := 7
	colMin := 0

	colPart := pass[7:10]
	for _, v := range colPart {
		if string(v) == "L" {
			colMax = (colMin + colMax) / 2
		} else if string(v) == "R" {
			colMin = (colMin + colMax + 1) / 2
		}

		fmt.Println("Code:", string(v), "min:", colMin, "max:", colMax)
	}

	return ((rowMax+rowMin)/2)*8 + ((colMax + colMin) / 2)
}

func myIo() []string {
	f, err := os.Open("files\\day5")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var array []string

	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}
