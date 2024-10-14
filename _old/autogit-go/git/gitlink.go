package git

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/toast.v1"
	"gopkg.in/yaml.v3"
)

const (
	_CODEFILE = "gitlink.go"
)

func GitUpdate(_yamlconfig string) {
	var __func = "GitUpdate(_yamlconfig string)"

	Write(_CODEFILE, __func, "Initialize Config-Yaml")

	_git, ready := _readReposFromYAML(_yamlconfig)

	if ready {

		for u := 0; u < len(_git.Collection); u++ {

			_r := _git.Collection[u]

			Write(_CODEFILE, __func, "Collection: "+_r.Collection)

			for i := 0; i < len(_r.Repos); i++ {

				os.Chdir(_r.Directory + _r.Repos[i].Path)

				Write(_CODEFILE, __func, "git.exe pull - "+_r.Repos[i].URL+" | "+_r.Repos[i].Branch)

				cmd := exec.Command("git.exe", "pull")

				cmd.Run()
			}
		}

		_notify("AutoGit-GO", "AutoGit-GO", "--- SUCCESS ---")

	} else {

		_notify("AutoGit-GO", "AutoGit-GO", "--- FAILURE ---")

	}
}

func _readReposFromYAML(_file string) (Git, bool) {

	yamlfile, err := os.Open(_file)

	var gitrepos Git

	if err != nil {

		return gitrepos, false

	} else {

		defer yamlfile.Close()

		yamlInByte, _ := ioutil.ReadAll(yamlfile)

		yaml.Unmarshal(yamlInByte, &gitrepos)

		return gitrepos, true
	}
}

func _notify(_appid string, _title string, _msg string) {

	notification := toast.Notification{
		AppID:   _appid,
		Title:   _title,
		Message: _msg,
	}

	err := notification.Push()

	if err != nil {

		log.Fatalln(err)
	}
}
