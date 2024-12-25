package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"path/filepath"
	in "github.com/vinewz/alGOs/internal"
)

var dayIdx int
func generate() {
  dayP := createDayDir()

	for _, alg := range in.Schema {
	  fName := fmt.Sprintf("%s.go", alg[0])
	  fPath := filepath.Join(dayP, fName)
	  content := fmt.Sprintf("package src\n\nfunc %s(%s) %s {\n\n} \n", alg[0], alg[1], alg[2])
	  byteContent := []byte(content)

	  err := os.WriteFile(fPath, byteContent, os.ModePerm)
	  if err != nil {
	    log.Fatalf("Error creating %s.go: %v", alg[0], err)
	  }
	}
}

func createDayDir() string{
	err := os.MkdirAll("src", os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating SRC dir: %v", err)
	}

	dirs, err := os.ReadDir("./src")
	if err != nil {
		fmt.Printf("Cant read SRC dir: %v", err)
	}

  days := 1
	for _, entry := range dirs {
		if !entry.IsDir() {
			continue
		}

    if strings.Contains(entry.Name(), "day") {
      days++
    }
	}

  dayIdx = days

  dayP := fmt.Sprintf("day%d", dayIdx)
  currentDayPath := path.Join("src", dayP)
	err = os.MkdirAll(currentDayPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating SRC dir: %v", err)
	}

  return currentDayPath
}
