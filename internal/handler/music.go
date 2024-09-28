package handler

import (
	"github.com/gin-gonic/gin"
	"library-music/internal/domain"
	"net/http"
)

func (h *Handler) AddMusic(c *gin.Context) {
	var input domain.Music
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Add(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (h *Handler) UpdateMusic(c *gin.Context) {
	var input domain.Music
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Update(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (h *Handler) DeleteMusic(c *gin.Context) {
	song := c.Param("song")
	err := h.service.Delete(song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (h *Handler) GetMusicList(c *gin.Context) {
	musics, err := h.service.GetAll()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"musics": musics,
	})
}

func (h *Handler) GetMusic(c *gin.Context) {
	song := c.Param("song")
	music, err := h.service.Get(song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, music)
}

func (h *Handler) GetTextMusic(c *gin.Context) {
	song := c.Param("song")
	text, err := h.service.GetText(song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}
