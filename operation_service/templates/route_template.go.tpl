package routes

import (
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "{{.ProjectName}}/internal/handler"
    "{{.ProjectName}}/internal/repository"
    "{{.ProjectName}}/internal/service"
)

func Register{{.ModelName}}Routes(router *mux.Router, db *mongo.Database) {
    {{.ModelNameLower}}Repo := repository.New{{.ModelName}}Repository(db)
    {{.ModelNameLower}}Service := service.New{{.ModelName}}Service({{.ModelNameLower}}Repo)
    {{.ModelNameLower}}Handler := handler.New{{.ModelName}}Handler({{.ModelNameLower}}Service)

    router.HandleFunc("/{{.ModelNameLower}}s", {{.ModelNameLower}}Handler.Create{{.ModelName}}).Methods("POST")
    router.HandleFunc("/{{.ModelNameLower}}s/{id}", {{.ModelNameLower}}Handler.Get{{.ModelName}}ByID).Methods("GET")
    router.HandleFunc("/{{.ModelNameLower}}s", {{.ModelNameLower}}Handler.GetAll{{.ModelNames}}).Methods("GET")
}
