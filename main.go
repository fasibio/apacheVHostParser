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

	fmt.Println("lets go", exPath)
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
	f, err := os.Create(newFilename + ".conf")
	w := bufio.NewWriter(f)
	n4, err := w.WriteString(file)
	fmt.Println("Generate VHost file", n4)
	w.Flush()

}
