package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
)

type File struct {
	size int64
	name string
}

var fileList sync.Map

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger = logger.With(zap.String("goos", runtime.GOOS))

	dir := flag.String("dir", "./", "Scan directory")
	deleteFile := flag.Bool("delete", false, "Delete duplicates flag")
	flag.Parse()

	workers := make(chan struct{}, runtime.NumCPU()*2)
	result := scanner(workers, *dir, *deleteFile)
	<-result
}

func scanner(workers chan struct{}, dir string, deleteFile bool) chan struct{} {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, _ := config.Build()
	logger = logger.With(zap.String("goos", runtime.GOOS))
	defer logger.Sync()

	workers <- struct{}{}
	files, err := ioutil.ReadDir(dir)
	chanels := make([]chan struct{}, 0, len(files))
	result := make(chan struct{})

	if err != nil {
		logger.Fatal(err.Error())
	}

	go func() {
		defer func() {
			logger.With(
				zap.String("time", time.Now().String()),
				zap.String("dir", dir),
			).Debug("thread finished")
			<-workers
			result <- struct{}{}
		}()

		for _, f := range files {
			if f.IsDir() {
				dirName := dir + string(os.PathSeparator) + f.Name()
				logger.With(
					zap.String("time", time.Now().String()),
					zap.String("dir", dirName),
				).Debug("start new thread")
				chanels = append(chanels, scanner(workers, dirName, deleteFile))

			} else {
				file := File{name: f.Name(), size: f.Size()}
				_, ok := fileList.Load(file)
				if ok {
					filename := dir + string(os.PathSeparator) + f.Name()
					if deleteFile {
						err = os.Remove(filename)
						if err == nil {
							logger.With(
								zap.String("time", time.Now().String()),
								zap.String("file", filename),
							).Info("file deleted")
						} else {
							logger.With(
								zap.String("time", time.Now().String()),
								zap.String("file", filename),
							).Warn("file is not deleted")
						}
					}
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
