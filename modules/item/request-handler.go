package item

import (
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/dto"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	CreateItem(c *gin.Context)
}

func (rq RequestHandler) CreateItem(c *gin.Context) {

	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	sellIn, err := strconv.ParseInt(c.Request.FormValue("sell_in"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	quality, err := strconv.ParseInt(c.Request.FormValue("quality"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	payload := CreateItemPayload{
		Nama:    c.Request.FormValue("name"),
		SellIn:  int(sellIn),
		Quality: int(quality),
	}

	form, _ := c.MultipartForm()
	files := form.File["images[]"]

	var images = make([]multipart.File, len(files))
	for i, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
			return
		}
		images[i] = file
		err = file.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.DefaultErrorResponseWithMessage(err.Error()))
			return
		}
	}
	payload.Images = images

	res, err := rq.ctrl.CreateItem(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, res)
}
