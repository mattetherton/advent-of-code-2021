package day3

import (
	"fmt"
	"strconv"
	"strings"

	reader "advent-of-code-2021/lib"
)

type day3 struct {
	input string
}

func NewDay3(input string) *day3 {
	return &day3{
		input: input,
	}
}

func Run() {
	fmt.Println("** Day 3 **")

	input, err := reader.ReadFile("./day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := NewDay3(input)

	part1 := d.Part1()
	fmt.Println("Part 1: ", part1)

	// part2 := d.Part2()
	// fmt.Println("Part 2: ", part2)
}

func (d *day3) Part1() int {
	input := strings.Split(d.input, "\n")

	var gamma string
	var epsilon string
	for i := 0; i < len(input[0]); i++ {
		var zeroes int
		var ones int

		for _, val := range input {
			if val[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)

	return int(epsilonDecimal * gammaDecimal)
}
