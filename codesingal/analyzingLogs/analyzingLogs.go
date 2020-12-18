// https://app.codesignal.com/challenge/egRo5er4TmikYj5qB
/*
Reference:
	https://github.com/tmrts/go-patterns/blob/master/concurrency/parallelism.go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"sync"
)

type Counter map[string]int

func (c *Counter) getSorted() []string {
	gps := make([]string, 0, len(*c))
	for grepped := range *c {
		gps = append(gps, grepped)
	}
	sort.Slice(gps, func(i, j int) bool {
		return (*c)[gps[i]] > (*c)[gps[j]] ||
			((*c)[gps[i]] == (*c)[gps[j]] && gps[i] < gps[j])
	})
	return gps
}

// ErrorPattern is regex for error log
var ErrorPattern = regexp.MustCompile(`\d+\|ERROR\|([\w.-]+)\|`)

func grepError(log string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, _ := ioutil.ReadFile(log)
	for _, matched := range ErrorPattern.
		FindAllStringSubmatch(string(data), -1) {
		ch <- matched[1]
	}
}

func main() {
	allLogs, _ := filepath.Glob("./root/devops/logs/*.log")
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(allLogs))
	for _, log := range allLogs {
		go grepError(log, ch, &wg)
	}
	go func() { wg.Wait(); close(ch) }()

	counter := Counter{}
	for grepped := range ch {
		counter[grepped]++
	}
	for _, file := range counter.getSorted() {
		fmt.Println(file, counter[file])
	}
}
