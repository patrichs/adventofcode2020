package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("files\\day3")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var a []string

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var product []int

	product = append(product, loopThrough(1, len(a), a, 1, 1), loopThrough(3, len(a), a, 1, 1),
		loopThrough(5, len(a), a, 1, 1), loopThrough(7, len(a), a, 1, 1), loopThrough(1, len(a), a, 2, 2))

	var total = 1
	for _, v := range product {
		fmt.Println("Total trees: ", v)
		total *= v
	}

	fmt.Println("Total product of all trees: ", total)
}

func loopThrough(x int, y int, a []string, iterate int, starti int) int {
	amount := 0
	xLen := len(a[0])
	xIndex := 0

	for i := starti; i < y; i += iterate {
		s := strings.Split(a[i], "")
		xIndex += x

		if xIndex >= xLen {
			if (xIndex)%x != 0 {
				distance := xIndex - 30
				xIndex = distance - 1
			} else if xIndex-x != 30 {
				distance := xIndex - 30
				xIndex = distance - 1
			} else {
				xIndex = x - 1
			}
		}

		if s[xIndex] == "#" {
			amount++
			s[xIndex] = "X"
		} else {
			s[xIndex] = "-"
		}

		fmt.Println("Splitting string", s, "trees found: ", amount, "iterations:", i, "currentX:", xIndex)
	}
	fmt.Println("trees found:", amount)

	return amount
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//func main() { // day 2
//	f, err := os.Open("files\\day2")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer f.Close()
//
//	scanner := bufio.NewScanner(f)
//
//	var a []string
//
//	for scanner.Scan() {
//		a = append(a, scanner.Text())
//	}
//
//	if err := scanner.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	var pattern = `^(\d+)-(\d+)\s(\w)$`
//	var compiledRegex = regexp.MustCompile(pattern)
//	var matchedCount = 0
//	var indexOccuredCount = 0
//
//	for _, v := range a {
//		var split = strings.Split(v, ":")
//		var submatches = compiledRegex.FindAllStringSubmatch(split[0], -1)
//		var min, _ = strconv.Atoi(submatches[0][1])
//		var max, _ = strconv.Atoi(submatches[0][2])
//		var charMatch = submatches[0][3]
//		var rightsideRegexp = regexp.MustCompile(charMatch)
//
//		amount := rightsideRegexp.FindAllStringSubmatch(split[1], -1)
//
//		var indexOccurence = rightsideRegexp.FindAllStringSubmatchIndex(split[1], -1)
//
//		minHit := false
//		maxHit := false
//
//		for _, v := range indexOccurence {
//			if v[0] == min {
//				minHit = true
//			} else if v[0] == max {
//				maxHit = true
//			}
//		}
//
//		if minHit && !maxHit {
//			indexOccuredCount++
//		}
//
//		if !minHit && maxHit {
//			indexOccuredCount++
//		}
//
//		matched := len(amount) >= min && max >= len(amount)
//		if matched {
//			matchedCount++
//		}
//
//		fmt.Println("Leftside:", split[0], "Rightside:", split[1], "Min:", min, "Max:", max, "character To Match:", charMatch, "matched:", matched, "indexOccur:", indexOccurence)
//		//fmt.Printf("%q\n", compiledRegex.FindAllStringSubmatch(split[0], -1))
//	}
//
//	fmt.Println("Matched Total: ", matchedCount, "Did not match: ", 1000 - matchedCount, "Index occurence Count: ", indexOccuredCount)
//}

//func main() { // day 1
//	f, err := os.Open("files\\day1")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer f.Close()
//
//	scanner := bufio.NewScanner(f)
//
//	var a []int
//
//	for scanner.Scan() {
//		s, parseError := strconv.Atoi(scanner.Text())
//
//		if parseError != nil {
//			log.Fatal(parseError)
//		}
//
//		a = append(a, s)
//	}
//
//	if err := scanner.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	printSlice(a)
//
//	for i, v := range a {
//		for i2, v2 := range a {
//			if v == 2020 - v2 {
//				fmt.Println("value: ", v, "value2: ", v2, "iterations: ", i * i2, "sum: ", v * v2)
//			}
//		}
//	}
//
//	for i, v := range a {
//		for i2, v2 := range a {
//			for i3, v3 := range a {
//				if v == 2020 - v2 - v3 {
//					fmt.Println("value: ", v, "value2: ", v2, "value3:", v3, "iterations: ", i * i2 * i3, "sum: ", v * v2 * v3)
//				}
//			}
//		}
//	}
//}

//func main() {
//	content, err := ioutil.ReadFile("files\\day1")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(string(content))
//}
