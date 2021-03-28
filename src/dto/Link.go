package dto

import (
	"strconv"
	"url-shortner/src/repository"

	"github.com/mattheath/base62"
)

type LinkDto struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func MapLink(link *repository.Link) *LinkDto {
	id, _ := strconv.ParseInt(link.ID, 10, 64)
	encoded := base62.EncodeInt64(id)
	return &LinkDto{
		ID:  encoded,
		URL: link.URL,
	}
}

func MapLinkDto(linkDto *LinkDto) *repository.Link {
	decoded := uint(base62.DecodeToInt64(linkDto.ID))
	return &repository.Link{
		ID:  string(decoded),
		URL: linkDto.URL,
	}
}
