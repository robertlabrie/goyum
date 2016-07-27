package goyum

import (
	"bufio"
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
	InstalledOn   string
}

var YumPath = "/var/lib/yum"
var DBPath = "/var/lib/yum/yumdb"

func GetPackageInfo(path string) (yumpackage YumPackage) {

	return
}

func GetPackageDirs() (dirs []string) {
	dirs = []string{}
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
	//users[0] = "root"
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
