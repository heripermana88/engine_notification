package repository

import (
    "context"
    "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/entity"
    "{{.ProjectName}}/internal/domain/{{.ModelNameLower}}/repository"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type {{.ModelName}}RepositoryImpl struct {
    db *mongo.Collection
}

func New{{.ModelName}}Repository(db *mongo.Database) repository.{{.ModelName}}Repository {
    return &{{.ModelName}}RepositoryImpl{
        db: db.Collection("{{.TableName}}"),
    }
}

func (r *{{.ModelName}}RepositoryImpl) Create{{.ModelName}}({{.ModelNameLower}} *entity.{{.ModelName}}) error {
    _, err := r.db.InsertOne(context.Background(), {{.ModelNameLower}})
    return err
}

func (r *{{.ModelName}}RepositoryImpl) Get{{.ModelName}}ByID(id string) (*entity.{{.ModelName}}, error) {
    var {{.ModelNameLower}} entity.{{.ModelName}}
    err := r.db.FindOne(context.Background(), bson.M{"_id": id}).Decode(&{{.ModelNameLower}})
    return &{{.ModelNameLower}}, err
}

func (r *{{.ModelName}}RepositoryImpl) GetAll{{.ModelNames}}() ([]*entity.{{.ModelName}}, error) {
    var {{.ModelNameLower}}s []*entity.{{.ModelName}}
    cursor, err := r.db.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    for cursor.Next(context.Background()) {
        var {{.ModelNameLower}} entity.{{.ModelName}}
        cursor.Decode(&{{.ModelNameLower}})
        {{.ModelNameLower}}s = append({{.ModelNameLower}}s, &{{.ModelNameLower}})
    }
    return {{.ModelNameLower}}s, nil
}
