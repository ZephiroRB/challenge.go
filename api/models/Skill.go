package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Skill struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Skill) Prepare() {
	s.ID = 0
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Skill) Validate() error {

	if s.Name == "" {
		return errors.New("Required Name")
	}

	return nil
}

func (s *Skill) SaveSkill(db *gorm.DB) (*Skill, error) {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Skill{}, err
	}
	return s, nil
}
