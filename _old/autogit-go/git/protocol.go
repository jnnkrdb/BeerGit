package git

import (
	"fmt"
	"os"
	"time"
)

var (
	LOG     *bool
	CONSOLE *bool
)

const (
	_LOGDIR  = "Logs"
	_LOGFILE = _LOGDIR + "/log.txt"
)

func Init(_log *bool, _console *bool) {

	LOG = _log
	CONSOLE = _console
}

func Write(_file string, _func string, _msg string) {

	m := `{ "date": "` + time.Now().Format(time.RFC3339) + `", "src": "` + _file + `", "func": "` + _func + `", "msg": "` + _msg + `" }`

	if *CONSOLE {

		fmt.Println(m)
	}

	if *LOG {

		if _, err := os.Stat(_LOGFILE); os.IsNotExist(err) {

			err = os.Mkdir(_LOGDIR, 0755)

			if err != nil {
				panic(err)
			}

			time.Sleep(5 * time.Second)

			_, err = os.Create(_LOGFILE)

			if err != nil {
				panic(err)
			}
		}

		_protocol, err := os.OpenFile(_LOGFILE, os.O_APPEND, 0644)

		if err != nil {

			panic(err)
		}

		defer _protocol.Close()

		_protocol.WriteString(m + "\n")

		_protocol.Sync()
	}
}
