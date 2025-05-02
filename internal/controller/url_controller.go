package controller

import (
	"github.com/AshiishKarhade/url-shortner-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type URLController struct {
	urlService service.URLService
}

func NewURLController(urlService *service.URLService) *URLController {
	return &URLController{
		urlService: *urlService,
	}
}

func (c *URLController) ShortenURL(ctx *gin.Context) {
	var request struct {
		LongURL string `json:"long_url" binding:"required,url"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url, err := c.urlService.ShortenURL(ctx.Request.Context(), request.LongURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"short_url": url.ShortURL,
		"long_url":  url.LongURL,
	})
}

func (c *URLController) RedirectURL(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")

	longURL, err := c.urlService.GetLongURL(ctx.Request.Context(), shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, longURL)
}

func (c *URLController) UpdateURL(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")

	var req struct {
		LongURL string `json:"long_url" binding:"required,url"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.urlService.UpdateURL(ctx.Request.Context(), shortURL, req.LongURL); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "URL updated successfully"})
}
