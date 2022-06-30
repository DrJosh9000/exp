package exp // import "github.com/DrJosh9000/exp"

import (
	"bufio"
	"log"
	"os"
)

// MustForEachLineIn calls cb with each line in the file.
// It uses a bufio.Scanner internally, which can fail on longer lines.
// If an error is encountered, it calls log.Fatal.
func MustForEachLineIn(path string, cb func(line string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("MustForEachLineIn opening file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		cb(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("MustForEachLineIn: scanner: %v", err)
	}
}