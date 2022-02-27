package roles

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
)

var ctx = context.Background()

func CopyFile(files string, pathSrc string, pathDst string, pathDstGcp string) {

	if len(files) > 0 {

		sourceFile, err := os.Open(pathSrc + files)
		if err != nil {
			log.Fatal(err)
		}

		defer sourceFile.Close()

		switch os.Getenv("SITEENV") {
		case "GCP":
			CopyFileGcp(files, pathDstGcp, pathSrc)
		case "DCCM":
			destFile, err := os.Create(pathDst + files)

			if err != nil {
				log.Fatal(err)
			}
			defer destFile.Close()

			bytesCopied, err := io.Copy(destFile, sourceFile)

			if err != nil {
				log.Fatal(err)
			} else {
				os.Remove(pathSrc + files)
			}

			log.Printf("Copied %d bytes of file %s", bytesCopied, files)

		}

	}
}

func CopyFileGcp(file string, pathDstGcp string, pathSrc string) {
	fmt.Println(pathDstGcp, file)
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := gcp.NewHTTPClient(
		gcp.DefaultTransport(),
		gcp.CredentialsTokenSource(creds))
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := gcsblob.OpenBucket(ctx, client, pathDstGcp, nil)

	if err != nil {
		log.Fatalln(err)
	}

	defer bucket.Close()

	w, err := bucket.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Copied file %s", file)
		os.Remove(pathSrc + file)
	}

	closeErr := w.Close()

	if closeErr != nil {
		log.Fatal(closeErr)
	} else {
		log.Default()
	}
}
