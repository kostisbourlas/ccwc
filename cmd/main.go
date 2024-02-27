package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
    
    reader := bufio.NewReader(file)
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
	lineCount, err := countLines(*reader)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", lineCount, fileName)
	os.Exit(0)

    case *wordFlag == true:
	wordCount, err := countWords(*reader)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	fmt.Printf("%d %s\n", wordCount, fileName)
	os.Exit(0)
    
    case *charFlag == true:
	charCount, err := countChars(*reader)
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
	lineCount, err := countLines(*reader)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}
	file.Seek(0, 0)
	wordCount, err := countWords(*reader)
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

func countLines(reader bufio.Reader) (int, error) {
    var lineCount int 
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

func countWords(reader bufio.Reader) (int, error) {
    var wordCount int
    for {
        line, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            return 0, err
        }
        words := strings.Fields(line)
        wordCount += len(words)
        if err == io.EOF {
            break
        }
    }
    return wordCount, nil
}

func countChars(reader bufio.Reader) (int, error) {
    var charCount int 
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

