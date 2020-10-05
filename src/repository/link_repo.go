package repository

import (
	"gorm.io/gorm"
	"url-shortner/util"
)

type Link struct {
	gorm.Model
	URL     string `gorm:"url"`
	MD5Hash string `gorm:"type:VARCHAR(100);index:md5_idex,unique;"`
}

func (l *Link) BeforeCreate(tx *gorm.DB) (err error) {
	l.MD5Hash = util.GenerateMD5Hash(l.URL)
	return nil
}

type ILinkRepo interface {
	FindLinkById(id string) (*Link, bool)
	FindByHashedUrl(hashedUrl string) (*Link, bool)
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

func (l *LinkRepo) FindByHashedUrl(url string) (*Link, bool) {
	var link Link
	db := l.BaseRepo.GetContext()
	rows := db.Where("md5_hash=?", url).First(&link).RowsAffected
	return &link, rows > 0
}
