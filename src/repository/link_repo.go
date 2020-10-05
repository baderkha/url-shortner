package repository

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	URL string `gorm:"url"`
}

type ILinkRepo interface {
	FindLinkById(id string) (*Link, bool)
	FindByUrl(url string) (*Link, bool)
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
	return l.Create(&link)
}

func (l *LinkRepo) DeleteLinkById(id string) bool {
	var link Link
	return l.DeleteById(id, &link)
}

func (l *LinkRepo) FindByUrl(url string) (*Link, bool) {
	var link Link
	db := l.BaseRepo.GetContext()
	rows := db.Where("url=?", url).First(link).RowsAffected
	return &link, rows > 0
}
