package main

import (
	"fmt"
	"sort"
)

func day5() {
	passes := myIo()

	highest := 0
	var seatIds []int
	for _, v := range passes {
		result := calcSeat(v)

		seatIds = append(seatIds, result)

		if result > highest {
			highest = result
		}
	}

	fmt.Println("Highest seat ID: ", highest)
	sort.Ints(seatIds)

	previous := seatIds[0] - 1
	for _, v := range seatIds {
		fmt.Print(" ID: ", v)
		if v != previous+1 {
			fmt.Println("Found missing seat at: ", previous+1)
		}
		previous = v
	}
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
	}

	return ((rowMax+rowMin)/2)*8 + ((colMax + colMin) / 2)
}
