package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"duplicates/internal/filecrcalt"
	"duplicates/pkg/filecrc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
)

type File struct {
	fileName  string
	filePath  string
	fileMD5   string
	fileSize  int64
	fileCheck bool
	fileDel   bool
}

func main() {
	listFiles := make(map[string]File)
	var wg sync.WaitGroup

	log.Info().Msg("start service")
	pathToDir, needDelete, debug := getPath()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().Msg("scan files")
	wg.Add(1)
	go getFiles(pathToDir, listFiles, &wg)
	wg.Wait()

	log.Info().Msg("start find duplicates")
	findDups(listFiles, &wg)

	log.Info().Msg("start del duplicates")
	delDups(listFiles, *needDelete)

	log.Info().Msg("finish service")
}

// Функция получает параметры из командной строки или назначает их по-умолчанию
func getPath() (pathToDir string, needDelete, debug *bool) {
	flag.StringVar(&pathToDir, "path", "/Users/skosovsky/Downloads/Test", "path to dir")
	needDelete = flag.Bool("D", false, "need delete")
	debug = flag.Bool("debug", false, "need debug")
	flag.Parse()

	return pathToDir, needDelete, debug
}

// Функция обходит все папки и под папки и собирает первичные данные о файлах
func getFiles(pathToDir string, listFiles map[string]File, wg *sync.WaitGroup) {
	var m sync.Mutex
	defer wg.Done()

	dir, err := os.Open(pathToDir)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by open files")
	}
	defer func(dir *os.File) {
		err = dir.Close()
		if err != nil {
			log.Error().Stack().Err(err).Msg("error by close file")
		}
	}(dir)

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by read files")
	}

	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go getFiles(fmt.Sprint(pathToDir, "/", file.Name()), listFiles, wg)
		} else {
			log.Info().Msg("current file" + file.Name())
			m.Lock()
			currFile := File{file.Name(), pathToDir + "/", "", file.Size(), false, false}
			listFiles[fmt.Sprint(pathToDir, "/", file.Name())] = currFile
			m.Unlock()

			// Cleat output for visual
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				log.Error().Stack().Err(err).Msg("error by clean Stdout")
			}
		}
	}
}

// Функция ищет дубликаты по размеру файлов, при совпадении сравнивает их контрольную сумму
func findDups(listFilesMap map[string]File, wg *sync.WaitGroup) {
	var m sync.Mutex
	appFs := afero.NewOsFs() // test
	log.Debug().Stack().Msg("star func findDups")

	for key, val := range listFilesMap {
		for key2, val2 := range listFilesMap {
			if val.fileSize == val2.fileSize && key != key2 {
				wg.Add(2)

				go func() {
					defer wg.Add(-2) // wg.Done() * 2
					m.Lock()
					currFile := listFilesMap[key]
					currFile.fileMD5 = filecrc.FileMD5(key)
					fmt.Println(filecrcalt.FileMD5(key, appFs)) // test
					currFile.fileCheck = true
					listFilesMap[key] = currFile
					currFile2 := listFilesMap[key2]
					currFile2.fileMD5 = filecrc.FileMD5(key2)
					fmt.Println(filecrcalt.FileMD5(key2, appFs)) // test
					listFilesMap[key2] = currFile2
					m.Unlock()
				}()

				wg.Wait()

				if listFilesMap[key].fileMD5 == listFilesMap[key2].fileMD5 && key != key2 && listFilesMap[key].fileMD5 != "" &&
					listFilesMap[key].fileCheck && !listFilesMap[key2].fileCheck {
					currFile2 := listFilesMap[key2]
					currFile2.fileCheck = true
					currFile2.fileDel = true
					listFilesMap[key2] = currFile2
				}
			}
		}
	}
}

// Функция удаляет или делает вид, что удаляет дубликаты файлов, один раз запрашивая подтверждение
func delDups(listFilesMap map[string]File, needDel bool) {
	log.Debug().Stack().Msg("star func delDups")

	count := 0
	confirmation := true
	var yes string

	for key, val := range listFilesMap {
		if val.fileDel {
			if needDel {
				if confirmation {
					fmt.Println("Are you sure you need to delete the files? Yes or No?")
					_, err := fmt.Scan(&yes)
					if err != nil {
						log.Error().Stack().Err(err).Msg("error by Scan")
					}

					if !strings.EqualFold(yes, "yes") {
						fmt.Println("Nothing delete")
						return
					}

					confirmation = false
				}
				err := os.Remove(key)
				if err != nil {
					fmt.Println("Error, not delete", key)
				} else {
					fmt.Println("Delete", key)
					count++
				}
			} else {
				fmt.Println("Need delete", key)
				count++
			}
		}
	}

	if count == 0 {
		fmt.Println("Nothing delete")
	} else {
		if needDel {
			fmt.Println("Delete", count, "files")
		} else {
			fmt.Println("Need delete", count, "files")
		}
	}
}
