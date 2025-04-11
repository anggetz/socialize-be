package handlers

import (
	"encoding/json"
	"smart/cores"
	"smart/models"
	"strconv"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Topic struct{}

func (t *Topic) Get(c *fiber.Ctx) error {

	limitParam := c.Query("limit")
	pageParam := c.Query("page")

	if limitParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "limit required",
		})
	}

	if pageParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "page required",
		})
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "page invalid value",
		})
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "limit invalid value",
		})
	}

	topics := []models.Topic{}

	err = cores.PGDB.Model(&models.Topic{}).Limit(limit).Offset((page - 1) * limit).Select(&topics)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"data": topics})
}

func (t *Topic) GetLikes(c *fiber.Ctx) error {

	id := c.Params("id")

	count, err := cores.PGDB.Model(&models.TopicsLike{}).Where("topics_id = ?", id).Count()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"data": count})
}

func (t *Topic) Insert(c *fiber.Ctx) error {

	topic := models.Topic{
		CreatedAt: time.Now(),
	}

	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	_, err := cores.PGDB.Model(&topic).Insert()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "data successfully saved"})
}

func (t *Topic) Delete(c *fiber.Ctx) error {

	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid comment ID",
		})
	}

	topic := models.Topic{ID: id}

	_, err = cores.PGDB.Model(&topic).WherePK().Delete()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	jsonBytes, err := json.Marshal(struct {
		Data   interface{}
		Action string
	}{
		Data:   id,
		Action: "delete-topic",
	})
	if err != nil {
		return err
	}

	WSconns := cores.TopicClients[id]

	for conn, _ := range WSconns {
		conn.WriteMessage(websocket.TextMessage, jsonBytes)
	}

	return c.JSON(fiber.Map{"message": "data successfully deleted"})
}

func (t *Topic) LikeTopic(c *fiber.Ctx) error {

	db := cores.PGDB

	type LikeRequest struct {
		Author   string `json:"author"`
		TopicsID int64  `json:"topics_id"`
	}

	var req LikeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	err := db.RunInTransaction(c.Context(), func(tx *pg.Tx) error {
		// Lock the topic row to prevent race
		topic := new(models.Topic)
		err := tx.Model(topic).
			Where("id = ?", req.TopicsID).
			For("UPDATE").
			Select()
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Topic not found")
		}

		// Check for existing like
		like := new(models.TopicsLike)
		err = tx.Model(like).
			Where("author = ?", req.Author).
			Where("topics_id = ?", req.TopicsID).
			For("UPDATE").
			Select()

		if err == pg.ErrNoRows {
			// Not exists → Insert like
			newLike := &models.TopicsLike{
				Author:    req.Author,
				TopicsID:  req.TopicsID,
				CreatedAt: time.Now(),
			}
			_, err = tx.Model(newLike).Insert()
			if err != nil {
				return err
			}
			c.Locals("likeStatus", "liked")
		} else if err == nil {
			// Exists → Delete like
			_, err = tx.Model(like).WherePK().Delete()
			if err != nil {
				return err
			}
			c.Locals("likeStatus", "unliked")
		} else {
			// Some DB error
			return err
		}

		return nil
	})

	likes, err := cores.PGDB.Model(&models.TopicsLike{}).Where("topics_id = ?", req.TopicsID).Count()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	WSconns := cores.TopicClients[req.TopicsID]

	jsonBytes, err := json.Marshal(struct {
		Data   interface{}
		Action string
	}{
		Data:   likes,
		Action: "update-likes",
	})
	if err != nil {
		return err
	}
	for conn, _ := range WSconns {
		conn.WriteMessage(websocket.TextMessage, jsonBytes)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return toggle result
	return c.JSON(fiber.Map{
		"status": c.Locals("likeStatus"),
	})
}
