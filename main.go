package main

import (
	"csvparser/entities"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	sourceDir  string
	resultFile string
	limit      int
	idLimit    int
)

func initFlags() {
	flag.StringVar(&sourceDir, "source", "./source", "source")
	flag.StringVar(&resultFile, "result", "./result.csv", "result.csv")
	flag.IntVar(&limit, "limit", 1000, "1000")
	flag.IntVar(&idLimit, "maxObject", 20, "20")
	flag.Parse()
}

func main() {
	initFlags()

	resultArray := parse(limit, sourceDir)
	err := writeResult(resultArray, idLimit, resultFile)
	if err != nil {
		log.Fatalf("write error: %v\n", err)
	}
}

// parse files into array
func parse(limit int, sourceDir string) []entities.Place {
	var wg = &sync.WaitGroup{}
	var chanPlaces = make(chan entities.Place)
	var resultTree = entities.BTree{MaxHeight: limit}
	var resultArray []entities.Place

	fileList, err := getFileList(sourceDir)
	if err != nil {
		log.Fatalf("read dir error: %v\n", err)
	}

	for _, filePath := range fileList {
		wg.Add(1)
		go worker(filePath, chanPlaces, wg)
	}

	go func() {
		for data := range chanPlaces {
			resultTree.Insert(data)
		}
	}()

	wg.Wait()

	return resultTree.Root.ToArray(resultArray)
}

func writeResult(resultArray []entities.Place, idLimit int, resultFile string) error {
	var placeIDCounter = make(map[int]int)

	file, err := os.Create(resultFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, data := range resultArray {
		placeIDCounter[data.Id]++
		if placeIDCounter[data.Id] > idLimit {
			continue
		}
		err := writer.Write([]string{
			strconv.Itoa(data.Id),
			data.Name,
			data.Condition,
			data.State,
			data.Price,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func worker(filePath string, chanPlaces chan entities.Place, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file error: %v\n", err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("read csv error: %v\n", err)
	}

	for _, line := range lines {
		tmp := strings.Split(line[0], ";")
		price, err := strconv.Atoi(strings.TrimSuffix(tmp[4], "RUB"))
		if err != nil {
			log.Fatalf("parse price error: %v\n", err)
		}
		id, err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		data := entities.Place{
			Id:        id,
			Name:      tmp[1],
			Condition: tmp[2],
			State:     tmp[3],
			Price:     tmp[4],
			PriceInt:  price,
		}
		chanPlaces <- data
	}
}

func getFileList(sourceDir string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), "csv") {
			fileList = append(fileList, path)
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	return fileList, nil
}
