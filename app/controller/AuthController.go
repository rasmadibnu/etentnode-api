package controller

import (
	"net/http"

	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"etentnode-api/security"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.UserService
	auth    service.AuthService
}

func NewAuthController(s service.UserService, a service.AuthService) AuthController {
	return AuthController{
		service: s,
		auth:    a,
	}
}

// @Summary Login user
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error, nil
// @method [POST]
// @Router /auth/login
func (controller AuthController) Login(ctx *gin.Context) {
	var loginReq entity.User

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp := helper.Response("Login Failed", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	findUser, _ := controller.service.FindByUsername(loginReq.Username)

	if findUser.ID == 0 {
		resp := helper.Response("The Username or Password is Incorrect", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	loggedIn, err := controller.auth.Login(loginReq)

	if err != nil {
		resp := helper.Response("The Username or Password is Incorrect", http.StatusBadRequest, nil)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	token, err := security.CreateToken(loggedIn, 0)

	if err != nil {
		resp := helper.Response("There was an error generating the API token. Please try again", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	active := 1

	_, err = controller.service.Update(entity.User{IsActive: &active}, int(loggedIn.ID))

	if err != nil {
		resp := helper.Response("Failed update user to active", http.StatusInternalServerError, err.Error())

		ctx.JSON(http.StatusInternalServerError, resp)

		return
	}

	resp := helper.Response("Login Successfully", http.StatusOK, map[string]interface{}{"token": token, "user": loggedIn})

	ctx.JSON(http.StatusOK, resp)
}

// // @Description REST API User
// // @Author RasmadIbnu
// // @Success 200 {object} entity.User
// // @Failure 400, 404 {object} err.Error, nil
// // @method [POST]
// // @Router /auth/login
// // @Summary Login user
// //
// func (controller AuthController) RefreshToken(ctx *gin.Context) {
// 	var refTokenReq request.RefreshToken

// 	if err := ctx.ShouldBindJSON(&refTokenReq); err != nil {
// 		resp := helper.Response( "Refresh Token Failed", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	user, err := controller.service.FindById(refTokenReq.UserID)

// 	if err != nil {
// 		resp := helper.Response( "User not found", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	newToken, err := security.CreateToken(user, refTokenReq.Expire)

// 	if err != nil {
// 		resp := helper.Response( "There was an error generating the API token. Please try again", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	resp := helper.Response( "Refresh token successfully", http.StatusOK, newToken)

// 	ctx.JSON(http.StatusOK, resp)
// }

// @Summary logout user
// @Description REST API User
// @Author RasmadIbnu
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error, nil
// @method [POST]
// @Router /auth/logout
func (controller AuthController) Logout(ctx *gin.Context) {
	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	notActive := 0
	_, err := controller.service.Update(entity.User{IsActive: &notActive}, int(user_id.(float64)))

	if err != nil {
		resp := helper.Response("Failed update user to active", http.StatusInternalServerError, err.Error())

		ctx.JSON(http.StatusInternalServerError, resp)

		return
	}

	resp := helper.Response("Logout Successfully", http.StatusOK, nil)

	ctx.JSON(http.StatusOK, resp)
}
