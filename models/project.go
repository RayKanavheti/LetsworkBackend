package models

import (
	"github.com/jinzhu/gorm"
)

// Project model
type Project struct {
	gorm.Model
	Title      string
	Desciption string
	DocPath1   string
	DocPath2   string
	DocPath3   string
	Duration   string
	IsComplete bool
	Status     string
	Assisted   string
	Budget     struct {
		minimum float32
		maximum float32
	} `sql:"budget,type:jsonb"`
	ProjectType string // fixed or hourly
	Project     ProjectCategory
	Tasks       []*Task
	OwnerID     uint
}
