package main

import (
	"fmt"
	rss "imartins/parserss/parse"
	"time"
)

func main() {
	timestart := time.Now()

	fmt.Println("Starting get")
	url := "http://golangweekly.com/rss/1g2bo910"
	parserss := rss.New(url)
	parserss.RunParse()
	fmt.Printf("The Title is: %s\n", parserss.GetTitle())
	fmt.Printf("The Description is: %s\n", parserss.GetDescription())
	fmt.Printf("The Link is: %s\n", parserss.GetLink())
	for i := 0; i < len(parserss.GetCategories()); i++ {
		fmt.Printf("The Categories is %s\n", parserss.GetCategories()[i])
	}
	fmt.Printf("The total of Items is: %d\n\n", parserss.GetTotalItems())

	timeend := time.Now()
	fmt.Printf("Benchmark: %s\n", timeend.Sub(timestart))

}
