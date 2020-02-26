package sort

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/mitinarseny/HSEProgTechLab1/students"
	"gopkg.in/yaml.v2"
)

var dataFlag = flag.String("data", "", "")

var data sort.Interface

func TestMain(m *testing.M) {
	flag.Parse()
	if *dataFlag != "" {
		var err error
		data, err = getData()
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	} else {

	}
	os.Exit(m.Run())
}

func getData() (sort.Interface, error) {
	f, err := os.Open(*dataFlag)
	if err != nil {
		return nil, err
	}
	var data []students.Student
	if err := yaml.NewDecoder(f).Decode(&data); err != nil {
		return nil, fmt.Errorf("unable to parse YAML: %s", err)
	}
	return students.Order(data, students.ByTotalPoints, students.ByFullName, students.ByFaculty, students.BySpeciality), nil
}
