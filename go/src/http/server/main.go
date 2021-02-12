package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rakshith-r/rh-init/go/src/http/student"
)

var stud student.Student = student.Student{}

func main() {
	stud.Init()
	http.HandleFunc("/", handler)
	log.Println("Listening on port 3000")
	go func() {
		http.ListenAndServe(":3001", nil)
	}()
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	log.Printf("Request type : %s \n", method)
	if method == "POST" {
		stud.Name = req.PostFormValue("Name")
		stud.Class = req.PostFormValue("Class")
		i, err := strconv.Atoi(req.PostFormValue("Age"))
		if err != nil {
			log.Fatalf("In hello Not a interger : %v", err)
		}
		stud.Age = i
		stud.Write()
	}
	resp := stud.GetHTML()
	log.Println(resp)
	fmt.Fprint(w, resp)
}
