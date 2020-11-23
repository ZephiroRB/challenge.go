package seed

import (
	"log"

	"github.com/ZephiroRB/challenge.go/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Skill{}, &models.User{}, &models.UserSkill{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Skill{}, &models.UserSkill{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}
