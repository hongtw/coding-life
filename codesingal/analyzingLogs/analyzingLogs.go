// https://app.codesignal.com/challenge/egRo5er4TmikYj5qB
/*
Reference:
	https://github.com/tmrts/go-patterns/blob/master/concurrency/parallelism.go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type targetFile struct {
	name string
	num  int
}

func (tf targetFile) String() string {
	return fmt.Sprintf("%v %v", tf.name, tf.num)
}

type Counter map[string]int

func (c *Counter) getSorted() (tfs []targetFile) {
	for filename, num := range *c {
		tfs = append(tfs, targetFile{filename, num})
	}
	sort.Slice(tfs, func(i, j int) bool {
		return tfs[i].num > tfs[j].num ||
			(tfs[i].num == tfs[j].num && tfs[i].name < tfs[j].name)
	})
	return tfs
}

// ErrorPattern is regex for error log
var ErrorPattern = regexp.MustCompile(`\d+\|ERROR\|([\w.-]+)\|`)

func visitAllFiles(root string) <-chan []string {
	ch := make(chan []string)
	go func() {
		var wg sync.WaitGroup
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(path, ".log") {
				return nil
			}
			wg.Add(1)
			go func() {
				data, _ := ioutil.ReadFile(path)
				for _, matched := range ErrorPattern.
					FindAllStringSubmatch(string(data), -1) {
					ch <- matched
				}
				wg.Done()
			}()
			return nil
		})

		go func() {
			wg.Wait()
			close(ch)
		}()
	}()
	return ch
}

func main() {
	ch := visitAllFiles("./root")
	counter := Counter{}
	for matched := range ch {
		counter[matched[1]]++
	}
	for _, file := range counter.getSorted() {
		fmt.Println(file)
	}
}
