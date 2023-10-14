package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	registers := make(map[string]int)
	mulCounter := 0

	for i := 0; i < len(instructions); {
		parts := strings.Fields(instructions[i])
		cmd, x, y := parts[0], parts[1], parts[2]
		valY, _ := strconv.Atoi(y)

		if y >= "a" && y <= "z" {
			valY = registers[y]
		}

		switch cmd {
		case "set":
			registers[x] = valY
		case "sub":
			registers[x] -= valY
		case "mul":
			registers[x] *= valY
			mulCounter++
		case "jnz":
			valX, _ := strconv.Atoi(x)
			if x >= "a" && x <= "z" {
				valX = registers[x]
			}
			if valX != 0 {
				i += valY
				continue
			}
		}
		i++
	}

	fmt.Println(mulCounter)
}

func partTwo() {
	b := 57*100 + 100000
	c := b + 17000
	h := 0

	for x := b; x <= c; x += 17 {
		if !isPrime(x) {
			h++
		}
	}

	fmt.Println(h)
}

func main() {
	partOne()
	partTwo()
}
