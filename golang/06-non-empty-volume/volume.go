package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		err = readtxt(mounted("hello.txt"))
		if err != nil {
			log.Fatal(err)
		}
		// rewrite file
		err = writetxt(mounted("hi.txt"), "hi")
		if err != nil {
			fmt.Printf("error writing hi.txt: %s\n", err.Error())
		}
		if i == 0 {
			// append file
			err = appendtxt(mounted("hello.txt"), "hi")
			if err != nil {
				fmt.Printf("error appending hello.txt: %s", err.Error())
			}
		}
	}
}

func writetxt(src, content string) error {
	return ioutil.WriteFile(src, []byte(content), 0644)
}

func appendtxt(src, content string) error {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
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
