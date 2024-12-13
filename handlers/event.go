package handlers

import (
	"context"
	"time"

	"github.com/Men-fish/ticket-v1/models"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    events,
	})
}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	eventId := ctx.Params("eventId")

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	event, err := h.repository.GetOne(context, eventId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    event,
	})
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	event := &models.Event{}

	// Создаём контекст с таймаутом
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Разбираем тело запроса
	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// Переименовываем переменную в createdEvent, чтобы не перезаписывать оригинальный event
	createdEvent, err := h.repository.CreateOne(context, event)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// Возвращаем успешный ответ
	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Event created",
		"data":    createdEvent,
	})
}

func (h *EventHandler) UpdateOne(ctx *fiber.Ctx) error {
	eventId := ctx.Params("eventId")
	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	event, err := h.repository.UpdateOne(context, eventId, updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Event updated",
		"data":    event,
	})
}

func (h *EventHandler) DeleteOne(ctx *fiber.Ctx) error {
	eventId := ctx.Params("eventId")

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context, eventId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId", handler.DeleteOne)
}
