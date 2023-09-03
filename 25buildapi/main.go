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

// Model for course
type Course struct {
	CourseId string `json:"course_id"`
	CourseName string `json:"course_name"`
	CoursePrice int `json:"course_price"`
	Author *Author `json:"author"`

}

type Author struct {
	Fullname string `json:"full_name"`
	Website string `json:"website"`
}

//fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool { 
  return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to APIs")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})
    
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourses).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}


// controllers

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Pingggg!!!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Courses")
	w.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(r)

	for _,course := range courses {
		if course.CourseId == params["id"] {
          json.NewEncoder(w).Encode(course)
		  return
		}
	}
	json.NewEncoder(w).Encode(`No course found`)
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create the Courses")
	w.Header().Set("Content-Type","application/json")

	// what if body is empty
	if r.Body != nil {
		json.NewEncoder(w).Encode(`Please send some data`)
		return
	}

	// what is we get empty data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	
	if course.IsEmpty() {
		json.NewEncoder(w).Encode(`Please provide data inside JSON`)
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
    courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

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

	json.NewEncoder(w).Encode(`Course not present in the database`)
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(`Course deleted from the database`)
	       return
		}
	}
	json.NewEncoder(w).Encode(`Course not present in the database`)
	return
}