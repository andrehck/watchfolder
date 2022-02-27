package roles

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func CopyFolder(file fs.FileInfo, pathSrc string, pathDst string, pathDstGcp string) {

	stats := file.Mode()

	files := file.Name()
	source := pathSrc + files
	dest := pathDst + files

	err := os.MkdirAll(dest, stats)
	if err != nil {
		log.Fatal(err)
	} else {
		countFiles, err := ioutil.ReadDir(source)
		if err != nil {
			log.Fatal(err)
		} else {
			if len(countFiles) > 0 {
			}
			os.Remove(source)
		}
	}
}
