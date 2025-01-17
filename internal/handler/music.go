package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"library-music/internal/domain/models"
	"library-music/internal/handler/responses"
	"library-music/internal/services"
	"library-music/internal/services/music"
	"net/http"
	"strconv"
	"time"
)

const (
	ErrInvalidArguments = "invalid arguments"
	ErrAlreadyExists    = "already exists"
	ErrRecordNotFound   = "record not found"
	ErrInternalServer   = "internal server error"
	ErrBadRequest       = "Bad request"
)

// @Summary AddMusic
// @Tags music
// @Description A method for creating a new song
// @ID create-music
// @Accept json
// @Produce json
// @Param input body services.MusicToAdd true "Music info to add"
// @Success 200 {object} responses.SuccessID
// @Failure 400 {object} responses.ErrorResponse
// @Failure 409 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/add [post]
func (h *Handler) AddMusic(ctx *gin.Context) {
	var input services.MusicToAdd
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	songDetails, err := h.service.ExternalApi.Info(input.Song, input.Group)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	releaseDate, err := time.Parse("02.01.2006", songDetails.ReleaseDate)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	msc := models.Music{
		Song: input.Song,
		Group: models.Group{
			Name: input.Group,
		},
		Text:        songDetails.Text,
		Link:        songDetails.Link,
		ReleaseDate: releaseDate,
	}

	if err = validateParams(msc); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	id, err := h.service.Music.Add(msc)
	if err != nil {
		if errors.Is(err, music.ErrMusicAlreadyExists) {
			responses.NewErrorResponse(ctx, http.StatusConflict, ErrAlreadyExists)
			return
		}
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		responses.SuccessID{
			ID: id,
		},
	)
}

// @Summary UpdateMusic
// @Tags music
// @Description A method for fully updating song parameters
// @ID update-music
// @Accept json
// @Produce json
// @Param id query int true "Id song"
// @Param input body services.MusicToUpdate true "Music to update"
// @Success 200 {object} responses.SuccessStatus
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 409 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/update [put]
func (h *Handler) UpdateMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 0 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	var input services.MusicToUpdate
	if err = c.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = validateParams(input)
	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.defaultUpdate(c, input, id)
}

// @Summary UpdatePartialMusic
// @Tags music
// @Description A method for updating some song parameters
// @ID update-partial-music
// @Accept json
// @Produce json
// @Param id query int true "Id song"
// @Param input body services.MusicToPartialUpdate true "Music info to update"
// @Success 200 {object} responses.SuccessStatus
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 409 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/update [patch]
func (h *Handler) UpdatePartialMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 0 {
		responses.NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input services.MusicToPartialUpdate
	if err = c.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = validateParams(input)
	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	upd := input.ParsePartial()
	h.defaultUpdate(c, upd, id)
}

func (h *Handler) defaultUpdate(c *gin.Context, upd services.MusicToUpdate, id int) {
	err := h.service.Music.Update(upd, id)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			responses.NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}

		if errors.Is(err, music.ErrMusicAlreadyExists) {
			responses.NewErrorResponse(c, http.StatusConflict, ErrAlreadyExists)
			return
		}

		responses.NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessStatus{
		Status: "success",
	})
}

// @Summary DeleteMusic
// @Tags music
// @Description Method for deleting a song
// @ID delete-music
// @Accept json
// @Produce json
// @Param id query int true "Id song"
// @Success 200 {object} responses.SuccessStatus
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/delete [delete]
func (h *Handler) DeleteMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 0 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.service.Music.Delete(id)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			responses.NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		responses.NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessStatus{
		Status: "success",
	})
}

// @Summary GetAllMusic
// @Tags music
// @Description A method for getting all songs with the ability to filter and paginate
// @ID get-all-music
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param song query string false "Song name"
// @Param group query string false "Music group"
// @Param link query string false "Link song"
// @Param text query string false "Text song"
// @Param releaseDate query string false "Release date" example:"DD.MM.YYYY"
// @Param countSongs query int true "Count songs"
// @Success 200 {object} responses.SuccessMusics
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/getAllMusic/{page} [get]
func (h *Handler) GetAllMusic(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	filters := services.MusicFilterParams{
		Song:        c.Query("song"),
		Group:       c.Query("group"),
		Link:        c.Query("link"),
		ReleaseDate: c.Query("releaseDate"),
	}

	if err = validateParams(filters); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	countSongs, err := strconv.Atoi(c.Query("countSongs"))
	if err != nil || countSongs < 1 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	musics, err := h.service.Music.GetAll(filters, countSongs, page)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			responses.NewErrorResponse(c, http.StatusBadRequest, ErrRecordNotFound)
			return
		}
		responses.NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessMusics{
		Music: musics,
	})
}

// @Summary GetMusic
// @Tags music
// @Description A method for getting information about a specific song
// @ID get-music
// @Accept json
// @Produce json
// @Param group query string true "Music group"
// @Param song query string true "Song name"
// @Success 200 {object} models.Music
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/getMusic [get]
func (h *Handler) GetMusic(c *gin.Context) {
	song := c.Query("song")
	group := c.Query("group")
	res, err := h.service.Music.Get(song, group)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			responses.NewErrorResponse(c, http.StatusBadRequest, ErrBadRequest)
			return
		}
		responses.NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary GetTextMusic
// @Tags music
// @Description A method for getting the lyrics of a song
// @ID get-text-music
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param song query string true "Song name"
// @Param group query string true "Music group"
// @Param countVerse query int true "Count verse"
// @Success 200 {object} responses.SuccessText
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/getTextMusic [get]
func (h *Handler) GetTextMusic(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	song := c.Query("song")
	group := c.Query("group")
	countVerse, err := strconv.Atoi(c.Query("countVerse"))
	if err != nil || countVerse <= 0 {
		responses.NewErrorResponse(c, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	text, err := h.service.Music.GetText(song, group, countVerse, page)
	if err != nil {
		if errors.Is(err, music.ErrMusicNotFound) {
			responses.NewErrorResponse(c, http.StatusNotFound, ErrRecordNotFound)
			return
		}
		responses.NewErrorResponse(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, responses.SuccessText{
		Text: text,
	})
}

func validateParams(value interface{}) error {
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		return fmt.Errorf(ErrInvalidArguments)
	}
	return nil
}
