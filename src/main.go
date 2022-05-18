package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ChunksSlower(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

func readFile(nametxt string) {
	textnametxt := "text/" + nametxt
	file, err := os.Open(textnametxt)

	if err != nil {
		log.Fatalf("Error when opening file : %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	stringIndex := []string{}
	arraySliceIndex := [][]string{}
	i := 0
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		stringIndex = ChunksSlower(fileScanner.Text(), 3)
		j := 0
		arrayStringIndex := []string{}
		for _, stringIndex1 := range stringIndex {
			arrayStringIndex = append(arrayStringIndex, stringIndex1)
			j++
		}
		//num1 = st[0]
		arraySliceIndex = append(arraySliceIndex, arrayStringIndex)
		i++

	}
	readArrayIndex(arraySliceIndex)
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file : %s", err)
	}

	file.Close()
}
func readArrayIndex(arraySliceIndex [][]string) {
	i := 0
	indexstring := ""
	for i < 9 {
		all := ""

		for _, s := range arraySliceIndex {
			all = all + s[i]
		}

		num := parseInNumber(all)
		indexstring = indexstring + num
		i++

	}

	fmt.Printf("=> %s", indexstring)
	fmt.Println("\t")
}

func parseInNumber(all string) string {
	switch all {

	case "     |  |":
		return "1"

	case " _  _||_ ":
		return "2"

	case " _  _| _|":
		return "3"

	case "   |_|  |":
		return "4"

	case " _ |_  _|":
		return "5"

	case " _ |_ |_|":
		return "6"

	case " _   |  |":
		return "7"

	case " _ |_||_|":
		return "8"

	case " _ |_| _|":
		return "9"

	case " _ | ||_|":
		return "0"

	default:
		return "?"
	}

}
func main() {
	readFile("000000000.txt")
	readFile("111111111.txt")
	readFile("222222222.txt")
	readFile("333333333.txt")
	readFile("444444444.txt")
	readFile("555555555.txt")
	readFile("666666666.txt")
	readFile("777777777.txt")
	readFile("888888888.txt")
	readFile("999999999.txt")

	readFile("123456789.txt")

	readFile("123456790.txt")

	readFile("error.txt")

}
