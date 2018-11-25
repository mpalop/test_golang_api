package tests

import (
	"io/ioutil"
	"os"
	"path"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return dat
}

func getFixture(fileName string) string {
	exec,err := os.Executable()
	if err == nil {
		finalPath := path.Join(path.Dir(exec), "../fixtures", fileName)
		return finalPath
	} else {
		panic(err)
	}

}