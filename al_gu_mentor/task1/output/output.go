package output

import (
	"errors"
	"log"
	"os"
)

func SaveToNewFile(fileName string, values ...string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(errors.New("problem with create file"))
	}
	defer f.Close()

	for _, n := range values {
		_, err = f.WriteString(n + "\n")
		if err != nil {
			log.Println(errors.New("problem with save file"))
		}
	}
}
