package dto

import (
	"fmt"
	"strconv"
	"url-shortner/src/repository"

	"github.com/mattheath/base62"
)

type LinkDto struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func MapLink(link *repository.Link) *LinkDto {
	fmt.Println(link.ID)
	id, _ := strconv.ParseUint(link.ID, 10, 64)
	fmt.Println(id)
	encoded := base62.EncodeInt64(int64(id))
	fmt.Printf("encoded" + encoded)
	return &LinkDto{
		ID:  encoded,
		URL: link.URL,
	}
}

func MapLinkDto(linkDto *LinkDto) *repository.Link {

	decoded := base62.StdEncoding.DecodeToInt64(linkDto.ID)
	fmt.Print("decoded" + fmt.Sprintf("%d", decoded))
	return &repository.Link{
		ID:  fmt.Sprintf("%d", decoded),
		URL: linkDto.URL,
	}
}
