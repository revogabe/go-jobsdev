package main

import "gorm.io/gorm"

type Jobs struct {
  gorm.Model

	Title string
	Description string

	Role string 
	Company string
	Location string
	Remote bool
	Experience string
	Salary string
	Link string
}