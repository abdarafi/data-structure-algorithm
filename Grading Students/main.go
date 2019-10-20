package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {

	for i, val := range grades{
		y := val % 5  // Not really necessary to declare this here,
		x := val % 10 // but it's easier to read the code like this
		if val >= 38 &&  10 - x < 3{
			grades[i] = val + 10 - x
		} else 	if val >= 38 && 5 - y < 3 {
			grades[i] = val + 5 - y
		}
	}
	return grades
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create("OUTPUT_PATH")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var grades []int32

	for i := 0; i < int(gradesCount); i++ {
		gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		gradesItem := int32(gradesItemTemp)
		grades = append(grades, gradesItem)
	}

	result := gradingStudents(grades)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
