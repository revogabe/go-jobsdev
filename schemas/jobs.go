package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Jobs struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Role        string             `bson:"role"`
	Company     string             `bson:"company"`
	Location    string             `bson:"location"`
	Remote      bool               `bson:"remote"`
	Experience  string             `bson:"experience"`
	Salary      string             `bson:"salary"`
	Link        string             `bson:"link"`
	Approved    bool               `bson:"approved"`
}

type JobsResponse struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt  time.Time          `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
	Title      string             `json:"title" bson:"title"`
	Descricao  string             `json:"description" bson:"description"`
	Role       string             `json:"role" bson:"role"`
	Company    string             `json:"company" bson:"company"`
	Location   string             `json:"location" bson:"location"`
	Remote     bool               `json:"remote" bson:"remote"`
	Experience string             `json:"experience" bson:"experience"`
	Salary     string             `json:"salary" bson:"salary"`
	Link       string             `json:"link" bson:"link"`
	Approved   bool               `json:"approved" bson:"approved"`
}
