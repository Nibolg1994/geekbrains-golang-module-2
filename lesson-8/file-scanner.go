package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
)

type File struct {
	size int64
	name string
}

var fileList sync.Map

func main() {
	dir := flag.String("dir", "./", "Scan directory")
	deleteFile := flag.Bool("delete", false, "Delete duplicates flag")
	flag.Parse()

	workers := make(chan struct{}, runtime.NumCPU()*2)
	result := scanner(workers, *dir, *deleteFile)
	<-result
}

func scanner(workers chan struct{}, dir string, deleteFile bool) chan struct{} {
	workers <- struct{}{}
	files, err := ioutil.ReadDir(dir)
	chanels := make([]chan struct{}, 0, len(files))
	result := make(chan struct{})

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer func() { result <- struct{}{} }()
		defer func() { <-workers }()

		for _, f := range files {
			if f.IsDir() {
				chanels = append(chanels, scanner(workers, dir+string(os.PathSeparator)+f.Name(), deleteFile))
			} else {
				file := File{name: f.Name(), size: f.Size()}
				_, ok := fileList.Load(file)
				if ok {
					err = os.Remove(dir + string(os.PathSeparator) + f.Name())
					fmt.Println(f.Name())
				} else {
					fileList.Store(file, struct{}{})
				}
			}
		}

		for _, ch := range chanels {
			<-ch
		}
	}()

	return result
}
