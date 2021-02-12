package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", "localhost:3000/", nil)

	if err != nil {
		t.Fatalf("request not created %v", err)
	}

	rec := httptest.NewRecorder()

	index(rec, request)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("excepted status ok, got %v ", res.Status)
	}

}

/*package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status, "\n")
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}*/
