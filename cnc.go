package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var array []string
var Header map[string]string
var Details map[string]map[string]int
var DetailIndex int
var CutDepth int
var ZforMoving int

func main() {
	CutDepth = 3.0   //Enter deep
	ZforMoving = 5.0 //Safety Я for moving cuter to points
	Header = make(map[string]string)
	Details = make(map[string]map[string]int)
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
	DetailStopIndex()

	DetailStartIndex, _ := strconv.Atoi(Header["DetailStartIndex"])
	DetailStopIndex, _ := strconv.Atoi(Header["DetailStopIndex"])
	// for i := DetailStartIndex; i <= DetailStopIndex; i++ {
	// 	fmt.Println("details ", i, " ", array[i])
	// }
	DetailIndex = DetailStartIndex

	for DetailIndex < DetailStopIndex {
		MakeDetail(DetailIndex)

	}

	// printarray(Header)
	// printdetails(Details)
	// fmt.Println(Details)
	CallPlate()
}

func MakeGcode(X, Y, DetailLenght, DetailWidth float32) {
	var CutClear float32
	MaterialThickness, _ := strconv.Atoi(Header["MaterialThickness"])
	Steps := MaterialThickness / CutDepth
	CutThickness, _ := strconv.Atoi(Header["CutThickness"])
	CutClear = float32(CutThickness) / float32(2)
	//fmt.Println("MaterialThickness: ", MaterialThickness, " CutDepth: ", CutDepth, " Steps: ", Steps, " CutClear: ", CutClear)
	fmt.Print("G0Z", ZforMoving, "\n")
	fmt.Print("G0", "X", X-CutClear, "Y", Y-CutClear, "\n")
	for i := 1; i <= Steps; i++ {
		fmt.Print("G1Z", -1*i*CutDepth, "\n")
		fmt.Print("G1", "X", X-CutClear, "Y", Y+DetailLenght+CutClear, "\n")
		fmt.Print("G1", "X", X+DetailWidth+CutClear, "Y", Y+DetailLenght+CutClear, "\n")
		fmt.Print("G1", "X", X+DetailWidth+CutClear, "Y", Y-CutClear, "\n")
		fmt.Print("G1", "X", X-CutClear, "Y", Y-CutClear, "\n")
	}
	fmt.Print("G1Z", -1*MaterialThickness-1, "\n")
	fmt.Print("G1", "X", X-CutClear, "Y", Y+DetailLenght+CutClear, "\n")
	fmt.Print("G1", "X", X+DetailWidth+CutClear, "Y", Y+DetailLenght+CutClear, "\n")
	fmt.Print("G1", "X", X+DetailWidth+CutClear, "Y", Y-CutClear, "\n")
	fmt.Print("G1", "X", X-CutClear, "Y", Y-CutClear, "\n")
	fmt.Print("G0Z", ZforMoving, "\n")
}

func CallPlate() {
	PlatesQty, _ := strconv.Atoi(Header["PlatesQty"])
	for j := 1; j <= PlatesQty; j++ {
		for i, ii := range Details {
			if ii["PlateNumber"] == j {
				fmt.Println(";PlateNum :", j, " ", i, " - ", ii["X"], " ", ii["Y"], " ", ii["Rotate90"], " ", ii["DetailLenght"], " ", ii["DetailWidth"])
				if ii["Rotate90"] == 1 {
					MakeGcode(float32(ii["X"]), float32(ii["Y"]), float32(ii["DetailLenght"]), float32(ii["DetailWidth"]))
				} else {
					MakeGcode(float32(ii["X"]), float32(ii["Y"]), float32(ii["DetailWidth"]), float32(ii["DetailLenght"]))
				}
			}
		}
	}
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

	// 	for i, arrayslice := range array {
	// 		fmt.Println(i, " ", arrayslice)
	// 	}
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
	keys := make([]string, 0, len(array))
	for k := range array {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, i := range keys {
		fmt.Println(i)
		fmt.Print(" PlateNumber ", array[i]["PlateNumber"], " Rotate ", array[i]["Rotate90"], " DetailLenght ", array[i]["DetailLenght"], " DetailWidth ", array[i]["DetailWidth"], " X ", array[i]["X"], " Y ", array[i]["Y"], " KL ", array[i]["DetailKantLenght"], " KW ", array[i]["DetailKantWidth"], " GL ", array[i]["DetailGrooveLenght"], " GW ", array[i]["DetailGrooveWidth"])
		fmt.Println()
	}
}

// This function adds DetailFinishIndex line number to Header
func DetailStopIndex() {
	for i, ii := range array {
		if strings.Contains(ii, "<USnips>") {
			Header["DetailStopIndex"] = strconv.Itoa(i - 1)
		}
	}
}

func MakeDetail(Index int) {
	DetailNam := array[Index+2]
	DetailName := DetailNam[2:]
	DetailList := DetailsConvert(Index)
	DetailKant := DetailsConvert(Index + 1)
	// Details = map[string]map[string]int{
	// 	DetailName: {"DetailLenght": DetailList[0], "DetailWidth": DetailList[1], "DetailQty": DetailList[2]},
	// }
	DetailIndex += 3
	DetailNumber := 0
	DetailScanStop := DetailIndex + DetailList[2]
	for DetailIndex < DetailScanStop {
		DetailNumber++
		DetailProperty := DetailsConvert(DetailIndex)
		Rotate := 0
		PlateNumber := 0
		if DetailProperty[0] < 999 {
			Rotate = DetailProperty[0] / 10 % 10
			PlateNumber = DetailProperty[0] % 10
		} else {
			Rotate = DetailProperty[0] / 100 % 10
			PlateNumber = DetailProperty[0] % 100
		}
		Details[DetailName+" "+strconv.Itoa(DetailList[0])+" "+strconv.Itoa(DetailList[1])+" "+strconv.Itoa(DetailNumber)] = map[string]int{

			"DetailLenght":       DetailList[1],
			"DetailWidth":        DetailList[0],
			"DetailQty":          DetailList[2],
			"DetailNumber":       DetailNumber,
			"DetailKantLenght":   DetailKant[1],
			"DetailKantWidth":    DetailKant[0],
			"DetailGrooveLenght": DetailKant[3],
			"DetailGrooveWidth":  DetailKant[2],
			"Rotate90":           Rotate,
			"PlateNumber":        PlateNumber,
			"Y":                  DetailProperty[2],
			"X":                  DetailProperty[3],
			//добавить поворот
		}
		DetailIndex++
	}

}

func DetailsConvert(line int) []int {
	DetailString := make([]string, 4)
	DetailStringInt := make([]int, 4)
	key := 0
	for _, ii := range array[line] {
		if string(ii) != "X" && string(ii) != "_" && string(ii) != ">" {
			DetailString[key] += string(ii)
		} else {
			key++
		}
	}
	DetailStringInt[0], _ = strconv.Atoi(DetailString[0])
	DetailStringInt[1], _ = strconv.Atoi(DetailString[1])
	DetailStringInt[2], _ = strconv.Atoi(DetailString[2])
	DetailStringInt[3], _ = strconv.Atoi(DetailString[3])
	return DetailStringInt
}
