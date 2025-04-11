package handlers

import (
	"encoding/json"
	"smart/cores"
	"smart/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type TopicsComment struct{}

func (t *TopicsComment) Insert(c *fiber.Ctx) error {

	comment := models.TopicsComment{
		CreatedAt: time.Now(),
	}

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	_, err := cores.PGDB.Model(&comment).Insert()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comments := []models.TopicsComment{}

	err = cores.PGDB.Model(&models.TopicsComment{}).Where("topics_id = ?", comment.TopicsID).Select(&comments)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	WSconns := cores.TopicClients[comment.TopicsID]

	jsonBytes, err := json.Marshal(struct {
		Data   interface{}
		Action string
	}{
		Data:   comments,
		Action: "update-comment",
	})
	if err != nil {
		return err
	}
	for conn, _ := range WSconns {
		conn.WriteMessage(websocket.TextMessage, jsonBytes)
	}

	return c.JSON(fiber.Map{"data": comment})
}

func (t *TopicsComment) Delete(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid comment ID",
		})
	}

	comment := models.TopicsComment{ID: id}

	_, err = cores.PGDB.Model(&comment).WherePK().Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	comments := []models.TopicsComment{}

	err = cores.PGDB.Model(&models.TopicsComment{}).Where("topics_id = ?", comment.TopicsID).Select(&comments)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	WSconns := cores.TopicClients[comment.TopicsID]

	jsonBytes, err := json.Marshal(struct {
		Data   interface{}
		Action string
	}{
		Data:   comments,
		Action: "update-comment",
	})
	if err != nil {
		return err
	}
	for conn, _ := range WSconns {
		conn.WriteMessage(websocket.TextMessage, jsonBytes)
	}

	return c.JSON(fiber.Map{"message": "data successfully inserted"})
}

func (t *TopicsComment) GetByTopicsId(c *fiber.Ctx) error {

	topicId := c.Params("topic_id")

	comments := []models.TopicsComment{}

	err := cores.PGDB.Model(&models.TopicsComment{}).Where("topics_id = ?", topicId).Select(&comments)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"data": comments})
}
