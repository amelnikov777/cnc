package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var array []string
var Header map[string]string
var Details map[string]string

func main() {
	Header = make(map[string]string)
	ReadFile()
	MakeHeader()

}

func MakeHeader() { //сделать функцию стения одной строки с передачей yjvthf строки в масиве array и возврате среза разобранных значений
	CurrentString := ""
	key := 0
	for _, ii := range array[2] {
		if string(ii) != "X" && string(ii) != "_" {
			CurrentString = CurrentString + string(ii)
		} else {
			key++
			switch key {
			case 1:
				Header["PlateLenght"] = CurrentString
			case 2:
				Header["PlateWidth"] = CurrentString
			}
			CurrentString = ""
		}
	}
	Header["PlatesQty"] = CurrentString

	fmt.Println(Header)
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
