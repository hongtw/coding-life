// https://app.codesignal.com/challenge/uf5m9HooxrJQoYZMn
/*
Reference:
	* Read file
	https://yourbasic.org/golang/list-files-in-directory/
	https://flaviocopes.com/go-list-files/
	https://golangbot.com/read-files/

	* Sort
	https://golang.org/pkg/sort/#Slice
	https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
	https://stackoverflow.com/questions/23330781/sort-go-map-values-by-keys

	* Write file
	https://turreta.com/2019/07/30/golang-write-text-to-file-line-by-line/

	* Stringer
	https://golang.org/pkg/fmt/#Stringer
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// TextDir : Shoud rewrite it to "/root/devops" when submitting to codesignal
var TextDir string = "./root/devops"

func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// SubjectManger has subject as key and point of Students as value
type SubjectManger map[string]Students

// SubjectStudents has subject name and students
type SubjectStudents struct {
	Name     string
	Students Students
}

func (sm *SubjectManger) feed(ss *SubjectStudents) {
	subName := ss.Name
	students := ss.Students
	(*sm)[subName] = append((*sm)[subName], students...)
}

func (sm *SubjectManger) sortAll() (keys []string) {
	for k, students := range *sm {
		students.sortByScore()
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

// Student contains student's basic information
type Student struct {
	Name  string
	Score int
}

func (s Student) String() string {
	// https: //golang.org/pkg/fmt/#Stringer
	return s.Name
}

// Students contains many student
type Students []Student

func (sts *Students) sortByScore() {
	// descending
	sort.Slice(*sts, func(i, j int) bool { return (*sts)[i].Score > (*sts)[j].Score })
}

// Head return Top n students
func (sts Students) Head(count int) Students {
	if leng := len(sts); count > leng {
		count = leng
	}
	return sts[:count]
}

func (sts Students) String() string {
	// https: //golang.org/pkg/fmt/#Stringer
	output := ""
	for _, student := range sts {
		output += student.Name + "\n"
	}
	return output
}

// ImportStudents import students and its score from file
func ImportStudents(filepath string) (students Students) {
	// Read a file line by line: https://golangbot.com/read-files/
	f, err := os.Open(filepath)
	check(err)
	defer func() { err = f.Close(); check(err) }()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		studentName := fields[0] + " " + fields[1]
		score, _ := strconv.Atoi(fields[2])
		students = append(students, Student{studentName, score})
	}
	return students
}

func visit(count *int, ch chan *SubjectStudents) filepath.WalkFunc {
	// https://flaviocopes.com/go-list-files/
	return func(path string, info os.FileInfo, err error) error {
		check(err)
		if !info.IsDir() {
			*count++
			go func() {
				subName := strings.TrimSuffix(info.Name(), ".txt")
				students := ImportStudents(path)
				ch <- &SubjectStudents{subName, students}
			}()
		}
		return nil
	}
}

func main() {
	sm := SubjectManger{}
	count := 0
	ch := make(chan *SubjectStudents)
	filepath.Walk(TextDir, visit(&count, ch))
	for i := 0; i < count; i++ {
		sm.feed(<-ch)
	}
	sortedSubjects := sm.sortAll()
	output := ""
	for _, subject := range sortedSubjects {
		topStudents := sm[subject].Head(4)
		output += fmt.Sprintf("%v:\n%v\n", subject, topStudents)
	}

	fmt.Println(output)
}
