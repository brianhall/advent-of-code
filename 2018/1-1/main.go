package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	target := 0

	file, err := os.Open("data/input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		target = target + num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(target)
}
