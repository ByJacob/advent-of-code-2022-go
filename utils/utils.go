package utils

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadFile(fileName string) string {

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(data[:])

}

func GetDataPerDay(day int, isTest bool) string {
	var fileName string
	testFilePart := ""
	if isTest {
		testFilePart = "_test"
	}
	if day < 10 {
		fileName = "day0" + strconv.Itoa(day) + testFilePart + ".txt"
	} else {
		fileName = "day" + strconv.Itoa(day) + testFilePart + ".txt"
	}
	absoluteFileName, _ := filepath.Abs("./data/" + fileName)
	return ReadFile(absoluteFileName)
}

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
