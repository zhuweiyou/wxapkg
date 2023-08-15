package unpacker

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

type UnpackFile struct {
	NameLen uint32
	Name    string
	Offset  uint32
	Size    uint32
}

const DefaultUnpackTo = "_unpack"

func Unpack(from string) error {
	f, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	root := filepath.Dir(from)
	name := filepath.Base(from) + DefaultUnpackTo

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
		if err := binary.Read(f, binary.BigEndian, &data.NameLen); err != nil {
			return fmt.Errorf("error reading NameLen: %v", err)
		}

		nameBytes := make([]byte, data.NameLen)
		if _, err := f.Read(nameBytes); err != nil {
			return fmt.Errorf("error reading Name: %v", err)
		}
		data.Name = string(nameBytes)

		if err := binary.Read(f, binary.BigEndian, &data.Offset); err != nil {
			return fmt.Errorf("error reading Offset: %v", err)
		}

		if err := binary.Read(f, binary.BigEndian, &data.Size); err != nil {
			return fmt.Errorf("error reading Size: %v", err)
		}

		fmt.Println("readFile =", data.Name, "at Offset =", data.Offset)

		fileList = append(fileList, data)
	}

	for _, d := range fileList {
		d.Name = "/" + name + d.Name
		path := filepath.Join(root, filepath.Dir(d.Name))

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}

		w, err := os.Create(filepath.Join(root, d.Name))
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}

		_, err = f.Seek(int64(d.Offset), 0)
		if err != nil {
			w.Close()
			return fmt.Errorf("error seeking to Offset: %v", err)
		}

		buf := make([]byte, d.Size)
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
		fmt.Println("writeFile =", filepath.Join(root, d.Name))
	}

	return nil
}
