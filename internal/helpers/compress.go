package helpers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func compress(filename string) error {

	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(filename + ".gz")

	defer out.Close()

	gzout := gzip.NewWriter(out)

	_, err = io.Copy(gzout, in)

	return err

}

//Zip handles individual zipping of files
func Zip(filename []string, jsonOpt bool) {

	type zippedFiles struct {
		Files []string `json:"Compressed Files"`
	}

	var wg sync.WaitGroup

	var file string
	var compressedFiles []string

	for _, file = range filename {
		wg.Add(1)
		go func(filename string, jsonOpt bool) {

			if jsonOpt {
				if err := compress(filename); err != nil {
					log.Println(err)
				}
				//won't print any additional info out so we can have a
				//clean output in case someone wants to pipe the info to a file.
				compressedFiles = append(compressedFiles, filename+".gz")
			} else {
				fmt.Printf("Compressing %s\n", filename)
				if err := compress(filename); err != nil {
					log.Println(err)
				}
				fmt.Printf("Compressed %s\n", filename)
				compressedFiles = append(compressedFiles, filename+".gz")
			}
			wg.Done()
		}(file, jsonOpt)
	}

	wg.Wait()

	if jsonOpt {

		var jsonData []byte
		jsonData, err := json.MarshalIndent(zippedFiles{compressedFiles}, "", " ")
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(jsonData))
	} else {

		fmt.Printf("Compressed %d files\n", len(compressedFiles))
		fmt.Println("--- Compressed Files ---")
		fmt.Println(strings.Join(compressedFiles, "\n"))
	}

}
