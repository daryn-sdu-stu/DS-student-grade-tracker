package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

var ctx = context.Background()

func initDb() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	initDb()

	var grades1 = make(map[string]int)
	var grades2 = make(map[string]int)

	grades1["123"] = 1
	grades1["456"] = 2

	grades2["555"] = 1
	grades2["666"] = 2

	var course1 = Course{"321", grades1}
	var course2 = Course{"flGTnyq", grades2}

	var s = Student{"Ivan", "123", []Course{course1}}
	var s2 = Student{"Empty", "1234", []Course{course2}}
	s.AddCourse(course1)
	s.ViewCourses(s.ID)

	s2.AddCourse(course2)
	s2.ViewCourses(s2.ID)

}

type Course struct {
	Code          string
	StudentsGrade map[string]int
}

func (c *Course) initCourseDb() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, c.Code, c.StudentsGrade, 0).Err()

	if err != nil {
		panic(err)
	}
}

func (c *Course) getCourseCode() {

}

type Student struct {
	Name    string
	ID      string
	Courses []Course
}

func (s *Student) AddCourse(c Course) {
	s.Courses = append(s.Courses, c)
	val, err := json.Marshal(c)

	if err != nil {
		panic(err)
	}

	clientErr := client.Set(ctx, s.ID, val, 0).Err()
	if clientErr != nil {
		panic(clientErr)
	}

}

func (s *Student) UpdateCourse(c Course) {
	val, err := json.Marshal(c)

	if err != nil {
		panic(err)
	}

	client.Set(ctx, s.ID, val, 0)

	for j := 0; j < len(s.Courses); j++ {
		if s.Courses[j].Code == c.Code {
			s.Courses[j] = c
		}
	}
}

func (s *Student) DeleteCourse(c Course) { // todo use redis
	for j := 0; j < len(s.Courses); j++ {
		if s.Courses[j].Code == c.Code {
			s.Courses = append(s.Courses[:j], s.Courses[j+1:]...)
		}
	}
}

func (s *Student) ViewCourses(id string) {
	courses, err := client.Get(ctx, id).Result()

	val := json.Unmarshal([]byte(courses), &Course{})

	fmt.Println("Course: ", val)

	if err != nil {
		panic(err)
	} else if err == redis.Nil {
		fmt.Println("student does not exist")
	} else {
		fmt.Println("courses: ", courses)
		fmt.Println("student: ", s)
		for key, course := range courses {
			fmt.Println("key: ", key, "course: ", course)

			if string(key) == s.ID {
				fmt.Println("student Id: ", id, " course: ", course)
			}
		}
	}
}
