package main

import (
	"fmt"
	"io"
	"net/http"
)

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

type logWriter struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	lw := logWriter{};

	response, _ := http.Get("http://google.com")

	printGreeting(eb)
	printGreeting(sb)

	io.Copy(lw, response.Body)
}
// wil compile, but womt work
func (logWriter) Write([]byte) (int, error) {
	return 1, nil
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

// printGreeting(spanishBot sb){
// 	fmt.Println(sb.getGreeting())
// }


func (englishBot) getGreeting() string {
	return "ABC"
}

func (spanishBot) getGreeting() string {
	return "XYZ"
}