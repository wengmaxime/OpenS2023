package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	directoryPath := "."

	fileList, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	// Parcourez la liste des fichiers et imprimez les noms
	for _, file := range fileList {
		fmt.Println(file.Name())
	}
}
