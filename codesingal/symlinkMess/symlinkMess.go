package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func isSymlink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink != 0
}

func main() {
	linkTo := map[string]string{}
	symlinks := []string{}

	filepath.Walk("/root/devops/", func(path string, info os.FileInfo, err error) error {
		if isSymlink(info) {
			symlinks = append(symlinks, path)
			dst, _ := os.Readlink(path)
			linkTo[path] = dst
		}
		return nil
	})

	sort.Strings(symlinks)
	for _, symlink := range symlinks {
		// get final destination
		dst := symlink
		for ; linkTo[dst] != ""; dst = linkTo[dst] {
		}

		// check whether dst's parent-dir is symlink
		// If so, resolve and recombine dst path
		dir, file := filepath.Split(dst)
		dir = strings.TrimSuffix(dir, "/") // To make `os.Lstat` get symlink flag or correct expected output
		if file == "" {
			dst = dir
		} else if dirinfo, _ := os.Lstat(dir); isSymlink(dirinfo) {
			dst = filepath.Join(linkTo[dir], file)
		}

		// output
		fmt.Printf("%v %v\n", symlink, dst)
	}
}
