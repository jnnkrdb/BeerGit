package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.LUTC | log.Ldate)
	log.SetOutput(os.Stdout)

	if workdir, err := os.Getwd(); err != nil {
		log.Fatalln(err)
	} else {

		log.Printf("workdir: %s\n", workdir)

		if entries, err := os.ReadDir("./"); err != nil {
			log.Fatalln(err)
		} else {

			for _, e := range entries {

				if e.IsDir() {

					if e.Name() != ".git" {

						if dirEntries, err := os.ReadDir(fmt.Sprintf("%s\\%s", workdir, e.Name())); err != nil {
							log.Fatalln(err)
						} else {

							for _, i := range dirEntries {

								if i.IsDir() {

									log.Printf("dir: %s\\%s\n", e.Name(), i.Name())
								}
							}
						}
					}
				}
			}
		}
	}
}
