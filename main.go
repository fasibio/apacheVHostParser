package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Values struct {
	URL        string
	DOCKERPORT string
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		fmt.Println("Usage: Command {URL} {dockerport}")
		return
	}
	ex, err := os.Executable()
	exPath := filepath.Dir(ex)
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadFile(exPath + "/VHostTemplate.conf")
	if err != nil {
		panic(err)
	}
	file := string(buf)
	data := Values{
		URL:        os.Args[1],
		DOCKERPORT: os.Args[2],
	}
	file = strings.Replace(file, "{{.URL}}", data.URL, -1)
	file = strings.Replace(file, "{{.DOCKERPORT}}", data.DOCKERPORT, -1)
	newFilename := strings.Replace(data.URL, ".", "_", -1)
	newFilename = newFilename + ".conf"
	f, err := os.Create(newFilename)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(newFilename)
	w.Flush()

}
