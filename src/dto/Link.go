package dto

import (
	"github.com/mattheath/base62"
	"gorm.io/gorm"
	"url-shortner/src/repository"
)

type LinkDto struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func MapLink(link *repository.Link) *LinkDto {
	id := link.ID
	encoded := base62.EncodeInt64(int64(id))
	return &LinkDto{
		ID:  encoded,
		URL: link.URL,
	}
}

func MapLinkDto(linkDto *LinkDto) *repository.Link {
	decoded := uint(base62.DecodeToInt64(linkDto.ID))
	return &repository.Link{
		Model: gorm.Model{
			ID: decoded,
		},
		URL: linkDto.URL,
	}
}
