package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cat_random.go filename")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator with the current time

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := range lines {
		j := rand.Intn(i + 1)
		lines[i], lines[j] = lines[j], lines[i]
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
