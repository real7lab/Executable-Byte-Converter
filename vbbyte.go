package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the file path of the executable: ")
	var exePath string
	fmt.Scanln(&exePath)

	bytes, err := ioutil.ReadFile(exePath)
	if err != nil {
		return
	}

	var numericValues []byte
	for _, b := range bytes {
		numericValues = append(numericValues, strconv.Itoa(int(b))...)
	}

	fileName := strings.TrimSuffix(filepath.Base(exePath), filepath.Ext(exePath))
	executablePath, err := os.Executable()
	if err != nil {
		return
	}
	softwareDir := filepath.Dir(executablePath)

	outputFilePath := filepath.Join(softwareDir, fileName+".txt")
	err = ioutil.WriteFile(outputFilePath, numericValues, 0644)
	if err != nil {
		return
	}

	fmt.Println("Numeric byte values from the .exe file successfully written to", outputFilePath)
}
