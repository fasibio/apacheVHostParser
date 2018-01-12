package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	moveFileToApacheConfig(newFilename)
	doApacheConfigTest(newFilename)
	restartApache()
}

func moveFileToApacheConfig(configFile string) {
	fmt.Println("moveFileToApache sites-enabled")
	err := os.Rename(configFile, "/etc/apache2/sites-enabled/"+configFile)
	if err != nil {
		panic(err)
	}

}

func doApacheConfigTest(configFile string) {
	fmt.Println("test apache config")
	args := []string{"configtest"}
	cmd := exec.Command("/usr/sbin/apachectl", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error by Configtest remove Config file from sites-enable")
		os.Remove("/etc/apache2/sites-enabled/" + configFile)
		panic(err)
	}
	// output, err2 := cmd.Output()
	// fmt.Println("hier", output)
	// if err2 != nil {
	// 	panic(err2)
	// }
}

func restartApache() {
	fmt.Println("Restart Apache2")
	args := []string{"restart", "apache2"}
	cmd := exec.Command("/bin/systemctl", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error by restart Apache2 Server")
		panic(err)
	}
}
