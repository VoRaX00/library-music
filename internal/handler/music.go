package handler

import (
	"github.com/gin-gonic/gin"
	"library-music/internal/domain"
	"net/http"
	"strconv"
)

// @Summary AddMusic
// @Tags music
// @Description Create a new music
// @ID create-music
// @Accept json
// @Produce json
// @Param input body domain.MusicToAdd true "Music info to add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/add [post]
func (h *Handler) AddMusic(c *gin.Context) {
	var input domain.MusicToAdd
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
// @Param id query int true "Id song"
// @Param input body domain.MusicToUpdate true "Music info to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/update [put]
func (h *Handler) UpdateMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input domain.MusicToUpdate
	if err = c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Update(input, id)
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
// @Param id query int true "Id song"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/delete [delete]
func (h *Handler) DeleteMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(id)
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
// @Param page query int true "Page number"
// @Success 200 {object} map[string]domain.Music
// @Failure 500 {object} map[string]string
// @Router /api/getAll [get]
func (h *Handler) GetMusicList(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	musics, err := h.service.GetAll(page)
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
// @Param song query string true "Song name"
// @Param group query string true "Music group"
// @Success 200 {object} domain.Music
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/info [get]
func (h *Handler) GetMusic(c *gin.Context) {
	song := c.Query("song")
	group := c.Query("group")

	music, err := h.service.Get(song, group)
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
// @Param song query string true "Song name"
// @Param group query string true "Music group"
// @Param page query int true "Page number"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/getText [get]
func (h *Handler) GetTextMusic(c *gin.Context) {
	song := c.Query("song")
	group := c.Query("group")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.MusicToGet

	if err = c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	text, err := h.service.GetText(song, group, page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}
