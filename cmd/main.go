package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
    countFlag := flag.Bool("c", false, "Size of file in Bytes")
    file := os.Args[2]
    flag.Parse()

    var size int 
    var err error
    if *countFlag == true {
	size, err = getFileSize(file)
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)	
	}
    }
    fmt.Printf("%d %s\n", size, file)
    os.Exit(0)
}

func getFileSize(file string) (int, error) {
    fileStats, err := os.Stat(file)
    if err != nil {
	return 0, err 
    }
    return int(fileStats.Size()), nil
}
