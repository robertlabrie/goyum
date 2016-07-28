package goyum

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type YumPackage struct {
	Name          string
	Repo          string
	InstalledBy   string
	Reason        string
	Releasever    string
	RepoTimeStamp string
	InstalledOn   int64
	CommandLine   string
}

var YumPath = "/var/lib/yum"
var DBPath = "/var/lib/yum/yumdb"

func ListInstalled() (installed []YumPackage) {
	dirs := GetPackageDirs()
	for _, d := range dirs {
		p := GetPackageInfo(d)
		installed = append(installed, p)
	}
	return
}

func GetPackageInfo(path string) (yumpackage YumPackage) {

	var data []byte
	var err error
	users := GetUsers()
	// extract the package name from the path -- this method sucks -- use RegEx
	name := path
	bits := strings.Split(name, "-")
	name = strings.Replace(name, bits[0]+"-", "", 1)
	yumpackage.Name = name

	//Linux doesn't have btime, so we use the mtime of from installed_by
	info, err := os.Stat(path + "/installed_by")
	yumpackage.InstalledOn = info.ModTime().Unix()

	// get installed by
	data, err = ioutil.ReadFile(path + "/installed_by")
	yumpackage.InstalledBy = string(data)

	// check if we had a user name for it
	if val, exists := users[yumpackage.InstalledBy]; exists {
		yumpackage.InstalledBy = val
	}

	//repo
	data, err = ioutil.ReadFile(path + "/from_repo")
	yumpackage.Repo = string(data)

	//reason
	data, err = ioutil.ReadFile(path + "/reason")
	yumpackage.Reason = string(data)

	//commandline
	data, err = ioutil.ReadFile(path + "/command_line")
	yumpackage.CommandLine = string(data)

	//releasever
	data, err = ioutil.ReadFile(path + "/releasever")
	yumpackage.Releasever = string(data)

	//from_repo_timestamp
	data, err = ioutil.ReadFile(path + "/from_repo_timestamp")
	yumpackage.RepoTimeStamp = string(data)

	if err != nil {
		return
	}
	return
}

func GetPackageDirs() (dirs []string) {
	dirs = []string{}

	if _, err := os.Stat(DBPath); os.IsNotExist(err) {
		return
	}
	err := filepath.Walk(DBPath, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() && len(path)-len(DBPath) > 3 {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
func GetPackages() (packages []YumPackage) {

	var p YumPackage

	packages = append(packages, p)

	return
}

func Foo() (out string) {
	return "bar"
}

func GetUsers() (users map[string]string) {
	users = make(map[string]string)

	file, err := os.Open("/etc/passwd")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), ":")
		if len(bits) < 3 {
			continue
		}
		users[bits[2]] = bits[0]
	}

	if err := scanner.Err(); err != nil {
		return
	}
	return
}
