package transport

import (
	"context"
	"encoding/json"
	"log"
	"myproject/common/hub"
	"myproject/modules/biz"
	"myproject/modules/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type wsHandler struct {
	biz biz.MessageStorage
	hub *hub.Hub
}

func NewWSHandler(b biz.MessageStorage, hub *hub.Hub) *wsHandler {
	return &wsHandler{biz: b, hub: hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *wsHandler) HandleWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &hub.Client{Conn: conn, Send: make(chan []byte)}
	h.hub.Register <- client

	go h.writePump(client)
	h.readPump(client)
}

func (h *wsHandler) readPump(client *hub.Client) {
	defer func() {
		h.hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var msg model.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		msg.CreatedAt = time.Now()

		if err := h.biz.SaveMessage(context.Background(), &msg); err != nil {
			log.Println("Save error:", err)
			continue
		}

		response, _ := json.Marshal(msg)
		h.hub.Broadcast <- response
	}
}

func (h *wsHandler) writePump(client *hub.Client) {
	defer client.Conn.Close()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
