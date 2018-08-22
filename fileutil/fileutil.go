package fileutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//FileUtil This is admitedly kind of a cop out from a naming perspective
// as this gets more fleshed out hopefully a better name than
//"something"-util will make itself apparent
type FileUtil interface {
	GetFilesInDirectory(path string) []string
}

type FileUtility struct {
}

//method to check if path is a directory
func isDirectory(path string) (bool, error) {
	fd, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Error describing file")
	}
	switch mode := fd.Mode(); {
	case mode.IsDir():
		log.Println(path, " is a directory")
		return true, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, nil
}

//GetFilesInDirectory is a method to walk directory to get array of files
func (fileUtil FileUtility) GetFilesInDirectory(dirPath string) []string {
	fileList := []string{}
	filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		fmt.Println("PATH ==> " + path)
		log.Println("NAME ==> ", f.Name())

		isDirectory, err := isDirectory(path)
		log.Println("path is a directory is ", isDirectory)
		if err != nil {
			log.Println("error walking directory ", err)
			return err
		}
		if isDirectory {
			// Do nothing
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})
	return fileList
}
