package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jnnkrdb/corerdb/fnc"
)

// get the config
func getRepoConfig(file string) (r []PullRequest, e error) {
	if e = fnc.LoadStructFromFile("json", file, &r); e != nil {
		e = fmt.Errorf("error loading config from [%s]: %#v", file, e)
	}
	return
}

// create the actual config object, which will be used for the routine execution
type PullRequest struct {
	Base   string `json:"base"`
	Link   string `json:"link"`
	Branch string `json:"branch"`
}

func (pr PullRequest) Clone() (e error) {

	log.Printf("Base: %s\n", pr.Base)

	dir := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pr.Link, pr.Base, ""), "/", "_"), ".git", "")

	if e = os.Mkdir(dir, 0777); e == nil {
		var f *os.File
		if f, e = os.Create(fmt.Sprintf("./%s/temp.txt", dir)); e == nil {

			var currDir string
			if currDir, e = os.Getwd(); e == nil {
				defer os.Chdir(currDir)

				if e = os.Chdir(dir); e == nil {

					log.Printf("dir: %s\n", dir)
					log.Printf("Link: %s\n", pr.Link)
					log.Printf("Branch: %s\n", pr.Branch)
					log.Printf("File: %#v\n", f.Name())
				}
			}
		}
	}

	return e
}
