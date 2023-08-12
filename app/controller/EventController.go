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

type EventController struct {
	service service.EventService
}

func NewEventController(s service.EventService) EventController {
	return EventController{
		service: s,
	}
}

// @Summary Get Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 404 {object} nil
// @method [GET]
// @Router /events
func (controller EventController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Event, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Found", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 404 {object} nil
// @method [GET]
// @Router /events/assign/:id
func (controller EventController) ListByUserAssign(ctx *gin.Context) {
	param := ctx.Request.URL.Query()
	id := ctx.Param("id")

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Event, err := controller.service.ListByUserAssign(m, id)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Found", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 404 {object} nil
// @method [GET]
// @Router /events/count
func (controller EventController) GetCountEvent(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	Event, err := controller.service.GetCountEvent(m, int(user_id.(float64)))

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Found", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /events
func (controller EventController) Store(ctx *gin.Context) {
	var req entity.Event
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

	req.Description = ctx.PostForm("description")
	req.TypeDescription = ctx.PostForm("type_description")
	req.OtherDescription = ctx.PostForm("other_description")
	req.TypeVehicleInvolved = ctx.PostForm("type_vehicle_involved")
	req.VictimInvolved = ctx.PostForm("victim_involved")
	req.Responsible = ctx.PostForm("responsible")
	req.Lat = ctx.PostForm("lat")
	req.Lng = ctx.PostForm("lng")
	req.Location = ctx.PostForm("location")
	req.EventCategoryID, _ = strconv.Atoi(ctx.PostForm("event_category_id"))
	req.EventCategoryTypeID, _ = strconv.Atoi(ctx.PostForm("event_category_type_id"))
	req.StatusID, _ = strconv.Atoi(ctx.PostForm("status_id"))
	req.CreatedBy = int(user_id.(float64))
	req.Image = newFileName

	Event, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Event", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Event", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 404 {object} nil
// @method [GET]
// @Router /events/:id
func (controller EventController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Event, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Found", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /events/:id
func (controller EventController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}
	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Event
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

	req.Description = ctx.PostForm("description")
	req.TypeDescription = ctx.PostForm("type_description")
	req.OtherDescription = ctx.PostForm("other_description")
	req.TypeVehicleInvolved = ctx.PostForm("type_vehicle_involved")
	req.VictimInvolved = ctx.PostForm("victim_involved")
	req.Responsible = ctx.PostForm("responsible")
	req.Lat = ctx.PostForm("lat")
	req.Lng = ctx.PostForm("lng")
	req.Location = ctx.PostForm("location")
	req.EventCategoryID, _ = strconv.Atoi(ctx.PostForm("event_category_id"))
	req.EventCategoryTypeID, _ = strconv.Atoi(ctx.PostForm("event_category_type_id"))
	req.StatusID, _ = strconv.Atoi(ctx.PostForm("status_id"))

	req.UpdatedBy = int(user_id.(float64))

	Event, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Event", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Event", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Event
// @Description REST API Event
// @Author RasmadIbnu
// @Success 200 {object} entity.Event
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /events/:id
func (controller EventController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Event, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Event", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Event", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Assign User Handling Role
// @Description REST API Assign User Handling Role
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /events/assign
func (controller EventController) AssignUser(ctx *gin.Context) {
	var req []entity.EventUserHandling

	user_id, exists := ctx.Get("user_id")

	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to Assign User Handling ", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategory, err := controller.service.AssignUser(req, int(user_id.(float64)))

	if err != nil {
		resp := helper.Response("Failed to Assign User Handling Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Assign User Handling Role", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}
