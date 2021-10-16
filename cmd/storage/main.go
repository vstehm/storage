package main

import (
	"fmt"
	"github.com/vstehm/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found File", restoredFile)
}
