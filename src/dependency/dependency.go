package dependency

import (
	"gorm.io/gorm"
	"url-shortner/src/controller"
	"url-shortner/src/repository"
)

type Dependency struct {
	*controller.LinkController
}

func MakeDependencies(db *gorm.DB) *Dependency {
	return &Dependency{
		LinkController: &controller.LinkController{
			LinkRepo: &repository.LinkRepo{
				BaseRepo: repository.BaseRepo{
					Context: db,
				},
			},
		},
	}
}
