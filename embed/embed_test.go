package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

//go:embed hello.txt
var hello string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(hello)
}

//go:embed wisuda.jpg
var wisuda []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("wisuda_new.jpg", wisuda, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/alpha.txt
//go:embed files/beta.txt
//go:embed files/first.txt

var files embed.FS

func TestMultipleFiles(t *testing.T) {
	file1, err := files.ReadFile("files/first.txt")
	file2, err := files.ReadFile("files/alpha.txt")
	file3, err := files.ReadFile("files/beta.txt")
	fmt.Println(string(file1))
	fmt.Println(string(file2))
	fmt.Println(string(file3))
	if err != nil {
		panic(err)
	}
}

//go:embed files/*.txt

var path embed.FS

func TestPathMatch(t *testing.T) {
	dirEnt, err := path.ReadDir("files")
	for _, entry := range dirEnt {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
	if err != nil {
		fmt.Println(err)
	}
}
