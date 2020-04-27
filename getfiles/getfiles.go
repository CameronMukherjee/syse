package getfiles

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gookit/color"
)

// Directory struct
type Directory struct {
	Directory string
	Files     []string
	SizeKB    int64
}

// File struct
type File struct {
	Name   string
	SizeKB int64
}

// StartScan starts scanning files in parameter directory.
func StartScan(directory string) {
	files, err := ioutil.ReadDir(directory)
	if (err) != nil {
		color.Red.Println(err)
	}
	for _, f := range files {

		currentFolder := directory + "\\" + f.Name()
		fi, err := os.Stat(currentFolder)
		if err != nil {
			color.Red.Println(err)
		}

		switch mode := fi.Mode(); {

		case mode.IsDir():
			extension := filepath.Ext(directory)
			if extension == ".exe" {
				color.Red.Println(directory)
			}
			StartScan(currentFolder)

		case mode.IsRegular():
			color.Blue.Println(currentFolder)
			name := string(f.Name())
			initFile(name, getFileSize(currentFolder))
		}
	}
}

func initFile(filename string, filesize int64) {

}

func initDir(directory string, files []File, directorySize int64) {

}

func getFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

// JSONAddToFile adds struct to JSON file.
func JSONAddToFile(input interface{}) {
	var jsonData []byte
	jsonData, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	data := string(jsonData)
	// data = data + "," + "\n"
	data = data + ","
	// f, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f, err := os.OpenFile("output.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Fatal(err)
	}
}
