package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventCategoryController struct {
	service service.EventCategoryService
}

func NewEventCategoryController(s service.EventCategoryService) EventCategoryController {
	return EventCategoryController{
		service: s,
	}
}

// @Summary Get Event Category
// @Description REST API Event Category
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-categorys
func (controller EventCategoryController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	EventCategory, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Event Category not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Found", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Event Category
// @Description REST API Event Category
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /event-categorys
func (controller EventCategoryController) Store(ctx *gin.Context) {
	var req entity.EventCategory

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate Event Category", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategory, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Event Category", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Event Category", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Event Category
// @Description REST API Event Category
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 404 {object} nil
// @method [GET]
// @Router /event-categorys/:id
func (controller EventCategoryController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	EventCategory, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Category Found", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Event Category
// @Description REST API Event Category
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /event-categorys/:id
func (controller EventCategoryController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.EventCategory

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Event Category", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategory, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Event Category", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Event Category", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Event Category
// @Description REST API Event Category
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /event-categorys/:id
func (controller EventCategoryController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Event Category not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	EventCategory, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Event Category", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Event Category", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Assign Event Category Role
// @Description REST API Assign Event Category Role
// @Author RasmadIbnu
// @Success 200 {object} entity.EventCategory
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /event-categorys/assign
func (controller EventCategoryController) AssignRole(ctx *gin.Context) {
	var req []entity.EventCategoryRole

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to Assign Event Category Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	EventCategory, err := controller.service.AssignRole(req)

	if err != nil {
		resp := helper.Response("Failed to Assign Event Category Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Assign Event Category Role", http.StatusOK, EventCategory)

	ctx.JSON(http.StatusOK, resp)
}
