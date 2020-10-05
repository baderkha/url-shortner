package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"url-shortner/src/dto"
	"url-shortner/src/repository"
)

type LinkRequest struct {
	URL string ` form:"url" json:"url" binding:"required"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LinkController struct {
	LinkRepo repository.ILinkRepo
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func (lc *LinkController) ShortenLink(c *gin.Context) {
	var link repository.Link
	var linkRequest LinkRequest
	err := c.ShouldBindJSON(&linkRequest)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{
			Message: "You should have a url in the post request",
		})
		return
	}
	if !isValidUrl(linkRequest.URL) {
		c.AbortWithStatusJSON(400, ErrorResponse{
			Message: "Link is Not Valid",
		})
		return
	}
	link.URL = linkRequest.URL
	isCreated := lc.LinkRepo.CreateLink(&link)
	if isCreated {
		dtoLink := dto.MapLink(&link)
		c.JSON(201, dtoLink)
		return
	}
	c.AbortWithStatusJSON(500, ErrorResponse{
		Message: "Link could not be created",
	})
}

func (lc *LinkController) grabLink(c *gin.Context) (*repository.Link, bool) {
	id := c.Param("id")
	link := dto.MapLinkDto(&dto.LinkDto{
		ID: id,
	})
	link, isFound := lc.LinkRepo.FindLinkById(fmt.Sprintf("%d", link.ID))
	return link, isFound
}

func (lc *LinkController) FetchLink(c *gin.Context) {
	link, isFound := lc.grabLink(c)
	if !isFound {
		c.AbortWithStatusJSON(404,
			ErrorResponse{
				Message: "Link was not found for this id",
			},
		)
		return
	}
	dtoLink := dto.MapLink(link)
	c.JSON(200, dtoLink)
}

func (lc *LinkController) ForwardLink(c *gin.Context) {
	link, isFound := lc.grabLink(c)
	if !isFound {
		c.AbortWithStatusJSON(404,
			ErrorResponse{
				Message: "Link was not found for this id",
			},
		)
		return
	}
	dtoLink := dto.MapLink(link)
	c.Redirect(http.StatusPermanentRedirect, dtoLink.URL)
}
