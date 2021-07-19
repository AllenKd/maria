package cleaner

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"maria/internal/pkg/configs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type cleaner struct {
	RootFolder string
	targetFile string
}

var (
	once     sync.Once
	instance *cleaner

	filePrefix = time.Now().Format("20060102")
)

func New() *cleaner {
	once.Do(func() {
		instance = &cleaner{
			RootFolder: configs.New().Cleaner.RootFolder,
		}
		log.Debug("cleaner initialized")
	})
	return instance
}

func (c cleaner) Clean() {

	for _, file := range c.listFiles(c.RootFolder) {
		if file.IsDir() {
			lFile, err := c.largest(fmt.Sprintf("%s/%s", c.RootFolder, file.Name()))
			if err != nil {
				log.Error("fail to find largest file from: ", file.Name())
				continue
			}
			log.Info(lFile)
			if err := c.moveOut(lFile); err != nil {
				log.Error("fail to move file out: ", err.Error())
				return
			}

			if err := os.RemoveAll(fmt.Sprintf("%s/%s", c.RootFolder, file.Name())); err != nil {
				log.Error("fail to remove folder: ", file.Name(), ". ", err.Error())
				return
			}
		}
	}
}

func (c cleaner) moveOut(filePath string) (err error) {
	log.Debug("move file out: ", filePath)
	fileName := fmt.Sprintf("%s_%s", filePrefix, filepath.Base(filePath))
	err = os.Rename(filePath, fmt.Sprintf("%s/%s", c.RootFolder, fileName))
	return
}

func (c cleaner) listFiles(path string) (files []fs.FileInfo) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c cleaner) largest(folder string) (largestFile string, err error) {
	log.Debug("get largest from: ", folder)

	var f os.FileInfo
	err = filepath.Walk(folder,
		func(path string, file os.FileInfo, err error) error {
			log.Debug("walk file: ", path)
			if err != nil {
				log.Error(err.Error())
				return err
			}
			if f == nil {
				f = file
				return nil
			}
			if f.Size() < file.Size() {
				f = file
				largestFile = path
			}
			return nil
		},
	)

	//var file fs.FileInfo
	//var mfd string
	//var fd string
	//for _, f := range c.listFiles(folder) {
	//	if f.IsDir() {
	//		fd = fmt.Sprintf("%s/%s", folder, f.Name())
	//		ff, _ := os.Stat(c.largest(fd))
	//
	//		if ff != nil {
	//			f = ff
	//		} else {
	//			continue
	//		}
	//	}
	//	if file == nil {
	//		file = f
	//
	//		continue
	//	}
	//	if file.Size() < f.Size() {
	//		file = f
	//		mfd = fd
	//	}
	//}
	//
	//if file != nil {
	//	log.Debug("got largest file from: ", mfd, ", ", file.Name(), ". size: ", file.Size(), " bytes")
	//	largestFile = fmt.Sprintf("%s/%s", mfd, file.Name())
	//} else {
	//	log.Debug("no largest file from: ", folder)
	//}
	return
}
