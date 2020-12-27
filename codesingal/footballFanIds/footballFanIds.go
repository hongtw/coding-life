// https://app.codesignal.com/challenge/L9YzWMLMJbwSTGRkk

// Repetition (greedy and non-greedy) https://yourbasic.org/golang/regexp-cheat-sheet/#repetition-greedy-and-non-greedy
// Regexp	Meaning
// x*	    zero or more x, prefer more
// x*?	    prefer fewer (non-greedy)
// x+	    one or more x, prefer more
// x+?	    prefer fewer (non-greedy)
// x?	    zero or one x, prefer one
// x??	    prefer zero
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
	pattern := regexp.MustCompile(fmt.Sprintf(
		`/(%v)?((%v).*)??/\w+\.info$`,
		strings.Join(bans, "|"),
		strings.Join(invites, "|"),
	))

	fanIds := make([]string, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		matched := pattern.FindStringSubmatch(path)
		if len(matched) > 0 && matched[3] != "" {
			fanIds = append(fanIds, readFile(path)...)
		}
		return nil
	})
	sort.Strings(fanIds)
	for _, fanId := range fanIds {
		fmt.Println(fanId)
	}
}
