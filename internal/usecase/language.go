package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"mnc-finance/internal/contract/entity"
	"mnc-finance/internal/contract/model"
	"mnc-finance/util"
	"strconv"
)

func (u *UsecaseImpl) AddLanguage(param model.Language) (int, bool, error) {
	util.Items = append(util.Items, entity.Language{
		Language:       param.Language,
		Appeared:       param.Appeared,
		Created:        param.Created,
		Functional:     param.Functional,
		ObjectOriented: param.ObjectOriented,
		Relation: entity.RelationChild{
			InfluencedBy: param.Relation.InfluencedBy,
			Influences:   param.Relation.Influences,
		},
	})

	return fiber.StatusOK, true, nil
}
func (u *UsecaseImpl) UpdateLanguage(param model.Language, id string) (int, bool, error) {
	languageID, err := strconv.Atoi(id)
	if err != nil {
		return fiber.StatusInternalServerError, false, errors.New("invalid id")
	}

	if len(util.Items)-1 < languageID || languageID < 0 {
		return fiber.StatusBadRequest, false, errors.New("language id not found")
	}

	if param.Language != "" {
		util.Items[languageID].Language = param.Language
	}

	if param.Appeared != nil {
		util.Items[languageID].Appeared = param.Appeared
	}

	if param.Created != nil || len(param.Created) != 0 {
		util.Items[languageID].Created = param.Created
	}

	if param.Functional != nil {
		util.Items[languageID].Functional = param.Functional
	}

	if param.ObjectOriented != nil {
		util.Items[languageID].ObjectOriented = param.ObjectOriented
	}

	if param.Relation.InfluencedBy != nil || len(param.Relation.InfluencedBy) != 0 {
		util.Items[languageID].Relation.InfluencedBy = param.Relation.InfluencedBy
	}

	if param.Relation.Influences != nil || len(param.Relation.Influences) != 0 {
		util.Items[languageID].Relation.Influences = param.Relation.Influences
	}

	return fiber.StatusOK, true, nil
}

func (u *UsecaseImpl) GetDetailLanguage(id string) (*model.Language, int, bool, error) {
	languageID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fiber.StatusInternalServerError, false, errors.New("invalid id")
	}

	if len(util.Items)-1 < languageID || languageID < 0 {
		return nil, fiber.StatusNotFound, false, errors.New("language id not found")
	}

	return &model.Language{
		Language:       util.Items[languageID].Language,
		Appeared:       util.Items[languageID].Appeared,
		Created:        util.Items[languageID].Created,
		Functional:     util.Items[languageID].Functional,
		ObjectOriented: util.Items[languageID].ObjectOriented,
		Relation: model.RelationChild{
			InfluencedBy: util.Items[languageID].Relation.InfluencedBy,
			Influences:   util.Items[languageID].Relation.Influences,
		},
	}, fiber.StatusOK, true, nil
}

func (u *UsecaseImpl) DeleteLanguage(id string) (int, bool, error) {
	languageID, err := strconv.Atoi(id)
	if err != nil {
		return fiber.StatusInternalServerError, false, errors.New("invalid id")
	}

	if len(util.Items)-1 < languageID || languageID < 0 {
		return fiber.StatusNotFound, false, errors.New("language id not found")
	}

	util.Items = append(util.Items[:languageID], util.Items[languageID+1:]...)

	return fiber.StatusOK, true, nil
}
func (u *UsecaseImpl) GetListLanguage() ([]model.Language, int, bool, error) {
	resp := []model.Language{}

	if len(util.Items) > 0 {
		for _, v := range util.Items {
			resp = append(resp, model.Language{
				Language:       v.Language,
				Appeared:       v.Appeared,
				Created:        v.Created,
				Functional:     v.Functional,
				ObjectOriented: v.ObjectOriented,
				Relation: model.RelationChild{
					InfluencedBy: v.Relation.InfluencedBy,
					Influences:   v.Relation.Influences,
				},
			})
		}
	}

	return resp, fiber.StatusOK, true, nil
}
