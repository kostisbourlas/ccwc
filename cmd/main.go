package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Execute() {
    countFlag := flag.Bool("c", false, "Size of file in Bytes")
    linesFlag := flag.Bool("l", false, "Number of lines in file")

    file := os.Args[2]
    flag.Parse()

    if *countFlag == true {
	size, err := getFileSize(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)	
	}
	fmt.Printf("%d %s\n", size, file)
	os.Exit(0)
    }

    if *linesFlag == true {
	lineCount, err := countLines(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", lineCount, file)
	os.Exit(0)
    }
}

func getFileSize(fileName string) (int, error) {
    fileStats, err := os.Stat(fileName)
    if err != nil {
	return 0, err 
    }
    return int(fileStats.Size()), nil
}

func countLines(fileName string) (int, error) {
    file, err := os.Open(fileName)
    if err != nil {
	return 0, err
    }
    defer file.Close()

    var lineCount int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	lineCount++
    }

    return lineCount, nil
}
