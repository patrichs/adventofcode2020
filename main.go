package main

import (
	"adventofcode/passport"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("files\\day4")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var passports []string
	var temp string

	for scanner.Scan() {
		temp += scanner.Text()
		temp += " "
		if len(scanner.Text()) <= 0 {
			passports = append(passports, temp)
			temp = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	var validPassports []passport.Passport

	count := 0
	for _, v := range passports {
		valid := true
		for _, vv := range required {
			if !strings.Contains(v, vv) {
				valid = false
			}
		}
		if valid {
			count++

			splitted := strings.Split(v, " ")

			newPassport := passport.Passport{
				Byr: getWhereName(splitted, "byr"),
				Iyr: getWhereName(splitted, "iyr"),
				Eyr: getWhereName(splitted, "eyr"),
				Hgt: getWhereName(splitted, "hgt"),
				Hcl: getWhereName(splitted, "hcl"),
				Ecl: getWhereName(splitted, "ecl"),
				Pid: getWhereName(splitted, "pid"),
			}
			validPassports = append(validPassports, newPassport)
		}
	}
	validCount := 0
	for _, v := range validPassports {
		if v.Valid() {
			fmt.Println(v.ToString())
			validCount++
		}
	}

	fmt.Println("Amount of passports:", len(passports), "Amount of first part valid passports:", count)
	fmt.Println("Second part validated passports amount:", validCount)
}

func getWhereName(splittedPassport []string, which string) string {
	for _, v := range splittedPassport {
		if which == getLeftside(v) {
			return getRightside(v)
		}
	}
	return ""
}

func getLeftside(val string) string {
	return strings.Split(val, ":")[0]
}

func getRightside(val string) string {
	return strings.Split(val, ":")[1]
}
