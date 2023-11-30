package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Spécifiez le chemin du répertoire que vous souhaitez inspecter
	directoryPath := "/"

	// Obtenez la liste des fichiers dans le répertoire
	fileList, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	// Parcourez la liste des fichiers et imprimez les noms
	for _, file := range fileList {
		fmt.Println(file.Name())
	}
}
