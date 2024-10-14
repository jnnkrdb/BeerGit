package main

import (
	"log"
	"os"

	"github.com/jnnkrdb/BeerGit/version"
)

func main() {

	// init the log output
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.LUTC | log.Ldate)
	log.SetOutput(os.Stdout)

	log.Printf("Hello World! - Version: %s", version.CurrentVersion)

}
