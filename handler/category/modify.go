package category

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/utils"
	"github.com/fahza-p/synapsis/lib/validation"
	"github.com/fahza-p/synapsis/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Category.Handler", "Create")
	logger.Info("Create")

	var (
		res      response.Build
		reqModel model.CategoryCreateReq
		model    model.Category
		authData = c.Locals("authData").(map[string]interface{})
	)

	if err := c.BodyParser(&reqModel); err != nil {
		logger.Error("Unable to parsing body data")
		res.Msg = err.Error()
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	// Run Validation
	if err := validation.RunValidate(&reqModel, &res); err != nil {
		logger.Errorf("invalid request data %v", reqModel)
		return res.BuildResponse(c, http.StatusUnprocessableEntity)
	}

	model.Name = reqModel.Name
	model.Slug = utils.ToSlug(model.Name, "_")

	if err := h.service.Create(c.Context(), &model, authData["email"].(string)); err != nil {
		logger.WithError(err).Errorf("unable to create category with request data %v", reqModel)
		res.Msg = err.Error()

		switch res.Msg {
		case "data is already exist":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Data = map[string]interface{}{"id": model.Id}
	return res.BuildResponse(c, http.StatusCreated)
}

func (h *Handler) UpdatePatch(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Category.Handler", "UpdatePatch")
	logger.Info("UpdatePatch")

	var (
		res        response.Build
		reqModel   model.CategoryUpdateReq
		input      map[string]interface{}
		categoryId = c.Params("id")
		authData   = c.Locals("authData").(map[string]interface{})
	)

	if err := c.BodyParser(&reqModel); err != nil {
		logger.Error("Unable to parsing body data")
		res.Msg = err.Error()
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	// Run Validation
	if err := validation.RunValidate(&reqModel, &res); err != nil {
		logger.Errorf("invalid request data %v", reqModel)
		return res.BuildResponse(c, http.StatusUnprocessableEntity)
	}

	if err := c.BodyParser(&input); err != nil {
		logger.Error("Unable to parsing body data")
		res.Msg = err.Error()
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	if err := h.service.UpdatePatch(c.Context(), input, categoryId, authData["email"].(string)); err != nil {
		logger.WithError(err).Errorf("unable to update category with request data %v", input)
		res.Msg = err.Error()
		switch res.Msg {
		case "document not found":
			return res.BuildResponse(c, http.StatusNotFound)
		case "data is already exist":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Msg = "successfully updated data"
	return res.BuildResponse(c, http.StatusOK)
}

func (h *Handler) Remove(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Category.Handler", "Remove")
	logger.Info("Remove")

	var (
		res response.Build
		id  = c.Params("id")
	)

	if err := h.service.Remove(c.Context(), id); err != nil {
		logger.WithError(err).Errorf("unable to delete category with id %s", id)
		res.Msg = err.Error()

		switch res.Msg {
		case "document not found":
			return res.BuildResponse(c, http.StatusNotFound)
		case "can't delete this category as long as product still uses it":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	return res.BuildResponse(c, http.StatusOK)
}
