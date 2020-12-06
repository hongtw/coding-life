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

type Subject string
type Subjects map[string]Students

func (sb Subjects) sortAll() (index []string) {
	// https://stackoverflow.com/questions/23330781/sort-go-map-values-by-keys
	for k, students := range sb {
		students.sortByScore()
		index = append(index, k)
	}
	sort.Strings(index)
	return
}

type Student struct {
	Name  string
	Score int
}

func (s Student) String() string {
	// https: //golang.org/pkg/fmt/#Stringer
	return s.Name
}

type Students []Student

func (st *Students) sortByScore() {
	// descending
	sort.Slice(*st, func(i, j int) bool { return (*st)[i].Score > (*st)[j].Score })
}

func (sts Students) Head(count int) Students {
	// descending
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

func getStudents(filepath string) (students Students) {
	// Read a file line by line: https://golangbot.com/read-files/
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
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

func visit(subjectsAll *Subjects) filepath.WalkFunc {
	// https://flaviocopes.com/go-list-files/
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			subName := strings.TrimSuffix(info.Name(), ".txt")
			if _, isOK := (*subjectsAll)[subName]; !isOK {
				(*subjectsAll)[subName] = Students{}
			}
			(*subjectsAll)[subName] = append((*subjectsAll)[subName], getStudents(path)...)
		}
		return nil
	}
}

func main() {
	subjectsAll := Subjects{}
	filepath.Walk(TextDir, visit(&subjectsAll))
	sortedSubjects := subjectsAll.sortAll()
	output := ""
	for _, subject := range sortedSubjects {
		topStudents := subjectsAll[subject].Head(4)
		output += fmt.Sprintf("%v:\n%v\n", subject, topStudents)
	}

	fmt.Println(output)
}
