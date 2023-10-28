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
var Details map[string]string

func main() {
	Header = make(map[string]string)
	ReadFile()
	MakeHeader([]string{"PlateLenght", "PlateWidth", "PlatesQty"}, 2)
	MakeHeader([]string{"CutNum?", "CutThickness", "EdgeMargin", "CutBol?"}, 3)
	MakeHeader([]string{"OrderName"}, 4)
	MakeHeader([]string{"Material"}, 7)
	MakeHeader([]string{"MaterialThickness"}, 8)
	Details = make(map[string]string)

	printarray(Header)
}

func MakeHeader(param []string, line int) { //сделать функцию стения одной строки с передачей yjvthf строки в масиве array и возврате среза разобранных значений
	key := 0
	for _, ii := range array[line] {
		if string(ii) != "X" && string(ii) != "_" {
			Header[strconv.Itoa(line)+param[key]] += string(ii)
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
