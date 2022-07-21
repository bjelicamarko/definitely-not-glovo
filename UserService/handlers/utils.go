package handlers

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

func AdjustResponseHeaderJson(resWriter *http.ResponseWriter) {
	(*resWriter).Header().Set("Content-Type", "application/json")
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetB64Image(image string) string {
	bytes, err := ioutil.ReadFile(image)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding
}
