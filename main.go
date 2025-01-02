package main

import (
	"log"
	"os"

	goMacho "github.com/blacktop/go-macho"
)

// func GetFileName(path string) string {
// 	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
// 	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
// }

func convertCafToMidi_Desktop(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := goMacho.NewFile(f)
	if err != nil {
		panic(err)
	}

	return m.FileTOC.String()
}

func main() {
	if len(os.Args) == 2 {
		convertCafToMidi_Desktop(os.Args[1:2][0])
	} else {
		log.Fatal("please drag a file")
	}
}
