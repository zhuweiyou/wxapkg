package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

type UnpackFile struct {
	nameLen uint32
	name    string
	offset  uint32
	size    uint32
}

func Unpack(from string) error {
	f, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	root := filepath.Dir(from)
	name := filepath.Base(from) + "_unpack"

	var firstMark byte
	if err := binary.Read(f, binary.BigEndian, &firstMark); err != nil {
		return fmt.Errorf("error reading first header mark: %v", err)
	}
	fmt.Println("first header mark =", firstMark)

	var info1 uint32
	if err := binary.Read(f, binary.BigEndian, &info1); err != nil {
		return fmt.Errorf("error reading info1: %v", err)
	}
	fmt.Println("info1 =", info1)

	var indexInfoLength uint32
	if err := binary.Read(f, binary.BigEndian, &indexInfoLength); err != nil {
		return fmt.Errorf("error reading indexInfoLength: %v", err)
	}
	fmt.Println("indexInfoLength =", indexInfoLength)

	var bodyInfoLength uint32
	if err := binary.Read(f, binary.BigEndian, &bodyInfoLength); err != nil {
		return fmt.Errorf("error reading bodyInfoLength: %v", err)
	}
	fmt.Println("bodyInfoLength =", bodyInfoLength)

	var lastMark byte
	if err := binary.Read(f, binary.BigEndian, &lastMark); err != nil {
		return fmt.Errorf("error reading last header mark: %v", err)
	}
	fmt.Println("last header mark =", lastMark)

	if firstMark != 0xBE || lastMark != 0xED {
		return fmt.Errorf("it's not a wxapkg file")
	}

	var fileCount uint32
	if err := binary.Read(f, binary.BigEndian, &fileCount); err != nil {
		return fmt.Errorf("error reading fileCount: %v", err)
	}
	fmt.Println("fileCount =", fileCount)

	fileList := make([]UnpackFile, 0, fileCount)

	for i := uint32(0); i < fileCount; i++ {
		var data UnpackFile
		if err := binary.Read(f, binary.BigEndian, &data.nameLen); err != nil {
			return fmt.Errorf("error reading nameLen: %v", err)
		}

		nameBytes := make([]byte, data.nameLen)
		if _, err := f.Read(nameBytes); err != nil {
			return fmt.Errorf("error reading name: %v", err)
		}
		data.name = string(nameBytes)

		if err := binary.Read(f, binary.BigEndian, &data.offset); err != nil {
			return fmt.Errorf("error reading offset: %v", err)
		}

		if err := binary.Read(f, binary.BigEndian, &data.size); err != nil {
			return fmt.Errorf("error reading size: %v", err)
		}

		fmt.Println("readFile =", data.name, "at Offset =", data.offset)

		fileList = append(fileList, data)
	}

	for _, d := range fileList {
		d.name = "/" + name + d.name
		path := filepath.Join(root, filepath.Dir(d.name))

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}

		w, err := os.Create(filepath.Join(root, d.name))
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}

		_, err = f.Seek(int64(d.offset), 0)
		if err != nil {
			w.Close()
			return fmt.Errorf("error seeking to offset: %v", err)
		}

		buf := make([]byte, d.size)
		_, err = f.Read(buf)
		if err != nil {
			w.Close()
			return fmt.Errorf("Error reading from file: %v", err)
		}

		_, err = w.Write(buf)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}

		w.Close()
		fmt.Println("writeFile =", filepath.Join(root, d.name))
	}

	return nil
}
