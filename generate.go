package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	in "github.com/vinewz/alGOs/internal"
)

func generate() {
	srcP := "src"
	err := os.MkdirAll(srcP, os.ModePerm)
	if err != nil {
		log.Fatalf("Couldn't create SRC dir: %v", err)
	}

	dayP, err := createDayDir(srcP)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for key, val := range in.FnS {
		err := createAlgo(key, val.Args, val.Return, in.DEFAULT_RETURN[val.Return], dayP)
		if err != nil {
			log.Fatalf("Error creating algo %s: %v", key, err)
		}
	}
}

func createDayDir(srcP string) (dayPath string, err error) {
	tdPath := filepath.Join(srcP, "_today")
	dirs, err := os.ReadDir(srcP)
	if err != nil {
		return "", err
	}

	days := len(dirs)

	if days == 0 {
		err := createTodayDir(srcP)
		if err != nil {
			return "", err
		}
		return tdPath, nil
	} else {
		err := os.Rename(tdPath, filepath.Join(srcP, fmt.Sprintf("day%d", days)))
		if err != nil {
			return "", err
		}
		err = createTodayDir(srcP)
		if err != nil {
			return "", err
		}
		return tdPath, nil
	}
}

func createTodayDir(srcP string) (err error) {
	todayPath := filepath.Join(srcP, "_today")
	err = os.MkdirAll(todayPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func createAlgo(fnName, fnArgs, fnReturnT, fnReturnV, dayPath string) error {

	hasNonPrimitive := func() bool {
		return strings.Contains(fnArgs, "in.") || strings.Contains(fnReturnT, "in.")
	}

	importStatement := ""
	if hasNonPrimitive() {
		importStatement = "import in \"github.com/vinewz/alGOs/internal\"\n\n"
	}

	// Generate the content
	content := fmt.Sprintf(
		"package src\n\n"+
			"%s"+
			"func %s(%s) %s {\n\n"+
			"return %s\n"+
			"} \n",
		importStatement, fnName, fnArgs, fnReturnT, fnReturnV)

	// File name and path
	fName := fmt.Sprintf("%s.go", fnName)
	fPath := filepath.Join(dayPath, fName)

	// Write to the file
	byteContent := []byte(content)
	err := os.WriteFile(fPath, byteContent, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
