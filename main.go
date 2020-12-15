package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	data string
	err  error
}

func main() {
	c1 := make(chan result)
	c2 := make(chan result)
	go getTodos(c1)
	go getPosts(c2)
	for i := 0; i < 2; i++ {
		select {
		case todos := <-c1:
			fmt.Println(todos)
		case posts := <-c2:
			fmt.Println(posts)
		}
	}
}

func getTodos(data chan<- result) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		data <- result{err: err}
		return
	}
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		data <- result{err: err}
		return
	}
	data <- result{string(str), nil}
}

func getPosts(data chan<- result) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		data <- result{err: err}
		return
	}
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		data <- result{err: err}
		return
	}
	data <- result{string(str), nil}
