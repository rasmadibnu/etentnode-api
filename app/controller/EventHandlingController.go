package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EventHandlingController struct {
	service service.EventHandlingService
}

func NewEventHandlingController(s service.EventHandlingService) EventHandlingController {
	return EventHandlingController{
		service: s,
	}
}

// @Summary Get Event Handling
// @Description REST API Event Handling
// @Author RasmadIbnu
// @Success 200 {object} entity.EventHandling
// @Failure 404 {object} nil
// @method [GET]
// @Router /events
func (controller EventHandlingController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	EventHandling, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Event Handling not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Handling Found", http.StatusOK, EventHandling)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Event Handling
// @Description REST API Event Handling
// @Author RasmadIbnu
// @Success 200 {object} entity.EventHandling
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /events
func (controller EventHandlingController) Store(ctx *gin.Context) {
	var req entity.EventHandling
	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}
	file, err := ctx.FormFile("image")

	if err != nil {
		resp := helper.Response("Failed to get image", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	filename := filepath.Base(file.Filename)
	now := time.Now()

	newFileName := now.Format("20060102-150405") + "-" + filename
	image := filepath.Join("uploads", newFileName)

	err = os.MkdirAll("uploads", os.ModePerm)

	if err != nil {
		resp := helper.Response("Failed to crate upload directory", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	if err := ctx.SaveUploadedFile(file, image); err != nil {
		resp := helper.Response("Failed to save file", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	req.EventID, _ = strconv.Atoi(ctx.PostForm("event_id"))
	req.Description = ctx.PostForm("description")
	req.Identification = ctx.PostForm("type_description")
	req.MinorInjuries, err = strconv.Atoi(ctx.PostForm("minor_injuries"))
	req.SeriouslyInjuries, err = strconv.Atoi(ctx.PostForm("seriously_injuries"))
	req.Die, err = strconv.Atoi(ctx.PostForm("die"))
	req.VictimInvolved = ctx.PostForm("victim_involved")
	req.EventCategoryTypeID, err = strconv.Atoi(ctx.PostForm("event_category_type_id"))
	req.Location = ctx.PostForm("location")
	req.Lat = ctx.PostForm("lat")
	req.Lng = ctx.PostForm("lng")
	req.CreatedBy = int(user_id.(float64))
	req.Image = newFileName

	EventHandling, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create EventHandling", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created EventHandling", http.StatusOK, EventHandling)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Event Handling
// @Description REST API Event Handling
// @Author RasmadIbnu
// @Success 200 {object} entity.EventHandling
// @Failure 404 {object} nil
// @method [GET]
// @Router /events/:id
func (controller EventHandlingController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	EventHandling, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("EventHandling not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("EventHandling Found", http.StatusOK, EventHandling)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Event Handling
// @Description REST API Event Handling
// @Author RasmadIbnu
// @Success 200 {object} entity.EventHandling
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /events/:id
func (controller EventHandlingController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("EventHandling not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.EventHandling
	file, err := ctx.FormFile("image")

	if err == nil {
		filename := filepath.Base(file.Filename)
		now := time.Now()

		newFileName := now.Format("20060102-150405") + "-" + filename
		image := filepath.Join("uploads", newFileName)

		err = os.MkdirAll("uploads", os.ModePerm)

		if err != nil {
			resp := helper.Response("Failed to crate upload directory", http.StatusBadRequest, err.Error())

			ctx.JSON(http.StatusBadRequest, resp)

			return
		}

		if err := ctx.SaveUploadedFile(file, image); err != nil {
			resp := helper.Response("Failed to save file", http.StatusBadRequest, err.Error())

			ctx.JSON(http.StatusBadRequest, resp)

			return
		}

		req.Image = newFileName
	}

	req.EventID, _ = strconv.Atoi(ctx.PostForm("event_id"))
	req.Description = ctx.PostForm("description")
	req.Identification = ctx.PostForm("identification")
	req.MinorInjuries, err = strconv.Atoi(ctx.PostForm("minor_injuries"))
	req.SeriouslyInjuries, err = strconv.Atoi(ctx.PostForm("seriously_injuries"))
	req.Die, err = strconv.Atoi(ctx.PostForm("die"))
	req.VictimInvolved = ctx.PostForm("victim_involved")
	req.EventCategoryTypeID, err = strconv.Atoi(ctx.PostForm("event_category_type_id"))
	req.Location = ctx.PostForm("location")
	req.Lat = ctx.PostForm("lat")
	req.Lng = ctx.PostForm("lng")
	EventHandling, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Event Handling", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Event Handling", http.StatusOK, EventHandling)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete EventHandling
// @Description REST API EventHandling
// @Author RasmadIbnu
// @Success 200 {object} entity.EventHandling
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /events/:id
func (controller EventHandlingController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("EventHandling not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	EventHandling, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete EventHandling", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete EventHandling", http.StatusOK, EventHandling)

	ctx.JSON(http.StatusOK, resp)
}
