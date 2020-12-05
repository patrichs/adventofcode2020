package passport

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
}

func (p Passport) ToString() string {
	return "Byr: " + p.Byr + " Iyr: " + p.Iyr + " Eyr: " + p.Eyr + " Hgt: " + p.Hgt + " Hcl: " + p.Hcl + " Ecl: " + p.Ecl + " Pid: " + p.Pid
}

func (p Passport) Valid() bool {
	return validateBirthyear(p.Byr) && validateIssueyear(p.Iyr) &&
		validateExpyear(p.Eyr) && validateHeight(p.Hgt) && validateHaircolor(p.Hcl) &&
		validateEyecolor(p.Ecl) && validatePassportId(p.Pid)
}

func validateBirthyear(val string) bool {
	convert, _ := strconv.Atoi(val)
	return convert >= 1920 && convert <= 2002
}

func validateIssueyear(val string) bool {
	convert, _ := strconv.Atoi(val)
	return convert >= 2010 && convert <= 2020
}

func validateExpyear(val string) bool {
	convert, _ := strconv.Atoi(val)
	return convert >= 2020 && convert <= 2030
}

func validateHeight(val string) bool {
	pattern := `^(\d+)(cm|in)$`
	compiledRegex := regexp.MustCompile(pattern)
	submatches := compiledRegex.FindAllStringSubmatch(val, -1)
	if submatches == nil {
		return false
	}
	convert, _ := strconv.Atoi(submatches[0][1])
	if submatches[0][2] == "cm" {
		return convert >= 150 && convert <= 193
	} else if submatches[0][2] == "in" {
		return convert >= 59 && convert <= 76
	} else {
		return false
	}
}

func validateHaircolor(val string) bool {
	pattern := `^[#a-f0-9]{7}$`
	compiledRegex := regexp.MustCompile(pattern)
	if compiledRegex.FindString(val) == "" {
		return false
	} else {
		return true
	}
}

func validateEyecolor(val string) bool {
	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range valid {
		if strings.Contains(val, v) {
			return true
		}
	}
	return false
}

func validatePassportId(val string) bool {
	pattern := `^[0-9]{9}$`
	compiledRegex := regexp.MustCompile(pattern)
	if compiledRegex.FindString(val) == "" {
		return false
	} else {
		return true
	}
}
