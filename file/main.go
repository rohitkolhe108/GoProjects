package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	argsWithProg := os.Args
	fileName := os.Args[1]

	fmt.Println(argsWithProg)
	fmt.Println(fileName)

	//Read file
	// bs, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(bs))

	//#2
	// file, err := os.Open(fileName) // For read access.
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

	//#3
	r, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
