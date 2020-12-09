// https://app.codesignal.com/challenge/CGrwxb2b4JdojE2gj
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
	"sync"
)

// IPPattern is regex for valid IP
var IPPattern = regexp.MustCompile(`\b((1\d\d|2[0-4]\d|25[0-5]|\d\d|\d)\.){3}(1\d\d|2[0-4]\d|25[0-5]|\d\d|\d)\b`)

func visitAllFiles(root string) <-chan []string {
	ch := make(chan []string)
	go func() {
		var wg sync.WaitGroup
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			wg.Add(1)
			go func() {
				data, _ := ioutil.ReadFile(path)
				ch <- IPPattern.FindAllString(string(data), -1)
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
	cache := make(map[string]bool)
	allUniqueIPs := []string{}
	for IPs := range ch {
		for _, IP := range IPs {
			if _, isSeen := cache[IP]; isSeen {
				continue
			}
			cache[IP] = true
			allUniqueIPs = append(allUniqueIPs, IP)
		}
	}
	sort.Strings(allUniqueIPs)
	for _, ip := range allUniqueIPs {
		fmt.Println(ip)
	}
}
