package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"testing"

	mysort "github.com/mitinarseny/HSEProgTechLab1/sort"
	"github.com/mitinarseny/HSEProgTechLab1/students"
)

const usage = `Usage: %s FILE ...

FIlE: path to JSON file containing array of objects. Each object should be in the following format:

  %s
`

// Errors
type cliError error

// Exit codes
const (
	_ = iota
	errExitCode
	cliErrorExitCode
)

func main() {
	if err := run(); err != nil {
		switch err.(type) {
		case cliError:
			flag.Usage()
			fmt.Printf("error: %s\n", err)
			os.Exit(cliErrorExitCode)
		default:
			fmt.Printf("error: %s\n", err)
			os.Exit(errExitCode)
		}
	}
}

func init() {
	flag.Usage = func() {
		var buff bytes.Buffer
		e := json.NewEncoder(&buff)
		e.SetIndent("  ", "  ")
		_ = e.Encode(students.Student{
			FullName:    "Full Name",
			Faculty:     "Faculty",
			Speciality:  "Speciality",
			TotalPoints: 297,
		})
		_, _ = fmt.Fprintf(os.Stderr, usage, os.Args[0], buff.String())
	}
}

func run() error {
	flag.Parse()
	if flag.NArg() == 0 {
		return cliError(errors.New("no data was provided"))
	}

	data := make(map[int]sort.Interface, flag.NArg())
	for _, fname := range flag.Args() {
		f, err := os.Open(fname)
		if err != nil {
			return err
		}

		var s []students.Student

		if err := json.NewDecoder(f).Decode(&s); err != nil {
			return fmt.Errorf("unable to parse JSON: %w", err)
		}
		f.Close()
		data[len(s)] = students.Order(s,
			students.ByTotalPoints,
			students.ByFullName,
			students.ByFaculty,
			students.BySpeciality,
		)
	}

	testing.Init()
	for size, d := range data {
		fmt.Printf("SIZE: %d\n", size)
		fmt.Print("Select:\t")
		res := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mysort.Select(d)
			}
		})
		fmt.Printf("avg %d ns\n", res.NsPerOp())

		fmt.Print("Merge:\t")
		res = testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mysort.Merge(d)
			}
		})
		fmt.Printf("avg %d ns\n", res.NsPerOp())

		fmt.Print("Heap:\t")
		res = testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mysort.Heap(d)
			}
		})
		fmt.Printf("avg %d ns\n\n", res.NsPerOp())
	}
	return nil
}
