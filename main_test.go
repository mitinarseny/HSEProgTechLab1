package HSEProgTechLab1

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"

	mysort "github.com/mitinarseny/HSEProgTechLab1/sort"
	"github.com/mitinarseny/HSEProgTechLab1/students"
)

var data []sort.Interface

func init() {
	log.SetFlags(0)
}

func Benchmark(b *testing.B) {
	for _, d := range data {
		b.Run(strconv.Itoa(d.Len()), func(b *testing.B) {
			b.Run("Merge", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					mysort.Merge(d)
				}
			})
			b.Run("Heap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					mysort.Heap(d)
				}
			})
			b.Run("Select", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					mysort.Select(d)
				}
			})
		})
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Print("no data was provided")
		os.Exit(1)
	}

	data = make([]sort.Interface, 0, flag.NArg())
	for _, fname := range flag.Args() {
		f, err := os.Open(fname)
		if err != nil {
			log.Print(err)
			continue
		}

		var s []students.Student

		if err := json.NewDecoder(f).Decode(&s); err != nil {
			log.Printf("unable to parse JSON: %s", err)
		}
		f.Close()
		data = append(data, students.Order(s,
			students.ByTotalPoints,
			students.ByFullName,
			students.ByFaculty,
			students.BySpeciality,
		))
	}
	os.Exit(m.Run())
}
