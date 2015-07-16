package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Song provides convenient file-format-independent storage for the
// information about each song.
type Song struct {
	Title    string
	Filename string
	Seconds  int
}

// Reads an arbitrary .m3u music playlist file given on the command line and
// outputs an equivalent .pls playlist file.
// $ ./m3u2pls Bowie-Singles.m3u [> Bowie-Singles.pls]
func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	}
}

// Accepts the entire contents of an .m3u file as a single string and returns a slice
// of the songs it is able to parse from the string.
func readM3uPlaylist(data string) (songs []Song) {
	var song Song

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}

	return songs
}

// parse lines of the form: #EXTINF:duration,title and where the duration is expected
// to be an integer, either -1 or greater than zero.
func parseExtinfLine(line string) (title string, seconds int) {
	// Find the position of the first digit or the minus sign
	// i holds the position of the first digit of the duration (or of -)
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:] // after this, the line has the form: duration,title

		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]

			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n", title, err)
				seconds = -1
			}
		}
	}

	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")

	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}

	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}
