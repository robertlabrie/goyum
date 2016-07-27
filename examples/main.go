package main

import (
	"fmt"

	"github.com/robertlabrie/goyum"
)

func main() {
	goyum.DBPath = "/Users/robertlabrie/tmp/yum/yumdb"
	fmt.Println("bubbles")
	bont := goyum.GetPackages()
	fmt.Printf("%+v", bont)

	var users map[string]string
	users = goyum.GetUsers()
	fmt.Printf("%+v\n\n", users)

	dirs := goyum.GetPackageDirs()
	fmt.Printf("%+v", dirs)

}
