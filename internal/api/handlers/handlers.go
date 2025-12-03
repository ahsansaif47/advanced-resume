package handlers

import (
	"github.com/ahsansaif47/advanced-resume/internal/api/controllers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service controllers.IWeaviateService
}

func (h *Handler) AddNewResume(ctx *fiber.Ctx) error {
	var resume map[string]any
	if err := ctx.BodyParser(&resume); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := h.service.AddNewResume(resume)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ID": id,
	})
}

func (h *Handler) BatchAddResume(ctx *fiber.Ctx) error {
	var req []map[string]any // expecting array of resumes

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.service.BatchAddResume(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Message": "Added resumes",
	})
}

func (h *Handler) VectorSearch(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	if query == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Query cann't be empty",
		})
	}

	data, err := h.service.VectorSearch(query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Results": data,
	})
}
