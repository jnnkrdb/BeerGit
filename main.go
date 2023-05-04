package main

import (
	"log"
	"os"
)

const (
	_CONFIG             = "git.json"
	_MAX_SIMULTANE_PROC = 5
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.LUTC | log.Ldate)
	log.SetOutput(os.Stdout)

	if c, err := getRepoConfig(_CONFIG); err != nil {
		log.Fatalln(err)
	} else {
		for i := range c {
			log.Printf("BEGIN --- [%s]\n", c[i].Link)
			if err = c[i].Clone(); err != nil {
				log.Printf("error processing repo: %#v\n", err)
			}
			log.Printf("END --- [%s]\n\n", c[i].Link)
		}
	}
}
