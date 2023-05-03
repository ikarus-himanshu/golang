package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course -must go in seperate file
type Course struct {
	CourseId    string  `json:"course"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Model for author -must go in seperate file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake Db
var courses []Course

// middleware, helper -must go in separate file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - by Himanshu")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Himanshu", Website: "hk.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERNSTACK", CoursePrice: 199, Author: &Author{Fullname: "Himanshu", Website: "go.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//Listen to part
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers -must go in separate file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API by Himanshu</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter Some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data Insde")
		return
	}

	//generate a unique id ,string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	//first -grab id from request
	params := mux.Vars(r)

	// loop , id , remove , add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO:Send an response when Id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")
	//first -grab id from request
	params := mux.Vars(r)

	// loop , id , remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
