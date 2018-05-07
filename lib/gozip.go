package gozip

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"strings"
)

func CompressString(raw string) string {
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

func DecompressString(raw string) string {
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
