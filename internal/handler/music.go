package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"library-music/internal/services"
	"library-music/internal/services/music"
	"net/http"
	"strconv"
	"time"
)

const (
	ErrInvalidArguments = "invalid arguments"
	ErrRecordNotFound   = "record not found"
	ErrInternalServer   = "internal server error"
)

// @Summary AddMusic
// @Tags music
// @Description Create a new music
// @ID create-music
// @Accept json
// @Produce json
// @Param input body services.MusicToAdd true "Music info to add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/add [post]
func (h *Handler) AddMusic(c *gin.Context) {
	var input services.MusicToAdd
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	id, err := h.service.Music.Add(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
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
// @Param input body services.MusicToUpdate true "Music info to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/update [put]
func (h *Handler) UpdateMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	var input services.MusicToUpdate
	if err = c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.service.Music.Update(input, id)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// @Summary UpdatePartialMusic
// @Tags music
// @Description update partial music
// @ID update-partial-music
// @Accept json
// @Produce json
// @Param id query int true "Id song"
// @Param input body services.MusicToUpdate true "Music info to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/update [patch]
func (h *Handler) UpdatePartialMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	var input services.MusicToUpdate
	if err = c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.service.Music.Update(input, id)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
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
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/delete [delete]
func (h *Handler) DeleteMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.service.Music.Delete(id)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
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
// @Param song query string false "Song name"
// @Param group query string false "Music group"
// @Param link query string false "Link song"
// @Param text query string false "Text song"
// @Param page query int true "Page number"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/getAll [get]
func (h *Handler) GetMusicList(c *gin.Context) {
	song := c.Query("song")
	group := c.Query("group")
	link := c.Query("link")
	text := c.Query("text")

	inputDate := c.Query("releaseDate")
	var date time.Time
	if inputDate != "" {
		var err error
		date, err = h.checkedDate(inputDate)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
			return
		}
	}

	filters := services.NewMusicFilterParams(song, group, link, text, date)
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	musics, err := h.service.Music.GetAll(filters, page)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			NewErrorResponse(c, http.StatusBadRequest, ErrRecordNotFound)
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"musics": musics,
	})
}

func (h *Handler) checkedDate(date string) (time.Time, error) {
	val, err := time.Parse("02-01-2006", date)
	return val, err
}

// @Summary GetMusic
// @Tags music
// @Description get music
// @ID get-music
// @Accept json
// @Produce json
// @Param song query string true "Song name"
// @Param group query string true "Music group"
// @Success 200 {object} models.Music
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/info [get]
func (h *Handler) GetMusic(c *gin.Context) {
	song := c.Query("song")
	group := c.Query("group")
	msc, err := h.service.Music.Get(song, group)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"music": msc,
	})
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
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	text, err := h.service.Music.GetText(song, group, page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}
