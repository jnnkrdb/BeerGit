package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
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

	// create the list of repositories, which need to be pulled
	var RepositoriesToPullList []string

	for _, subFolder := range entries {
		log.Printf("subFolder: %s ", subFolder.Name())

		// make dir checks
		if err = checkDir(subFolder); err != nil {
			log.Printf("\tSKIPPED - errormsg: %v\n", err)
			continue
		}

		// read subdirectories of the current directory
		var resDir string = fmt.Sprintf("%s\\%s", workdir, subFolder.Name())
		if err = readGitProject(resDir); err != nil {
			log.Printf("\tSKIPPED - errormsg: %v\n", err)
			continue
		}

		// change dir and run git pull
		log.Printf("change dir and run git pull\n")
		RepositoriesToPullList = append(RepositoriesToPullList, resDir)

	}

	log.Printf("\n\n\nstarting git pull session")
	for i := range RepositoriesToPullList {
		gitPull(RepositoriesToPullList[i])
	}
}

// ---------------------------------------------------------------------- additional functions
func checkDir(dEntry fs.DirEntry) error {
	log.Printf("checking folder: %s\n", dEntry.Name())

	switch {
	case !dEntry.IsDir():
		return fmt.Errorf("dEntry is no directory")

	case dEntry.Name() == ".git":
		return fmt.Errorf("dEntry is '.git'-Directory")
	}
	return nil
}

func readGitProject(path string) error {
	log.Printf("checking directory: %s\n", path)

	var err error
	var subEntries []fs.DirEntry
	if subEntries, err = os.ReadDir(path); err != nil {
		return err
	}

	for i := range subEntries {
		if subEntries[i].Name() == ".git" {
			return nil
		}
	}

	return fmt.Errorf("this is not a .git Project directory")
}

func gitPull(dir string) {
	log.Printf("pulling repos-dir: %s\n", dir)

	if err := os.Chdir(dir); err != nil {
		log.Printf("error changing dir: %v\n", err)
		return
	}

	if stdout, err := exec.Command("git", "pull").Output(); err != nil {
		log.Printf("error pulling project: %v\n", err)
		return
	} else {
		log.Printf("pulling message: %s\n", string(stdout))
	}
}
