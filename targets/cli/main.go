package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var path = os.Args[1]
	fmt.Printf("for path: %s\n", path)
	for idx, str := range os.Args[2:] {
		fmt.Printf("\n-- %d : %s\n", idx, str)
	}
	fmt.Println("\nstart")
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("end")

}
