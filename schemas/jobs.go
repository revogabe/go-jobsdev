package schemas

import (
	"time"

	"gorm.io/gorm"
)

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
	Approved bool
}

type JobsResponse struct {
	ID uint `json:"id"`
	
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
 
	Title string `json:"title"`
	Description string `json:"description"`
 
	Role string  `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Experience string `json:"experience"`
	Salary string `json:"salary"`
	Link string `json:"link"`
	Approved bool `json:"approved"`
}