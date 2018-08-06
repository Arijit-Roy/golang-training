package main

import(
	"fmt"
	"net/http"
)

func main(){

	links:= []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		 go visit(link, c)
	}
	for i:=0; i< len(links); i++{
		fmt.Println(<-c)
	}
}

func visit(link string, c chan string){
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("error ", err)
		c <- "down"
		return
	}

	fmt.Println(link, "is up and running")
	c <- "up"
}
