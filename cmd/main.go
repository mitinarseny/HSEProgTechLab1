package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mitinarseny/HSEProgTechLab1/sort"
	"github.com/mitinarseny/HSEProgTechLab1/students"
)

// Flags
const (
	dataFlag = "data"
)

// Exit codes
const (
	_ = iota
	errExitCode
)

func main() {
	if err := run(); err != nil {
		switch err.(type) {
		default:
			fmt.Printf("error ocured: %s\n", err)
			os.Exit(errExitCode)
		}
	}
}

func run() error {
	data := flag.String(dataFlag, "", "path to JSON")
	flag.Parse()

	if *data == "" {
		return RequiredFlagIsEmpty(dataFlag)
	}

	f, err := os.Open(*data)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	var a []students.Student

	if err := json.NewDecoder(f).Decode(&a); err != nil {
		return fmt.Errorf("unable to parse JSON: %w", err)
	}

	sort.SelectSort(students.Order(a, students.ByFullName))

	for i, s := range a {
		fmt.Printf("%2d: %s\n", i, &s)
	}

	return nil
}
