package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		fmt.Println(decompressString(raw))
	} else {
		fmt.Println(compressString(raw))
	}
}

func compressString(raw string) string {
	var buff bytes.Buffer
	gz := gzip.NewWriter(&buff)

	if _, err := gz.Write([]byte(raw)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buff.Bytes())
}

func decompressString(raw string) string {
	raw = strings.TrimSpace(raw)

	zippedStr, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		panic(err)
	}

	rdata := bytes.NewReader(zippedStr)
	upzippedStr, err := gzip.NewReader(rdata)
	if err != nil {
		panic(err)
	}

	res, err := ioutil.ReadAll(upzippedStr)
	if err != nil {
		panic(err)
	}
	return string(res)
}
