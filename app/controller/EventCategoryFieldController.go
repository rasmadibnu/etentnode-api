package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventCategoryFieldController struct {
	service service.EventCategoryFieldService
}

func NewEventCategoryFieldService(s service.EventCategoryFieldService) EventCategoryFieldController {
	return EventCategoryFieldController{
		service: s,
	}
}

// @Summary Get Event Category Field
// @Description REST API Event Category Field
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryField
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-category-field
func (controller EventCategoryFieldController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	EventCategoryField, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Event Category Field not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Field Found", http.StatusOK, EventCategoryField)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Event Category Field
// @Description REST API Event Category Field
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryField
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /event-category-field
func (controller EventCategoryFieldController) Store(ctx *gin.Context) {
	var req entity.EventCategoryField

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate Event Category Field", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategoryField, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Event Category Field", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Event Category Field", http.StatusOK, EventCategoryField)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Event Category Field
// @Description REST API Event Category Field
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryField
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-category-field/:id
func (controller EventCategoryFieldController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	EventCategoryField, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Field not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Field Found", http.StatusOK, EventCategoryField)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Event Category Field
// @Description REST API Event Category Field
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryField
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /event-category-field/:id
func (controller EventCategoryFieldController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Field not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.EventCategoryField

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Event Category Field", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategoryField, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Event Category Field", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Event Category Field", http.StatusOK, EventCategoryField)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Event Category Field
// @Description REST API Event Category Field
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryField
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /event-category-field/:id
func (controller EventCategoryFieldController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Field not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	EventCategoryField, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Event Category Field", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Event Category Field", http.StatusOK, EventCategoryField)

	ctx.JSON(http.StatusOK, resp)
}
