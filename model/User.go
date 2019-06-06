package model

import "github.com/jinzhu/gorm"

// A User can have many Workouts
type User struct {
	gorm.Model
	Email          string
	HashedPassword string
	Role           string
	Workouts       []Workout
	UserID         uint
}

// A Workout Representation
type Workout struct {
	gorm.Model
	Name        string
	Excercises  string
	Description string
	Result      string
	Date        string
}
