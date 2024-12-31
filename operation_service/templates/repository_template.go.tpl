package repository

import "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/entity"

type {{.ModelName}}Repository interface {
    Create{{.ModelName}}({{.ModelNameLower}} *entity.{{.ModelName}}) error
    Get{{.ModelName}}ByID(id string) (*entity.{{.ModelName}}, error)
    GetAll{{.ModelNames}}() ([]*entity.{{.ModelName}}, error)
}
