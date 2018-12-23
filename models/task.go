package models

import(
  "github.com/jinzhu/gorm"
)
// Task model
type Task struct {
  gorm.Model
  Amount        float32
  Status        bool
  Reason        string
  Description   string
  DocPath       string
  ProjectID     uint
  AssigneeID    uint
  AssignerID    uint

}
