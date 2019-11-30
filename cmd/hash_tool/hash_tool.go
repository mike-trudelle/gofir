package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var folderName string
var fileName string
var fileOutput string
var algorithm string

func main() {

	flag.StringVar(&folderName, "folder", "", "a folder path")
	flag.StringVar(&fileName, "file", "", "a file path")
	flag.StringVar(&fileOutput, "output", "", "a folder path to dump data to")
	flag.StringVar(&algorithm, "algorithm", "", "hash algorithm to use (MD5/SHA1/SHA256)")

	flag.Parse()

	//open output file
	o, err := os.OpenFile(fileOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		//log library stuff
	}

	defer o.Close() //wait to close file

	//write the csv header data
	if _, err := o.WriteString("FileName,HashValue \r\n"); err != nil {
		//log library stuff
	}

	var files []string

	if folderName != "" {

		err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				//fmt.Printf("Skipping %v because it's a directory \r\n\r\n", info.Name())
			} else {

				files = append(files, path)
				return nil
			}
			return nil
		})

	} else {

		files = append(files, fileName)
	}

	//for the list of files, open, create hash value and write file name and hash value to file
	for _, file := range files {

		v := hash.getHashValue(file, algorithm)

		data := file + "," + v + "\r\n"

		if _, err := o.WriteString(data); err != nil {
			fmt.Println("Error Writing Hash to File")
			//log something
		}
	}
}
