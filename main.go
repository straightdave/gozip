package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	goziplib "github.com/straightdave/gozip/lib"
)

var de = flag.Bool("d", false, "to decompress")

func main() {
	flag.Parse()

	var c []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c = append(c, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	raw := strings.Join(c, "\n") // Note: support Linux only

	if *de {
		fmt.Println(goziplib.DecompressString(raw))
	} else {
		fmt.Println(goziplib.CompressString(raw))
	}
}
