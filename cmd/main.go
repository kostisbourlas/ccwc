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
	lineCount, err := countLines(fileName)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", lineCount, fileName)
	os.Exit(0)

    case *wordFlag == true:
	wordCount, err := countWords(fileName)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", wordCount, fileName)
	os.Exit(0)
    
    case *charFlag == true:
	charCount, err := countChars(fileName)
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
	lineCount, err := countLines(fileName)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	wordCount, err := countWords(fileName)
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

func countLines(fileName string) (int, error) {
    file, err := os.Open(fileName)
    if err != nil {
	return 0, err
    }
    defer file.Close()

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

func countWords(fileName string) (int, error) {
    file, err := os.Open(fileName)
    if err != nil {
	return 0, err
    }
    defer file.Close()

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

func countChars(fileName string) (int, error) {
    file, err := os.Open(fileName)
    if err != nil {
	return 0, err
    }
    defer file.Close()

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

