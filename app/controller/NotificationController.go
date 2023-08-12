package controller

import (
	"encoding/json"
	"errors"
	"etentnode-api/app/entity"
	"etentnode-api/app/service"
	"etentnode-api/helper"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

type NotificationRequest struct {
	Title       string   `json:"title"`
	EventID     int      `json:"event_id"`
	Type        int      `json:"type"`
	ChannelID   string   `json:"channel_id"`
	Message     string   `json:"message"`
	ExternalIDs []string `json:"external_ids"`
	Segments    []string `json:"segments"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow connections from any origin
		},
	}
)

type NotificationController struct {
	clientsMu sync.Mutex
	clients   map[*websocket.Conn]bool
	service   service.NotificationService
}

func NewNotificationController(s service.NotificationService) NotificationController {
	return NotificationController{
		clients: make(map[*websocket.Conn]bool),
		service: s,
	}
}

// @Summary Get Notification
// @Description REST API Notification
// @Author RasmadIbnu
// @Success 200 {object} entity.Notification
// @Failure 404 {object} nil
// @method [GET]
// @Router /notification/:role
func (controller NotificationController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()
	role := ctx.Param("role")
	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}
	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	Event, err := controller.service.List(m, int(user_id.(float64)), role)

	if err != nil {
		resp := helper.Response("Event not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.Response("Event Found", http.StatusOK, Event)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary POST Notification
// @Description REST API Notification
// @Author RasmadIbnu
// @Success 200 {object}
// @Failure 404 {object} nil
// @method [POST]
// @Router /send-notification
func (controller *NotificationController) SendNotificationHandler(ctx *gin.Context) {
	var request NotificationRequest
	user_id, exists := ctx.Get("user_id")
	if !exists {
		resp := helper.Response("User ID not found", http.StatusUnauthorized, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}
	if err := ctx.BindJSON(&request); err != nil {
		resp := helper.Response("Invalid body", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	if request.Title == "" || request.Message == "" {
		resp := helper.Response("Title and message are required", http.StatusBadRequest, request)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	res, err := controller.SendNotification(request.Title, request.Message, request.ChannelID, request.ExternalIDs, request.Segments)

	if err != nil {
		resp := helper.Response("Failed to send notification", http.StatusInternalServerError, err.Error())

		ctx.JSON(http.StatusInternalServerError, resp)

		return
	}

	controller.clientsMu.Lock()
	defer controller.clientsMu.Unlock()

	var notification []entity.Notification

	if len(request.ExternalIDs) > 0 {
		for _, externalID := range request.ExternalIDs {
			notification = append(notification, entity.Notification{
				ExternalID: externalID,
				EventID:    request.EventID,
				Type:       request.Type,
				Title:      request.Title,
				Message:    request.Message,
				SentBy:     int(user_id.(float64)),
			})
		}
	}

	if len(request.Segments) > 0 {
		for _, seg := range request.Segments {
			notification = append(notification, entity.Notification{
				Segment: seg,
				EventID: request.EventID,
				Type:    request.Type,
				Title:   request.Title,
				Message: request.Message,
				SentBy:  int(user_id.(float64)),
			})
		}
	}

	_, err = controller.service.Insert(notification)

	if err != nil {
		resp := helper.Response("Failed insert notification", http.StatusInternalServerError, err.Error())

		ctx.JSON(http.StatusInternalServerError, resp)

		return
	}

	message := NotificationRequest{Title: request.Title, Message: request.Message, ExternalIDs: request.ExternalIDs, Segments: request.Segments}
	jsonData, err := json.Marshal(message)
	if err != nil {
		// Handle error
		return
	}

	for conn := range controller.clients {
		if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Println(err)
		}
	}

	resp := helper.Response("Successfully sent notification", http.StatusOK, string(res))

	ctx.JSON(http.StatusOK, resp)

	return
}

func (controller *NotificationController) SendNotification(title, message, channelID string, externalIDs []string, segments []string) ([]byte, error) {
	notification := map[string]interface{}{
		"app_id": os.Getenv("ONESIGNAL_APP_ID"),
		"contents": map[string]string{
			"en": message,
		},
		"headings": map[string]string{
			"en": title,
		},
	}

	if len(externalIDs) > 0 {
		notification["include_external_user_ids"] = externalIDs
	}

	if channelID != "" {
		notification["android_channel_id"] = channelID
	}

	if len(segments) > 0 {
		notification["included_segments"] = segments
	}

	if len(externalIDs) == 0 && len(segments) == 0 {
		return nil, errors.New("No target specified for the notification")
	}

	client := resty.New()
	sucess, err := client.R().
		SetHeader("Authorization", "Basic "+os.Getenv("ONESIGNAL_API_KEY")).
		SetHeader("Content-Type", "application/json").
		SetBody(notification).
		Post("https://onesignal.com/api/v1/notifications")

	return sucess.Body(), err
}

func (controller *NotificationController) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	controller.clientsMu.Lock()
	controller.clients[conn] = true
	controller.clientsMu.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			controller.clientsMu.Lock()
			delete(controller.clients, conn)
			controller.clientsMu.Unlock()
			break
		}

	}
}

// @Summary Update Status
// @Description REST API Status
// @Author RasmadIbnu
// @Success 200 {object} entity.Status
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /status/:id
func (controller NotificationController) UpdateByEvent(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID
	role := ctx.Param("role")

	var req entity.Notification

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.Response("Failed to update Notification", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	Notification, err := controller.service.UpdateByEvent(req, ID, role)

	if err != nil {
		resp := helper.Response("Failed to update Notification", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.Response("Successfully to Update Notification", http.StatusOK, Notification)

	ctx.JSON(http.StatusOK, resp)
}
