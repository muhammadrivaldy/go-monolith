package usecase

import (
	"backend/config"
	"backend/handler/template"
	templateEntity "backend/handler/template/entity"
)

type templateUseCase struct {
	config         config.Configuration
	templateEntity templateEntity.TemplateEntity
}

func NewTemplateUseCase(
	config config.Configuration,
	templateEntity templateEntity.TemplateEntity) template.ITemplateUseCase {
	return templateUseCase{
		config:         config,
		templateEntity: templateEntity}
}
