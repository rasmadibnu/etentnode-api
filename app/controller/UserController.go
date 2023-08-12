package controller

import (
	"net/http"
	"strconv"

	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"

	"github.com/gin-gonic/gin"
)

type UserConstroller struct {
	service service.UserService
}

func NewUserConstroller(s service.UserService) UserConstroller {
	return UserConstroller{
		service: s,
	}
}

// @Summary Get User
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 404 {object} nil
// @method [GET]
// @Router /users
func (controller UserConstroller) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	user, err := controller.service.List(m)

	if err != nil {
		resp := helper.Response("User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("User Found", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Insert User
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /users
func (controller UserConstroller) Store(ctx *gin.Context) {
	var req entity.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to crate User", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	_, err := controller.service.FindByUsername(req.Username)

	if err == nil {
		resp := helper.Response("Username alredy taken", http.StatusBadRequest, nil)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	_, err2 := controller.service.FindByEmail(req.Email)

	if err2 == nil {
		resp := helper.Response("E-Mail alredy taken", http.StatusBadRequest, nil)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	user, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.Response("Failed to create User", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("User successfully created", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// // @Summary Get one User
// // @Description REST API User
// // @Author RasmadIbnu
// // @Success 200 {object} entity.User
// // @Failure 404 {object} nil
// // @method [GET]
// // @Router /users/:id
func (controller UserConstroller) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	user, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("User Found", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update User
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /users/:id
func (controller UserConstroller) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req entity.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	user, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.Response("Failed to update User", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("User successfully updated", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete User
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /users/:id
func (controller UserConstroller) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.Response("User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	user, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.Response("Failed to delete User", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.Response("User successfully deleted", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}
