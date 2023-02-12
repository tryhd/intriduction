package helpers

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FileUpload(fileNameBase, data string) (path string) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		panic("InvalidImage")
	}
	ImageType := data[11:idx]
	log.Println(ImageType)

	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)

	switch ImageType {
	case "png":
		path = filepath.Join("app", "upload", "photo", fileNameBase+".png")
		im, err := png.Decode(r)
		if err != nil {
			panic("Bad png")
		}

		f, err := os.Create(path)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, im)
		defer f.Close()
	case "jpeg":
		path = filepath.Join("app", "upload", "photo", fileNameBase+".jpeg")
		im, err := jpeg.Decode(r)
		if err != nil {
			panic("Bad jpeg")
		}

		f, err := os.Create(path)
		if err != nil {
			panic("Cannot open file")
		}

		jpeg.Encode(f, im, nil)
	case "jpg":
		path = filepath.Join("app", "upload", "photo", fileNameBase+".jpg")

		im, err := gif.Decode(r)
		if err != nil {
			panic("Bad jpg")
		}

		f, err := os.Create(path)
		if err != nil {
			panic("Cannot open file")
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		path = filepath.Join("app", "upload", "photo", fileNameBase+".gif")

		im, err := gif.Decode(r)
		if err != nil {
			panic("Bad gif")
		}

		f, err := os.Create(path)
		if err != nil {
			panic("Cannot open file")
		}

		gif.Encode(f, im, nil)
	}
	return path
}

func ImageBase(file string) string {
	imgFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	ext := strings.ReplaceAll(filepath.Ext(file), ".", "")
	res := "data:image/" + ext + ";base64," + imgBase64Str
	return res
}
