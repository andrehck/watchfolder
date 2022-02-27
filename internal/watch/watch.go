package watch

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gitlab.globoi.com/andre.luis/spy-bond/internal/roles"
)

func FolderWatch(pathSrc string, pathDst string, pathDstGcp string) {
	for {
		files, err := ioutil.ReadDir(pathSrc)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {

			if file.IsDir() {

				log.Println("Watch Folder habilitado somente para copia de arquivos " + file.Name() + " é um diretorio. Removendo path ...")
				os.Remove(pathSrc + file.Name())
			} else {
				file := file.Name()
				WaitingCopy(file, pathSrc)
				roles.CopyFile(file, pathSrc, pathDst, pathDstGcp)
			}
		}
	}
}

func WaitingCopy(file string, pathSrc string) {

	statFile, err := os.Stat(pathSrc + file)

	if err != nil {
		fmt.Println(err)
	}

	lastModifiedTime := statFile.ModTime()

	timeNow := time.Now()
	count := 5

	copyThen := timeNow.Add(time.Duration(-count) * time.Second)

	if lastModifiedTime.Second() == copyThen.Second() {
		return
	} else {
		WaitingCopy(file, pathSrc)
	}
}
