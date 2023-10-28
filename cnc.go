package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var array []string
var Header map[string]string
var Details map[string]map[string]int

func main() {
	Header = make(map[string]string)
	ReadFile()
	MakeHeader([]string{"PlateLenght", "PlateWidth", "PlatesQty"}, 2)
	MakeHeader([]string{"CutNum?", "CutThickness", "EdgeMargin", "CutBol?"}, 3)
	MakeHeader([]string{"OrderName"}, 4)
	MakeHeader([]string{"Material"}, 7)
	MakeHeader([]string{"MaterialThickness"}, 8)
	MakeHeader([]string{"_", "DetailsQty"}, 15)
	if Header["DetailsQty"] == "" {
		MakeHeader([]string{"_", "DetailsQty"}, 17)
		Header["DetailStartIndex"] = "18"
	} else {
		Header["DetailStartIndex"] = "16"
	}

	// Details = make(map[string]map[string]int)

	Details = map[string]map[string]int{
		"деталь1": {"длина": 1, "ширина": 2},
		"деталь2": {"длина": 3, "ширина": 4},
	}
	fmt.Println(Details)

	printarray(Header)
	printdetails(Details)
}

func MakeHeader(param []string, line int) {
	key := 0
	for _, ii := range array[line] {
		if string(ii) != "X" && string(ii) != "_" && string(ii) != ">" {
			// Header[strconv.Itoa(line)+param[key]] += string(ii)
			Header[param[key]] += string(ii)
		} else {
			key++
		}
	}
}

func ReadFile() { //This function reads user sct file and writes it to array.
	fmt.Print("Please input the name of file to read: ")
	filename, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	filename = strings.Trim(filename, "\r\n")
	f, err := os.Open(filename)
	if err == nil {
		readdata := bufio.NewScanner(f)
		for readdata.Scan() {
			array = append(array, readdata.Text())
		}
	} else {
		fmt.Println("File open error! Please check the filename. " + filename)
	}
	fmt.Println("Прочитано строк: ", len(array))

	for i, arrayslice := range array {
		fmt.Println(i, " ", arrayslice)
	}
}
func printarray(array map[string]string) {
	for i, ii := range array {
		iiint, iierr := strconv.Atoi(ii)
		if iierr == nil {
			fmt.Println(i, " ", iiint)
		} else {
			fmt.Println(i, " ", ii)
		}
	}
}
func printdetails(array map[string]map[string]int) {
	for i, ii := range array {
		for j, jj := range ii {
			fmt.Println(i, " ", j, " ", jj)
		}
	}
}
