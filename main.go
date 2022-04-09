package main

import (
	"fmt"
	"time"
)

func readXMLFile(path string) (*Node, error) {
	start := time.Now()
	root, err := ReadXMLFile(path)
	readingEnd := time.Now()

	if err != nil {
		fmt.Println(err.Error)
		return nil, err
	}

	readedIn := readingEnd.Sub(start)
	fmt.Printf("Readed file in: %d\n", readedIn.Milliseconds())
	return root, err
}

func readDocxFile(path string) {
	start := time.Now()
	doc, err := ReadDocxFile(path)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	err = doc.SaveToFile("./myFile.docx")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	end := time.Now()
	readedIn := end.Sub(start)
	fmt.Printf("Docx readed in in %d\n", readedIn.Milliseconds())
}

func main() {
	//generateRandomXMLFile("./test-generated.xml", 100000, 3000)

	//readXMLFile("./document.xml")
	//generateString(root)

	readDocxFile("./test.docx")

	//b, _ := os.ReadFile("./document.xml")
	//r, err := xmlquery.Parse(bytes.NewReader(b))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(r.FirstChild.Data)
}
