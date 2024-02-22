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

    var file string
    file = os.Args[len(os.Args)-1]
    flag.Parse()
    
    switch {
    case *countFlag == true:
	size, err := getFileSize(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)	
	}
	fmt.Printf("%d %s\n", size, file)
	os.Exit(0)

    case *linesFlag == true:
	lineCount, err := countLines(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", lineCount, file)
	os.Exit(0)

    case *wordFlag == true:
	wordCount, err := countWords(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", wordCount, file)
	os.Exit(0)
    
    case *charFlag == true:
	charCount, err := countChars(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", charCount, file)
	os.Exit(0)
    default:
	size, err := getFileSize(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	lineCount, err := countLines(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	wordCount, err := countWords(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}

	fmt.Printf("  %d  %d %d %s\n", lineCount, wordCount, size, file)
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

    if err := scanner.Err(); err != nil {
	return 0, err
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

