package handler

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "{{.ProjectName}}/internal/service"
    "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/entity"
)

type {{.ModelName}}Handler struct {
    {{.ModelName}}Service *service.{{.ModelName}}Service
}

func New{{.ModelName}}Handler({{.ModelNameLower}}Service *service.{{.ModelName}}Service) *{{.ModelName}}Handler {
    return &{{.ModelName}}Handler{ {{.ModelName}}Service: {{.ModelNameLower}}Service }
}

func (h *{{.ModelName}}Handler) Create{{.ModelName}}(w http.ResponseWriter, r *http.Request) {
    var {{.ModelNameLower}} entity.{{.ModelName}}
    json.NewDecoder(r.Body).Decode(&{{.ModelNameLower}})
    err := h.{{.ModelName}}Service.Create{{.ModelName}}(&{{.ModelNameLower}})
    if err != nil {
        http.Error(w, "Failed to create {{.ModelName}}", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode({{.ModelNameLower}})
}

func (h *{{.ModelName}}Handler) Get{{.ModelName}}ByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    {{.ModelNameLower}}, err := h.{{.ModelName}}Service.Get{{.ModelName}}ByID(id)
    if err != nil {
        http.Error(w, "{{.ModelName}} not found", http.StatusNotFound)
        return
    }w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode({{.ModelNameLower}})
}

func (h *{{.ModelName}}Handler) GetAll{{.ModelNames}}(w http.ResponseWriter, r *http.Request) {
    {{.ModelNameLower}}s, err := h.{{.ModelName}}Service.GetAll{{.ModelNames}}()
    if err != nil {
        http.Error(w, "Failed to fetch {{.ModelNames}}", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode({{.ModelNameLower}}s)
}
