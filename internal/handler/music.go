package handler

import (
	"github.com/gin-gonic/gin"
	"library-music/internal/domain"
	"net/http"
)

// @Summary AddMusic
// @Tags music
// @Description create music
// @ID create-music
// @Accept json
// @Produce json
// @Param input body domain.Music true "music info"
// @Success 200 {object} map[string]string
// @Router /api/add [post]
func (h *Handler) AddMusic(c *gin.Context) {
	var input domain.Music
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Add(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// @Summary UpdateMusic
// @Tags music
// @Description update music
// @ID update-music
// @Accept json
// @Produce json
// @Param input body domain.Music true "music info"
// @Success 200 {object} map[string]string
// @Router /api/update [put]
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

// @Summary DeleteMusic
// @Tags music
// @Description delete music
// @ID delete-music
// @Accept json
// @Produce json
// @Param song path string true "Song Name"
// @Success 200 {object} map[string]string
// @Router /api/delete/{song} [delete]
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

// @Summary GetAllMusic
// @Tags music
// @Description get all music
// @ID get-all-music
// @Accept json
// @Produce json
// @Success 200 {object} map[string]domain.Music
// @Router /api/getAll [get]
func (h *Handler) GetMusicList(c *gin.Context) {
	musics, err := h.service.GetAll()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"musics": musics,
	})
}

// @Summary GetMusic
// @Tags music
// @Description get music
// @ID get-music
// @Accept json
// @Produce json
// @Param song path string true "Song Name"
// @Success 200 {object} domain.Music
// @Router /api/get/{song} [get]
func (h *Handler) GetMusic(c *gin.Context) {
	song := c.Param("song")
	music, err := h.service.Get(song)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, music)
}

// @Summary GetTextMusic
// @Tags music
// @Description get text music
// @ID get-text-music
// @Accept json
// @Produce json
// @Param song path string true "Song Name"
// @Success 200 {object} map[string]string
// @Router /api/getText/{song} [get]
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
