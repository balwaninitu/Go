package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// experimentOne()
	// readingSimpleLineFromConsole()
}

// Example 1: Size Experiment
func experimentOne() {
	bigString := "Paste a very big string, more than 4KB in one line"
	reader := bufio.NewReader(strings.NewReader(bigString))
	line, _, _ := reader.ReadLine() // read 4KB at once

	fmt.Println("Output From reader.Readline() Below:")
	fmt.Println(string(line))

	fmt.Println("\n................................\n")

	scanner := bufio.NewScanner(strings.NewReader(bigString))
	scanner.Scan() // read 64KB at once
	fmt.Println("Output From scanner.Scan() Below:")
	fmt.Println(scanner.Text())
}

// Example 2: Reading a line
func readingSimpleLineFromConsole() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	fmt.Println("Readline:", string(line))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Scan Line:", scanner.Text())
}
