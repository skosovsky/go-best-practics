package entrance

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
)

func GetValue() (valueOne, valueTwo, valueThree int) {
	valueOne, valueTwo, valueThree, existValues := getFromPath()
	if existValues == true {
		return
	} else {
		valueOne, valueTwo, valueThree, existValues = getFromFile()
		if existValues == true {
			return
		} else {
			valueOne, valueTwo, valueThree, existValues = getFromFlag()
			if existValues == true {
				return
			} else {
				valueOne, valueTwo, valueThree = 5, 7, 10
				return
			}
		}
	}

}

func getFromPath() (valueOne, valueTwo, valueThree int, existValues bool) {
	valueOneString, existOne := os.LookupEnv("VALUE_ONE")
	valueTwoString, existTwo := os.LookupEnv("VALUE_TWO")
	valueThreeString, existThree := os.LookupEnv("VALUE_THREE")

	if existOne == true && existTwo == true && existThree == true {
		valueOne, err := strconv.Atoi(valueOneString)
		if err != nil {
			return 0, 0, 0, false
		}

		valueTwo, err2 := strconv.Atoi(valueTwoString)
		if err2 != nil {
			return 0, 0, 0, false
		}

		valueThree, err3 := strconv.Atoi(valueThreeString)
		if err3 != nil {
			return 0, 0, 0, false
		}

		return valueOne, valueTwo, valueThree, true
	}

	return 0, 0, 0, false
}

func getFromFile() (valueOne, valueTwo, valueThree int, existValues bool) {
	type ContentFromFile struct {
		ValueOne   int `json:"value_one"`
		ValueTwo   int `json:"value_two"`
		ValueThree int `json:"value_three"`
	}

	var myContent ContentFromFile

	contentFile, err := os.ReadFile("config.json")
	if err != nil {
		return 0, 0, 0, false
	}

	err = json.Unmarshal(contentFile, &myContent)
	if err != nil {
		return 0, 0, 0, false
	}

	return myContent.ValueOne, myContent.ValueTwo, myContent.ValueThree, true

}

func getFromFlag() (valueOne, valueTwo, valueThree int, existValues bool) {
	flag.IntVar(&valueOne, "value-one", 0, "contents value one")
	flag.IntVar(&valueTwo, "value-two", 0, "contents value two")
	flag.IntVar(&valueThree, "value-three", 0, "contents value three")

	flag.Parse()

	checkFlag := 0
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			checkFlag++
		}
	})

	if checkFlag == 0 {
		return 0, 0, 0, false
	} else {
		return valueOne, valueTwo, valueThree, true
	}

}
