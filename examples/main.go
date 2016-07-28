package main

import (
	"fmt"

	"github.com/robertlabrie/goyum"
)

func main() {
	goyum.DBPath = "/Users/robertlabrie/tmp/yum/yumdb"
	goyum.DBPath = "/fake"
	fmt.Println("bubbles")
	bont := goyum.GetPackages()
	fmt.Printf("%+v", bont)

	var users map[string]string
	users = goyum.GetUsers()
	fmt.Printf("%+v\n\n", users)

	// dirs := goyum.GetPackageDirs()
	// fmt.Printf("%+v", dirs)

	p := goyum.GetPackageInfo("/Users/robertlabrie/tmp/yum/yumdb/z/b72c0834bd1cdf15d2486e8415ec61edf7c8086c-zlib-1.2.3-29.el6-x86_64")
	fmt.Printf("%+v", p)

	pkgs := goyum.ListInstalled()
	fmt.Printf("%+v", pkgs)
}
