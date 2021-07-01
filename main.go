package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		if files && !out.IsDir() {
			return nil
		}
		dropTheLine(path, path1, out, files)
		return nil
	})
	if test != nil {
		fmt.Println(test.Error())
	}
	return nil
}

func searchLatestFile(str string, arg bool) string{
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
		if file.IsDir() && arg {
			nameFiles = file.Name()
			continue
		}
	}

	return nameFiles
}

func dropTheLine(pathOs string, str string, out os.FileInfo, arg bool) {
	var newPath string
	pathSplit := strings.Split(strings.Replace(str, pathOs,"", -1), string(os.PathSeparator))
	for er, rt := range pathSplit {
		if pathOs == str {
			fmt.Println("Directory tree:")
			break
		}
		if er == len(pathSplit) -1 && searchLatestFile(str, arg) == rt {
			newPath += "└───" + out.Name()
			break
		} else if er == len(pathSplit) -1 && searchLatestFile(str, arg) != rt {
			newPath += "├───" + out.Name()
			break
		}
		if searchLatestFile(pathOs, arg) == rt {
			newPath += "\t"
		} else {
			newPath += "|\t"
		}
		pathOs += rt + "\\"
	}
	switch {
	case out.IsDir():
	case out.Size() == 0:
		newPath += " (empty)"
	default:
		newPath += " (" + strconv.Itoa(int(out.Size())) +"b)"
	}
	fmt.Println(newPath)
}
