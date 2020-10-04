package repository

import (
	"github.com/satori/go.uuid"
)

type Link struct {
	ID  string `gorm:"id"`
	URL string `gorm:"url"`
}

type ILinkRepo interface {
	FindLinkById(id string) (*Link, bool)
	CreateLink(link *Link) bool
	DeleteLinkById(id string) bool
}

type LinkRepo struct {
	BaseRepo
}

func (l *LinkRepo) FindLinkById(id string) (*Link, bool) {
	var link Link
	isFound := l.FindById(id, &link)
	return &link, isFound
}

func (l *LinkRepo) CreateLink(link *Link) bool {
	id := uuid.Must(uuid.NewV4(), nil)
	link.ID = id.String()
	return l.Create(&link)
}

func (l *LinkRepo) DeleteLinkById(id string) bool {
	var link Link
	return l.DeleteById(id, &link)
}
