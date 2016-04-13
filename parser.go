package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := "data" // Aloways look in 'data'
	files := LoadFiles(path)
	data := ExtractData(path, files)
	ExportData(data)
}

func LoadFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Print(err)
	}
	return files
}

func ExtractData(path string, files []os.FileInfo) string {
	var data = "nothin"

	// For each os.FileInfo
	for _, v := range files {
		// Open
		filename := path + "/" + v.Name()
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return data
		}

		// For each line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

		// Close the file
		file.Close()

	}

	return data
}

func ExportData(data string) {
	fmt.Println("Exporting Data")
	fmt.Println(data)
}
