package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// FunctionForSuffix ...
var FunctionForSuffix = map[string]func(string) ([]string, error){
	".gz": GzipFileList, ".tar": TarFileList, ".tar.gz": TarFileList,
	".tgz": TarFileList, ".zip": ZipFileList}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s archive1 [archive2 [... archiveN]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)

	}
	args := commandLineFiles(os.Args[1:])
	archiveFileList := ArchiveFileList
	if len(args[0]) == 1 && strings.IndexAny(args[0], "12345") != -1 {
		which := args[0][0]
		args = args[1:]
		switch which {
		case '2':
			archiveFileList = ArchiveFileList2
		case '3':
			archiveFileList = ArchiveFileList3
		case '4':
			archiveFileList = ArchiveFileList4
		case '5':
			archiveFileList = ArchiveFileListMap
		}
	}
	for _, filename := range args {
		fmt.Print(filename)
		lines, err := archiveFileList(filename)
		if err != nil {
			fmt.Println(" ERROR:", err)
		} else {
			fmt.Println()
			for _, line := range lines {
				fmt.Println(" ", line)
			}
		}
	}
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

// ArchiveFileList ...
func ArchiveFileList(file string) ([]string, error) {
	if suffix := Suffix(file); suffix == ".gz" {
		return GzipFileList(file)
	} else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".tgz" {
		return TarFileList(file)
	} else if suffix == ".zip" {
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

// ArchiveFileList2 ...
func ArchiveFileList2(file string) ([]string, error) {
	switch suffix := Suffix(file); suffix { // Naïve and noncanonical!
	case ".gz":
		return GzipFileList(file)
	case ".tar":
		fallthrough
	case ".tar.gz":
		fallthrough
	case ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

// ArchiveFileList3 ...
func ArchiveFileList3(file string) ([]string, error) {
	switch Suffix(file) {
	case ".gz":
		return GzipFileList(file)
	case ".tar":
		fallthrough
	case ".tar.gz":
		fallthrough
	case ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

// ArchiveFileList4 ...
func ArchiveFileList4(file string) ([]string, error) {
	switch Suffix(file) { // Canonical
	case ".gz":
		return GzipFileList(file)
	case ".tar", ".tar.gz", ".tgz":
		return TarFileList(file)
	case ".zip":
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

// ArchiveFileListMap ...
func ArchiveFileListMap(file string) ([]string, error) {
	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
		return function(file)
	}
	return nil, errors.New("unrecognized archive")
}

// Suffix ...
func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}

// ZipFileList ...
func ZipFileList(filename string) ([]string, error) {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()
	var files []string
	for _, file := range zipReader.File {
		files = append(files, file.Name)
	}
	return files, nil
}

// GzipFileList ...
func GzipFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return []string{gzipReader.Header.Name}, nil
}

// TarFileList ...
func TarFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var tarReader *tar.Reader
	if strings.HasSuffix(filename, ".gz") ||
		strings.HasSuffix(filename, ".tgz") {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		tarReader = tar.NewReader(gzipReader)
	} else {
		tarReader = tar.NewReader(reader)
	}
	var files []string
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return files, err
		}
		if header == nil {
			break
		}
		files = append(files, header.Name)
	}
	return files, nil
}