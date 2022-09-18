package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	m, err := NewModel(inputPath)
	if err != nil {
		return err
	}
	log.Printf("parsed model, got %d functions", len(m.Functions))
	fnbytes, _ := json.MarshalIndent(m.Functions, "", "  ")
	fmt.Println(string(fnbytes))
	err = Generate(m, outputPath)
	if err != nil {
		return err
	}

	err = GenerateToDo(m, outputPath)
	if err != nil {
		return err
	}

	return nil
}
