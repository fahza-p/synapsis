package category

import (
	"context"
	"errors"
	"html"
	"time"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/utils"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) Create(ctx context.Context, model *model.Category, userEmail string) error {
	logger := log.GetLogger(ctx, "Category.Service", "Create")
	logger.Info("Create")

	// Check Category Is Exists
	isExists, err := s.category.IsExists(ctx, "slug", model.Slug)
	if err != nil {
		logger.Errorf("can't get category data with slug: '%s'", model.Slug)
		return err
	}

	if isExists {
		return errors.New("data is already exist")
	}

	// Build Create Data
	model.SetCategoryCreateData(userEmail)

	// Create Category
	return s.category.Create(ctx, model)
}

func (s *Service) Remove(ctx context.Context, id string) error {
	logger := log.GetLogger(ctx, "Category.Service", "Remove")
	logger.Info("Remove")

	// Check product that uses this category
	isExists, err := s.product.IsExists(ctx, "category_id", id)
	if err != nil {
		logger.Errorf("can't get product data with id: '%s'", id)
		return err
	}

	if isExists {
		return errors.New("can't delete this category as long as product still uses it")
	}

	// Delete Category
	return s.category.Delete(ctx, "id", id)
}

func (s *Service) UpdatePatch(ctx context.Context, inputFields map[string]interface{}, id string, userEmail string) error {
	logger := log.GetLogger(ctx, "Category.Service", "UpdatePatch")
	logger.Info("UpdatePatch")

	// Build Update Map Data
	updateField, err := s.buildUpdateData(ctx, userEmail, inputFields)
	if err != nil {
		return err
	}

	// Update Category
	return s.category.UpdatePatch(ctx, id, updateField)
}

/* Local Functions */

func (s *Service) buildUpdateData(ctx context.Context, userEmail string, inputFields map[string]interface{}) (map[string]interface{}, error) {
	logger := log.GetLogger(ctx, "Category.Service", "buildUpdateData")
	logger.Info("buildUpdateData")

	var (
		results = make(map[string]interface{}, 0)
		now     = time.Now().Format("2006-01-02 15:04:05")
	)

	for k, v := range inputFields {
		if utils.IsExistFieldByTag(model.CategoryUpdateReq{}, "json", k) {
			switch k {
			case "name":
				results[k] = html.EscapeString(v.(string))
				results["slug"] = utils.ToSlug(inputFields["name"].(string), "_")

				// Check Category Is Exists
				isExists, err := s.category.IsExists(ctx, "slug", results["slug"])
				if err != nil {
					logger.Errorf("can't get category data with slug: '%s'", results["slug"])
					return nil, err
				}

				if isExists {
					return nil, errors.New("data is already exist")
				}
			default:
				results[k] = html.EscapeString(v.(string))
			}
		}
	}

	results["updated_at"] = now
	results["updated_by"] = userEmail

	return results, nil
}
