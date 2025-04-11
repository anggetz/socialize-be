package route

import (
	"fmt"
	"smart/cores"
	"smart/handlers"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RunRoute(app *fiber.App) {

	mu := sync.Mutex{}
	cores.TopicClients = make(map[int64]map[*websocket.Conn]bool)

	app.Get("/ws/:topic_id", func(c *fiber.Ctx) error {
		idParam := c.Params("topic_id")
		topicID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return websocket.New(func(conn *websocket.Conn) {
			mu.Lock()
			if cores.TopicClients[topicID] == nil {
				cores.TopicClients[topicID] = make(map[*websocket.Conn]bool)
			}
			cores.TopicClients[topicID][conn] = true
			mu.Unlock()

			fmt.Printf("Client connected to topic %s. Total: %d\n", topicID, len(cores.TopicClients[topicID]))

			defer func() {
				mu.Lock()
				delete(cores.TopicClients[topicID], conn)
				mu.Unlock()
				conn.Close()

				fmt.Printf("Client disconnected from topic %s. Total: %d\n", topicID, len(cores.TopicClients[topicID]))
			}()

			for {
				if _, _, err := conn.ReadMessage(); err != nil {
					break // disconnect
				}
			}
		})(c)
	})

	// Group: /api
	api := app.Group("/api")

	// Group: /api/v1
	v1 := api.Group("/v1")

	// Routes under /api/v1
	topic := v1.Group("/topic")
	topic.Get("/", new(handlers.Topic).Get)
	topic.Post("/", new(handlers.Topic).Insert)
	topic.Delete("/:id", new(handlers.Topic).Delete)
	topic.Post("/like_topic", new(handlers.Topic).LikeTopic)
	topic.Get("/likes/:id", new(handlers.Topic).GetLikes)

	comment := v1.Group("/topic_comment")
	comment.Get("/by-topic-id/:topic_id", new(handlers.TopicsComment).GetByTopicsId)
	comment.Post("/", new(handlers.TopicsComment).Insert)
	comment.Delete("/:id", new(handlers.TopicsComment).Delete)

	// List all routes before starting the server
	for _, route := range app.GetRoutes() {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}
}
