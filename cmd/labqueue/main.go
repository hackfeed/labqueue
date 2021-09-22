package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	inputFile string
)

func init() {
	flag.StringVar(&inputFile, "input", "students.txt", "file to read")
}

func main() {
	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	students := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		students = append(students, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(students), func(i, j int) { students[i], students[j] = students[j], students[i] })

	for i, student := range students {
		fmt.Printf("%d. %s\n", i+1, student)
	}
}
