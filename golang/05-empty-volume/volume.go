package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

var mounted = func(file string) string { return path.Join("mnt", file) }

type hello struct {
	Message string `json:"message"`
}

func main() {
	var err error
	// read-write files
	for i := 0; i < 2; i++ {
		readdir("mnt", i)
		// on 1st try, non-existing, should err
		err = readtxt(mounted("hello.txt"))
		if err != nil {
			fmt.Println(err)
		}
		// write file
		err = writetxt(mounted("hello.txt"), "hi")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func writetxt(src, content string) error {
	return ioutil.WriteFile(src, []byte(content), 0644)
}

func readtxt(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	fmt.Printf("reading %s:\n%s\n", src, string(b))
	return nil
}

func readdir(dir string, run int) {
	fmt.Printf("run #%d\n%s\n", run, dir)
	fis, _ := ioutil.ReadDir(dir)
	for _, fi := range fis {
		fmt.Printf("|-- %s\n", fi.Name())
	}
}
