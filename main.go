package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.LUTC | log.Ldate)
	log.SetOutput(os.Stdout)

	var err error

	// get working dir
	var workdir string
	if workdir, err = os.Getwd(); err != nil {
		log.Fatalln(err)
	}
	log.Printf("workdir: %s\n", workdir)

	// get all subdirectories in this directory
	var entries []fs.DirEntry
	if entries, err = os.ReadDir("./"); err != nil {
		log.Fatalln(err)
	}

	for i := range entries {
		if entries[i].IsDir() && !(entries[i].Name() == ".git") {

			var dirEntries []fs.DirEntry
			if dirEntries, err = os.ReadDir(fmt.Sprintf("%s\\%s", workdir, entries[i].Name())); err != nil {
				log.Printf("%#v\n", err)
				continue
			}

			for ii := range dirEntries {
				if dirEntries[ii].IsDir() {
					log.Printf("dir: %s\\%s\n", entries[i].Name(), dirEntries[ii].Name())
				}
			}
		}
	}
}
