package controller

import (
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service service.RoleService
}

func NewRoleController(s service.RoleService) RoleController {
	return RoleController{
		service: s,
	}
}

// @Summary Get Role
// @Description REST API Role
// @Author RasmadIbnu
// @Success 200 {object} entity.Role
// @Failure 404 {object} nil
// @method [GET]
// @Router /roles
func (controller RoleController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Role, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Role Found", http.StatusOK, Role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert Role
// @Description REST API Role
// @Author RasmadIbnu
// @Success 200 {object} entity.Role
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /roles
func (controller RoleController) Store(ctx *gin.Context) {
	var req entity.Role

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Role, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Created Role", http.StatusOK, Role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one Role
// @Description REST API Role
// @Author RasmadIbnu
// @Success 200 {object} entity.Role
// @Failure 404 {object} nil
// @method [GET]
// @Router /roles/:id
func (controller RoleController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	Role, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Role Found", http.StatusOK, Role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update Role
// @Description REST API Role
// @Author RasmadIbnu
// @Success 200 {object} entity.Role
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /roles/:id
func (controller RoleController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.Role

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Role, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Role", http.StatusOK, Role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Role
// @Description REST API Role
// @Author RasmadIbnu
// @Success 200 {object} entity.Role
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /roles/:id
func (controller RoleController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	Role, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete Role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("Successfully Delete Role", http.StatusOK, Role)

	ctx.JSON(http.StatusOK, resp)
}
