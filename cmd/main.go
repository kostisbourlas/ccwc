package cmd

import (
	"flag"
	"fmt"
	"os"
)
func Execute() {
    countFlag := flag.Bool("c", false, "Size of file in Bytes")
    file := os.Args[0]
    flag.Parse()

    if *countFlag == true {
        fmt.Printf("Count flag: %v\n", *countFlag)
        _ = getFileSize(file)
    }
    os.Exit(0)
}

func getFileSize(file string) int {
    fmt.Println("Hi from getFileSize")
    return 1
}
