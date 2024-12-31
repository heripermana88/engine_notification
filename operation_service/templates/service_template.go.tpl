package service

import (
    "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/entity"
    "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/repository"
)

type {{.ModelName}}Service struct {
    {{.ModelName}}Repo repository.{{.ModelName}}Repository
}

func New{{.ModelName}}Service(repo repository.{{.ModelName}}Repository) *{{.ModelName}}Service {
    return &{{.ModelName}}Service{
        {{.ModelName}}Repo: repo,
    }
}

func (s *{{.ModelName}}Service) Create{{.ModelName}}({{.ModelNameLower}} *entity.{{.ModelName}}) error {
    return s.{{.ModelName}}Repo.Create{{.ModelName}}({{.ModelNameLower}})
}

func (s *{{.ModelName}}Service) Get{{.ModelName}}ByID(id string) (*entity.{{.ModelName}}, error) {
    return s.{{.ModelName}}Repo.Get{{.ModelName}}ByID(id)
}

func (s *{{.ModelName}}Service) GetAll{{.ModelNames}}() ([]*entity.{{.ModelName}}, error) {
    return s.{{.ModelName}}Repo.GetAll{{.ModelNames}}()
}
