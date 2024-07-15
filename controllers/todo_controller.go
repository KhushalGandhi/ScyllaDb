package controllers

import (
	"encoding/base64"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"scylladb/models"
	"scylladb/services"
)

type TODOController struct {
	Service *services.TODOService
}

func (c *TODOController) CreateTODO(ctx *fiber.Ctx) error {
	todo := new(models.TODO)
	if err := ctx.BodyParser(todo); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := c.Service.Create(todo); err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create TODO item"})
	}

	return ctx.JSON(todo)
}

func (c *TODOController) GetTODO(ctx *fiber.Ctx) error {
	userID, err := gocql.ParseUUID(ctx.Params("user_id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	id, err := gocql.ParseUUID(ctx.Params("id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	todo, err := c.Service.GetByID(userID, id)
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot get TODO item"})
	}

	return ctx.JSON(todo)
}

func (c *TODOController) ListTODOs(ctx *fiber.Ctx) error {
	userID, err := gocql.ParseUUID(ctx.Params("user_id"))
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	status := ctx.Query("status")
	limit := ctx.QueryInt("limit", 10)
	sortBy := ctx.Query("sort_by", "asc")
	var pageState []byte

	if ctx.Query("page_state") != "" {
		pageState, err = base64.StdEncoding.DecodeString(ctx.Query("page_state"))
		if err != nil {
			log.Error(err)

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page state"})
		}
	}

	todos, newPageState, err := c.Service.List(userID, status, limit, pageState, sortBy)
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot list TODO items"})
	}

	return ctx.JSON(fiber.Map{
		"todos":      todos,
		"page_state": base64.StdEncoding.EncodeToString(newPageState),
	})
}

func (c *TODOController) UpdateTODO(ctx *fiber.Ctx) error {
	userID, err := gocql.ParseUUID(ctx.Params("user_id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	id, err := gocql.ParseUUID(ctx.Params("id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	todo := new(models.TODO)
	if err := ctx.BodyParser(todo); err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	todo.ID = id
	todo.UserID = userID

	if err := c.Service.Update(todo); err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update TODO item"})
	}

	return ctx.JSON(todo)
}

func (c *TODOController) DeleteTODO(ctx *fiber.Ctx) error {
	userID, err := gocql.ParseUUID(ctx.Params("user_id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	id, err := gocql.ParseUUID(ctx.Params("id"))
	if err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	if err := c.Service.Delete(userID, id); err != nil {
		log.Error(err)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot delete TODO item"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
