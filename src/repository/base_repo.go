package repository

import (
	"fmt"

	"github.com/guregu/dynamo"
)

type BaseRepo struct {
	Context *dynamo.DB
	Table   string
}

func (b *BaseRepo) GetContext() dynamo.Table {
	return b.Context.Table(b.Table)
}

func (b *BaseRepo) FindById(id string, model interface{}) bool {
	fmt.Println(id)
	return b.GetContext().Get("ID", id).One(model) == nil
}

func (b *BaseRepo) Create(model interface{}) bool {
	err := b.
		GetContext().
		Put(model).
		Run()
	return err == nil
}

func (b *BaseRepo) DeleteById(id string) bool {
	return b.
		GetContext().
		Delete("ID", id).
		Run() == nil
}
