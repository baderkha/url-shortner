package repository

import (
	"fmt"
	"url-shortner/util"
)

type Link struct {
	ID      string
	URL     string
	MD5Hash string
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
	link.MD5Hash = util.GenerateMD5Hash(link.URL)
	id := util.RandUint64()
	link.ID = fmt.Sprintf("%d", id)
	return l.Create(link)
}

func (l *LinkRepo) DeleteLinkById(id string) bool {
	return l.DeleteById(id)
}

func (l *LinkRepo) FindByHashedUrl(url string) (*Link, bool) {
	var links []Link
	db := l.BaseRepo.GetContext()
	err := db.Scan().Filter("'MD5Hash' = ?", url).All(&links)
	if err != nil {
		fmt.Print(err.Error())
	}
	if len(links) > 0 {
		return &links[0], true
	}
	return nil, false
}
