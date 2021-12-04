package day2

import (
	"fmt"
	"strconv"
	"strings"

	reader "advent-of-code-2021/lib"
)

type day2 struct {
	input string
}

func NewDay2(input string) *day2 {
	return &day2{
		input: input,
	}
}

func Run() {
	fmt.Println("** Day 2 **")

	input, err := reader.ReadFile("./day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := NewDay2(input)

	part1 := d.Part1()
	fmt.Println("Part 1: ", part1)

	part2 := d.Part2()
	fmt.Println("Part 2: ", part2)
}

func (d *day2) Part1() int {
	lines := strings.Split(d.input, "\n")
	var depth int
	var horizontal int
	for _, line := range lines {
		words := strings.Split(line, " ")

		direction := words[0]
		position, _ := strconv.Atoi(words[1])

		if direction == "up" {
			depth += -1 * position
		} else if direction == "down" {
			depth += position
		} else {
			horizontal += position
		}
	}

	return horizontal * depth
}

func (d *day2) Part2() int {
	lines := strings.Split(d.input, "\n")
	var depth int
	var horizontal int
	var aim int
	for _, line := range lines {
		words := strings.Split(line, " ")

		direction := words[0]
		position, _ := strconv.Atoi(words[1])

		if direction == "up" {
			aim += -1 * position
		} else if direction == "down" {
			aim += position
		} else {
			horizontal += position
			depth += aim * position
		}

	}

	return horizontal * depth
}
