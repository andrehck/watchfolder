package main

import (
	"os"

	"gitlab.globoi.com/andre.luis/spy-bond/internal/watch"
)

var (
	pathSrc    = os.Getenv("PATHSRC")
	pathDst    = os.Getenv("PATHDST")
	pathDstGcp = os.Getenv("PATHDSTGCP")
	nameProj   = os.Getenv("PROJECTGCP")
)

func main() {
	watch.FolderWatch(pathSrc, pathDst, pathDstGcp)
}
