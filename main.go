package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("You entered grade is: %s", passGrade())
	fmt.Printf("The size of the file is: %d bytes\n", fileSize("main.go"))
}

func passGrade() string {
	var status string
	level := 60

	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	if grade >= float64(level) {
		status = "passing"
	} else {
		status = "failing"
	}

	return status
}

func fileSize(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.Size()
}
