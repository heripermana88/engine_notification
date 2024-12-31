package entity

type {{.ModelName}} struct {
    ID    string `json:"id,omitempty"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
