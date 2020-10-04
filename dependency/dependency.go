package dependency

import (
	"gorm.io/gorm"
	"url-shortner/controller"
	"url-shortner/repository"
)

type Dependency struct {
	*controller.LinkController
}

func MakeDependencies(db *gorm.DB) *Dependency {
	return &Dependency{
		LinkController: &controller.LinkController{
			LinkRepo: &repository.LinkRepo{
				BaseRepo: repository.BaseRepo{
					Contex: db,
				},
			},
		},
	}
}
