package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserSkill struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	RUser   User   `json:"ruser"`
	RUserID uint32 `sql:"type:int REFERENCES users(id)" json:"ruser_id"`

	RSkill   Skill  `json:"rskill"`
	RSkillID uint32 `sql:"type:int REFERENCES skills(id)" json:"rskill_id"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (us *UserSkill) Prepare() {
	us.ID = 0

	us.RUser = User{}
	us.RSkill = Skill{}
	us.CreatedAt = time.Now()
	us.UpdatedAt = time.Now()
}

func (us *UserSkill) Validate() error {

	if us.RUserID < 1 {
		return errors.New("Required User")
	}

	if us.RSkillID < 1 {
		return errors.New("Required Skill")
	}

	return nil
}

func (us *UserSkill) SaveUserSkill(db *gorm.DB) (*UserSkill, error) {
	var err error
	err = db.Debug().Model(&UserSkill{}).Create(&us).Error
	if err != nil {
		return &UserSkill{}, err
	}
	if us.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", us.RUserID).Take(&us.RUser).Error
		if err != nil {
			return &UserSkill{}, err
		}

		err = db.Debug().Model(&Skill{}).Where("id = ?", us.RSkillID).Take(&us.RSkill).Error
		if err != nil {
			return &UserSkill{}, err
		}

	}
	return us, nil
}
