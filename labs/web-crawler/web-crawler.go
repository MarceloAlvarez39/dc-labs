// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
//
// Crawl3 adds support for depth limiting.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

// Creating a struct to get:
//   links: the urls of a particular search
//   level: the level of the search made by the crawl
type linkList struct {
	links []string
	level int
}

func crawl(url string, file os.File) []string {
	// Write the link into the file
	fmt.Println(url)
	_, err := file.WriteString(url + "\n")
	if err != nil {
		log.Fatal(err)
	}

	// Get the urls of the new link
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {

	worklist := make(chan linkList) //Open the channel
	var waitingList int             // Number of pending sends to worklist

	// Start with the command-line arguments.
	depth := flag.Int("depth", 1, "the depth of the worklist")
	response := flag.String("results", "res.txt", "the name of the txt file")
	flag.Parse()

	if len(flag.Args()) < 1 || len(os.Args) < 3 {
		fmt.Println("Error")
		os.Exit(0)
	}

	// Start the worklist
	go func() { worklist <- linkList{flag.Args(), 0} }()
	waitingList++

	//Create the text file.
	file, err := os.Create(*response)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; waitingList > 0; waitingList-- {
		list := <-worklist
		level := list.level
		for _, link := range list.links {
			if !seen[link] && level <= *depth {
				seen[link] = true
				waitingList++
				go func(link string) {
					// Send the linkList, the file and the level
					// The level increases because is going to make a search.
					worklist <- linkList{crawl(link, *file), level + 1}
				}(link)
			}
		}
	}
}

//!-
