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
    wordFlag := flag.Bool("w", false, "Number of word in file")
    charFlag := flag.Bool("m", false, "Number of characters in file")

    var fileName string
    fileName = os.Args[len(os.Args)-1]
    flag.Parse()

    file, err := os.Open(fileName)
    if err != nil {
	fmt.Println(err)
	os.Exit(1)
    }
    defer file.Close()
    
    switch {
    case *countFlag == true:
	size, err := getFileSize(fileName)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)	
	}
	fmt.Printf("%d %s\n", size, fileName)
	os.Exit(0)

    case *linesFlag == true:
	lineCount, err := countLines(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", lineCount, fileName)
	os.Exit(0)

    case *wordFlag == true:
	wordCount, err := countWords(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", wordCount, fileName)
	os.Exit(0)
    
    case *charFlag == true:
	charCount, err := countChars(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", charCount, fileName)
	os.Exit(0)

    default:
	size, err := getFileSize(fileName)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	lineCount, err := countLines(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	file.Seek(0, 0)
	wordCount, err := countWords(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}

	fmt.Printf("  %d  %d %d %s\n", lineCount, wordCount, size, fileName)
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

func countLines(file *os.File) (int, error) {
    var lineCount int 
    reader := bufio.NewReader(file)
    for {
	_, _, err := reader.ReadLine()
	if err != nil {
	    if err.Error() == "EOF" {
		break
	    } else {
		return 0, err
	    }
	}
	lineCount++
    }

    return lineCount, nil
}

func countWords(file *os.File) (int, error) {
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords) 

    var wordCount int
    for scanner.Scan() {
	wordCount++
    }

    if err := scanner.Err(); err != nil {
	return 0, err
    }

    return wordCount, nil
}

func countChars(file *os.File) (int, error) {
    var charCount int = 0
    reader := bufio.NewReader(file)
    for {
	_, _, err := reader.ReadRune()
	if err != nil {
	    if err.Error() == "EOF" {
		break
	    } else {
		return 0, err
	    }
	}
	charCount++
    }
    return charCount, nil
}

