package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventCategoryTypeController struct {
	service service.EventCategoryTypeService
}

func NewEventCategoryTypeController(s service.EventCategoryTypeService) EventCategoryTypeController {
	return EventCategoryTypeController{
		service: s,
	}
}

// @Summary Get Event Category Type
// @Description REST API Event Category Type
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryType
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-category-type
func (controller EventCategoryTypeController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	EventCategoryType, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Event Category Type not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Type Found", http.StatusOK, EventCategoryType)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Event Category Type
// @Description REST API Event Category Type
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryType
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /event-category-type
func (controller EventCategoryTypeController) Store(ctx *gin.Context) {
	var req entity.EventCategoryType

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate Event Category Type", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategoryType, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Event Category Type", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Event Category Type", http.StatusOK, EventCategoryType)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Event Category Type
// @Description REST API Event Category Type
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryType
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-category-type/:id
func (controller EventCategoryTypeController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	EventCategoryType, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Type not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Type Found", http.StatusOK, EventCategoryType)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Event Category Type
// @Description REST API Event Category Type
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryType
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /event-category-type/:id
func (controller EventCategoryTypeController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Type not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.EventCategoryType

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Event Category Type", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategoryType, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Event Category Type", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Event Category Type", http.StatusOK, EventCategoryType)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Event Category Type
// @Description REST API Event Category Type
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategoryType
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /event-category-type/:id
func (controller EventCategoryTypeController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category Type not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	EventCategoryType, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Event Category Type", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Event Category Type", http.StatusOK, EventCategoryType)

	ctx.JSON(http.StatusOK, resp)
}
