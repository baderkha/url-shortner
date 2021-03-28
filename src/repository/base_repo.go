package repository

import (
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
	return b.GetContext().Get("ID", id).One(dynamo.AWSEncoding(&model)) == nil
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
