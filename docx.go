package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

type DocxDocument struct {
	Content []*DocFile
}

type DocFile struct {
	Root       *xmlquery.Node
	RawContent []byte
	FileName   string
}

func (o *DocxDocument) GetXMLDocFile(fileName string) *DocFile {
	for _, f := range o.Content {
		if f.FileName == fileName {
			return f
		}
	}
	return nil
}

func (o *DocxDocument) SaveToFile(filePath string) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	return o.save(file)

}

func (o *DocxDocument) Save() ([]byte, error) {
	buf := make([]byte, 0)
	writer := bytes.NewBuffer(buf)
	err := o.save(writer)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (o *DocxDocument) save(writer io.Writer) error {
	var fileWriter io.Writer
	var err error

	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	for _, doc := range o.Content {
		fileWriter, err = zipWriter.Create(doc.FileName)

		if err != nil {
			return err
		}

		if doc.IsXML() {
			_, err = fileWriter.Write([]byte(doc.Root.OutputXML(true)))
		} else {
			_, err = fileWriter.Write(doc.RawContent)
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (o *DocFile) IsXML() bool {
	return isXML(o.FileName)
}

//TODO Use reader instead path arg to support []byte/file
func ReadDocxFile(path string) (*DocxDocument, error) {
	var fReader io.ReadCloser
	var buff []byte

	docx, err := zip.OpenReader(path)
	defer docx.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*DocFile, 0)

	for _, f := range docx.File {
		fReader, err = f.Open()
		if err != nil {
			return nil, err
		}

		buff = make([]byte, f.FileInfo().Size())
		_, err = fReader.Read(buff)
		if err != nil {
			return nil, err
		}

		fmt.Println("Parsing file: " + f.Name)
		if isXML(f.Name) {
			rootNode, err := xmlquery.Parse(fReader)
			if err != nil {
				fmt.Println("Error in file2: " + f.Name + " - " + err.Error())
				if rootNode == nil {
					fmt.Println("Root is")
				}
			}
			result = append(result, &DocFile{Root: rootNode, FileName: f.Name, RawContent: nil})
		} else {
			result = append(result, &DocFile{Root: nil, FileName: f.Name, RawContent: buff})
		}
		fReader.Close()
	}
	return &DocxDocument{
		Content: result,
	}, err
}

func isXML(fileName string) bool {
	return strings.Contains(fileName, ".xml")
}
