package gols

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Content struct {
	Name   string
	Isfile bool
}

func FindCurrentPath() string {

	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return ex

}

func GetPathContent(path string) ([]Content, error) {
	var con []Content
	files, err := ioutil.ReadDir(path)
	for _, file := range files {
		con = append(con, Content{Name: file.Name(), Isfile: !file.IsDir()})
	}
	return con, err
}

func PrintPathContent(c []Content, b bool) {
	if b {
		fmt.Printf("Files: \n")
		for _, con := range c {
			if con.Isfile {
				fmt.Printf(" %s\n", con.Name)
			}

		}
	} else {
		fmt.Printf("\nFolders: \n")
		for _, con := range c {
			if !con.Isfile {
				fmt.Printf(" %s\n", con.Name)
			}

		}
	}
}
