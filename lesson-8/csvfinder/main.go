package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type contentFromFile struct {
	PathToFile string `json:"path_to_file"`
	ValueTwo   int    `json:"value_two"`
	ValueThree int    `json:"value_three"`
}

func main() {
	getExPath()
	getLastCommit()

	pathToFile, _, _ := getSettings()
	operator := [5]string{">=", "<=", "=", ">", "<"}
	//logical := [2]string{"AND", "OR"}
	userInput := readInput(operator)
	fmt.Printf("%#v", userInput)

	file, err := os.Open(pathToFile)
	if err != nil {
		log.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = 0
	reader.Comment = '#'

	for {
		record, err := reader.Read()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(record)
	}
}

func getSettings() (string, int, int) {
	var myContent contentFromFile

	contentFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		return "", 0, 0
	}

	err = json.Unmarshal(contentFile, &myContent)
	if err != nil {
		log.Println(err)
		return "", 0, 0
	}

	return myContent.PathToFile, myContent.ValueTwo, myContent.ValueThree
}

func readInput(operator [5]string) []string {
	var vals []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		val := scanner.Text()
		for _, v := range operator {
			if strings.HasPrefix(val, v) {
				val = strings.Replace(val, v, v+" ", 1)
				break
			} else if strings.HasSuffix(val, v) {
				val = strings.Replace(val, v, " "+v, 1)
				break
			}
		}
		if strings.Contains(val, " ") {
			vals = append(vals, strings.Split(val, " ")...)
		} else {
			vals = append(vals, val)
		}
	}
	return vals
}

func getExPath() {
	exPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current paht:", exPath)
}

func getLastCommit() {
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	commitHash := string(out)
	fmt.Println("Last commit:", commitHash)
}
