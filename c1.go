package main

import (
	"fmt"
)

func main() {
	steps := 343
	buffer := []int{0}
	currentPos := 0

	for i := 1; i <= 2017; i++ {
		currentPos = (currentPos + steps) % len(buffer)
		buffer = append(buffer[:currentPos+1], buffer[currentPos:]...)
		buffer[currentPos+1] = i
		currentPos++
	}

	for i, val := range buffer {
		if val == 2017 {
			fmt.Println(buffer[(i+1)%len(buffer)])
			break
		}
	}
}
