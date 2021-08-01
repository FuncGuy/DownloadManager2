package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	d := Download{
		Url:           "https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-mp4-file.mp4",
		TargetPath:    "final.mp4",
		TotalSections: 10,
	}

	err := d.Do()

	if err != nil {
		log.Fatalf("An error occured while downloading the file: %s\n", err)
	}

	fmt.Printf("Download completed in %v seconds\n", time.Now().Sub(startTime).Seconds())

}

