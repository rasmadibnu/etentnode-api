package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	service service.StatusService
}

func NewStatusController(s service.StatusService) StatusController {
	return StatusController{
		service: s,
	}
}

// @Summary Get Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 404 {object} nil
// @method [GET]
// @Router /status
func (controller StatusController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	status, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Status Found", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /status
func (controller StatusController) Store(ctx *gin.Context) {
	var req entity.Status

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Status, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Status", http.StatusOK, Status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 404 {object} nil
// @method [GET]
// @Router /status/:id
func (controller StatusController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Status, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Status Found", http.StatusOK, Status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /status/:id
func (controller StatusController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Status

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Status, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Status", http.StatusOK, Status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /status/:id
func (controller StatusController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Status, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Status", http.StatusOK, Status)

	ctx.JSON(http.StatusOK, resp)
}
