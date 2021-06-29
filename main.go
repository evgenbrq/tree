package main

import (
	"fmt"
	"log"
	"strings"
	"path/filepath"
	"os"
)


func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTreee(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTreee(out *os.File, path string, files bool) error {
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	test := filepath.Walk(path, func(path1 string, out os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dropTheLine(path, path1, out, files)
		return nil
	})
	if test != nil {
		fmt.Println(test.Error())
	}
	return nil
}

func searchLatestFile(str string) string {
	var num int
	var nameFiles string
	path := filepath.Dir(str)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files{
		if len(files)-1 != num {
			num++
		} else {
			nameFiles = file.Name()
		}
	}
	return nameFiles
}

func dropTheLine(pathOs string, str string, out os.FileInfo, arg bool) {
	var as = "."
	var newPath, size string
	path := strings.Split(strings.ReplaceAll(str, pathOs+"\\",""), string(os.PathSeparator))
	//fmt.Println(path, pathOs, str)
	//dir, file := filepath.Split(str)
	//if len(path) > 1 {
	//	path = path[len(path)-1:]
	//}

	//fmt.Println("dir: ", dir, " file: ", file, " len: ", len(dir))
	for i:=0; i<len(path)-1;i++ {
		as = as + "\\" + path[i]
		if path[i] == searchLatestFile(as) {
			newPath = newPath + "\t"
		} else {
			newPath = newPath + " |\t"
		}
	}
	test := searchLatestFile(str)
	if out.Size() == 0 {
		size = "(empty)"
	} else {
		size = fmt.Sprint("(", out.Size(), "b)")
	}
	if arg {
		if test == path[len(path)-1] {
			if out.IsDir() {
				fmt.Println(newPath + " └───" + path[len(path)-1])
			}
		} else {
			if out.IsDir() {
				fmt.Println(newPath + " ├───" + path[len(path)-1])
			}
		}
	} else if !arg {
		if test == path[len(path)-1] {
			if out.IsDir() {
				fmt.Println(newPath + " └───" + path[len(path)-1])
			} else {
				fmt.Println(newPath + " └───" + path[len(path)-1], size)
			}
		} else {
			if out.IsDir() {
				fmt.Println(newPath + " ├───" + path[len(path)-1])
			} else {
				fmt.Println(newPath + " ├───" + path[len(path)-1], size)
			}
		}
	}
}
