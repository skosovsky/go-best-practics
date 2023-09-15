package entrance

import (
	"encoding/json"
	"flag"
	"os"
)

func GetValue() (valueOne, valueTwo, valueThree string) {
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
				valueOne, valueTwo, valueThree = "5", "7", "10"
				return
			}
		}
	}

}

func getFromPath() (valueOne, valueTwo, valueThree string, existValues bool) {
	valueOne, existOne := os.LookupEnv("VALUE_ONE")
	valueTwo, existTwo := os.LookupEnv("VALUE_TWO")
	valueThree, existThree := os.LookupEnv("VALUE_THREE")

	if existOne == true && existTwo == true && existThree == true {
		existValues = true
	} else {
		existValues = false
	}

	return
}

func getFromFile() (valueOne, valueTwo, valueThree string, existValues bool) {
	type ContentFromFile struct {
		ValueOne   string `json:"value_one"`
		ValueTwo   string `json:"value_two"`
		ValueThree string `json:"value_three"`
	}

	var myContent ContentFromFile

	contentFile, err := os.ReadFile("config.json")
	if err != nil {
		return "", "", "", false
	}

	err = json.Unmarshal(contentFile, &myContent)
	if err != nil {
		return "", "", "", false
	}

	return myContent.ValueOne, myContent.ValueTwo, myContent.ValueThree, true

}

func getFromFlag() (valueOne, valueTwo, valueThree string, existValues bool) {
	flag.StringVar(&valueOne, "value-one", "", "contents value one")
	flag.StringVar(&valueTwo, "value-two", "", "contents value two")
	flag.StringVar(&valueThree, "value-three", "", "contents value three")

	flag.Parse()

	checkFlag := 0
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			checkFlag++
		}
	})

	if checkFlag == 0 {
		return valueOne, valueTwo, valueThree, true
	} else {
		return "", "", "", false
	}

}
