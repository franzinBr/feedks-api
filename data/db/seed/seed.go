package seed

import (
	"github.com/franzinBr/feedks-api/constants"
	"github.com/franzinBr/feedks-api/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seed struct {
	Name  string
	Model interface{}
	Run   func(*gorm.DB)
}

func (s *Seed) Execute(db *gorm.DB) {
	count := 0
	db.Model(s.Model).Select("count(*)").Find(&count)

	if count != 0 {
		return
	}

	s.Run(db)
}

func AutoSeed(db *gorm.DB) {
	for _, seed := range All() {
		seed.Execute(db)
	}
}

func All() []Seed {
	return []Seed{
		{
			Name:  "CreateRoles",
			Model: &models.Role{},
			Run: func(db *gorm.DB) {
				db.Create(&models.Role{Name: constants.AdminRole})
				db.Create(&models.Role{Name: constants.DefaultRoleName})
			},
		},
		{
			Name:  "CreateAdminUser",
			Model: &models.User{},
			Run: func(db *gorm.DB) {

				var adminRole models.Role
				db.Where("name = ?", constants.AdminRole).First(&adminRole)

				hashPass, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

				db.Create(&models.User{
					FirstName: "Admin",
					LastName:  "Admin",
					UserName:  "admin",
					Email:     "admin@admin.com.br",
					Password:  string(hashPass),
					RoleID:    int(adminRole.ID),
				})
			},
		},
	}
}
