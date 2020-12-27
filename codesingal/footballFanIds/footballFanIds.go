// https://app.codesignal.com/challenge/L9YzWMLMJbwSTGRkk

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func readFile(path string) []string {
	data, _ := ioutil.ReadFile(path)
	return strings.Fields(strings.ReplaceAll(
		string(data), ".", string(os.PathSeparator)))
}

func main() {
	root := "/root/devops/"
	invites := readFile(root + "invite.info")
	bans := readFile(root + "ban.info")
	invitePat := regexp.MustCompile(
		fmt.Sprintf(`/(%v).*\.info$`, strings.Join(invites, "|")))
	banPat := regexp.MustCompile(
		fmt.Sprintf(`/(%v)/\w*\.info$`, strings.Join(bans, "|")))

	fanIds := make([]string, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if (!banPat.MatchString(path)) && invitePat.MatchString(path) {
			fanIds = append(fanIds, readFile(path)...)
		}
		return nil
	})
	sort.Strings(fanIds)
	for _, fanId := range fanIds {
		fmt.Println(fanId)
	}
}
