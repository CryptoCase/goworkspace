package main

import (
	//Must import the school package because that contains handler.go and model.go
	"SchoolTutorial/school"
	"net/http"
)

func main() {

	http.HandleFunc("/", school.Index)
	http.HandleFunc("/index", school.Index)

	// Hands off the functioning to the handler, routes the execution path.
	http.HandleFunc("/instructor/index", school.InstructorIndex)

	//Type into browser "http://localhost:8080". I increment this by one for each new compile.
	http.ListenAndServe(":8086", nil)
}
