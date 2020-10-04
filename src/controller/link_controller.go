package controller

import (
	"github.com/gin-gonic/gin"
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
	link.URL = linkRequest.URL
	isCreated := lc.LinkRepo.CreateLink(&link)
	if isCreated {
		c.JSON(200, link)
		return
	}
	c.AbortWithStatusJSON(500, ErrorResponse{
		Message: "Link could not be created",
	})
}

func (lc *LinkController) FetchLink(c *gin.Context) {
	id := c.Param("id")
	link, isFound := lc.LinkRepo.FindLinkById(id)
	if !isFound {
		c.AbortWithStatusJSON(404,
			ErrorResponse{
				Message: "Link was not found for this id",
			},
		)
		return
	}
	c.JSON(200, link)
}
