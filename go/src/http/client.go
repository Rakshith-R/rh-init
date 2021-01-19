package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
