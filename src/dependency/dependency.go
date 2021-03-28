package dependency

import (
	"url-shortner/src/controller"
	"url-shortner/src/repository"

	"github.com/guregu/dynamo"
)

type Dependency struct {
	*controller.LinkController
}

func MakeDependencies(db *dynamo.DB) *Dependency {
	return &Dependency{
		LinkController: &controller.LinkController{
			LinkRepo: &repository.LinkRepo{
				BaseRepo: repository.BaseRepo{
					Context: db,
					Table:   "shrter",
				},
			},
		},
	}
}
