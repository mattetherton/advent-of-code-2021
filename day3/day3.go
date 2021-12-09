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

	part2 := d.Part2()
	fmt.Println("Part 2: ", part2)
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

type finder func([]string, int) string

func (d *day3) Part2() int {
	input := strings.Split(d.input, "\n")

	mostCommonFinder := func(input []string, index int) string { return mostCommon(input, index) }
	oxygenRating, _ := ratingGenerator(input, mostCommonFinder)

	leastCommonFinder := func(input []string, index int) string { return leastCommon(input, index) }
	co2Rating, _ := ratingGenerator(input, leastCommonFinder)

	return int(oxygenRating * co2Rating)
}

func ratingGenerator(input []string, find finder) (int64, error) {
	for i := 0; i < len(input[0]); i++ {
		input = reducer(input, i, find(input, i))
		if len(input) == 1 {
			break
		}
	}

	return strconv.ParseInt(input[0], 2, 64)
}

func reducer(input []string, index int, target string) (reduced []string) {
	for _, line := range input {
		if string(line[index]) == target {
			reduced = append(reduced, line)
		}
	}
	return reduced
}

func filter(input []string, index int) (zeroes int, ones int) {
	for _, val := range input {
		if val[index] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	return zeroes, ones
}

func mostCommon(input []string, index int) string {
	zeroes, ones := filter(input, index)
	if zeroes > ones {
		return "0"
	}
	return "1"
}

func leastCommon(input []string, index int) string {
	zeroes, ones := filter(input, index)
	if zeroes > ones {
		return "1"
	}
	return "0"
}
