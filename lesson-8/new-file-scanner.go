package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	collections "github.com/golang-collections/collections/stack"
)

type File struct {
	size int64
	name string
}

type FileSystem interface {
	ReadDir(dirname string) ([]fs.FileInfo, error)
}

type OsSystem struct{}

func (s *OsSystem) ReadDir(dirname string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func main() {
	dir := flag.String("dir", "./", "Scan directory")
	deleteFile := flag.Bool("delete", false, "Delete duplicates flag")
	flag.Parse()
	fileSystem := new(OsSystem)
	duplicates := getDuplicateFiles(fileSystem, *dir)
	for _, file := range duplicates {
		fmt.Println(file)
		if *deleteFile {
			os.Remove(file)
		}
	}
}

func getDuplicateFiles(fileSystem FileSystem, dir string) []string {
	files, _ := fileSystem.ReadDir(dir)
	stack := new(collections.Stack)
	duplicateList := make([]string, 0)
	fileList := make(map[File]struct{})

	for _, f := range files {
		if f.IsDir() {
			dirName := dir + string(os.PathSeparator) + f.Name()
			stack.Push(dirName)
		} else {
			fileList[File{name: f.Name(), size: f.Size()}] = struct{}{}
		}

	}

	for ok := true; ok; ok = stack.Len() > 0 {

		if stack.Len() == 0 {
			break
		}

		dirRoot := (stack.Pop()).(string)
		files, _ := fileSystem.ReadDir(dirRoot)

		for _, f := range files {
			dirName := dirRoot + string(os.PathSeparator) + f.Name()
			if f.IsDir() {
				stack.Push(dirName)
			} else {
				file := File{name: f.Name(), size: f.Size()}
				_, okDuplicate := fileList[file]
				if okDuplicate {
					duplicateList = append(duplicateList, dirName)
				} else {
					fileList[file] = struct{}{}
				}
			}
		}
	}

	return duplicateList
}
