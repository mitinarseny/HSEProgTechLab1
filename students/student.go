package students

import "fmt"

type Student struct {
	FullName    string `json:"fullName"`
	Faculty     string `json:"faculty"`
	Speciality  string `json:"speciality"`
	TotalPoints uint   `json:"totalPoints"`
}

func (s *Student) String() string {
	return fmt.Sprintf("name: %q, faculty: %q, speciality: %q, totalPoints: %d", s.FullName, s.Faculty, s.Speciality, s.TotalPoints)
}

type lessFunc func(s1, s2 *Student) bool

func ByFullName(s1, s2 *Student) bool {
	return s1.FullName < s2.FullName
}

func ByFaculty(s1, s2 *Student) bool {
	return s1.Faculty < s2.Faculty
}

func BySpeciality(s1, s2 *Student) bool {
	return s1.Speciality < s2.Speciality
}

func ByTotalPoints(s1, s2 *Student) bool {
	return s1.TotalPoints < s2.TotalPoints
}

func Order(students []Student, by ...lessFunc) *multiSorter {
	return &multiSorter{
		students: students,
		less:     by,
	}
}

type multiSorter struct {
	students []Student
	less     []lessFunc
}

func (s *multiSorter) Len() int {
	return len(s.students)
}

func (s *multiSorter) Less(i, j int) bool {
	p, q := &s.students[i], &s.students[j]
	for _, l := range s.less {
		if l(p, q) {
			return true
		}
	}
	return false
}

func (s *multiSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}
