package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

// Function to check if a string is CamelCase
func isCamelCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// Function to convert CamelCase to snake_case
func toSnakeCase(s string) string {
	// Use regex to find places where uppercase letters occur
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	// Replace such occurrences with an underscore and lowercase letter
	snake := re.ReplaceAllString(s, "${1}_${2}")
	// Convert the entire string to lowercase
	return strings.ToLower(snake)
}

// Pluralize mengubah kata menjadi bentuk jamak
func Pluralize(word string) string {
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	} else if strings.HasSuffix(word, "y") && !isVowel(string(word[len(word)-2])) {
		return word[:len(word)-1] + "ies"
	}
	return word + "s"
}

// isVowel memeriksa apakah karakter adalah huruf vokal
func isVowel(char string) bool {
	return strings.Contains("aeiou", char)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a model name")
	}

	modelName := os.Args[1]
	projectName := "gitlab.com/nusakti/golang-api-boilerplate" // Ganti dengan nama proyekmu

	modelNameLower := strings.ToLower(modelName)
	tableName := Pluralize(toSnakeCase(modelName))
	if isCamelCase(modelName) {
		modelNameLower = toSnakeCase(modelName)
	}

	modelNames := modelName + "s" // For plural

	// Buat folder yang diperlukan
	os.MkdirAll(fmt.Sprintf("internal/domain/%s/entity", modelNameLower), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("internal/domain/%s/repository", modelNameLower), os.ModePerm)
	// os.MkdirAll(fmt.Sprintf("internal/repository", modelNameLower), os.ModePerm)
	// os.MkdirAll(fmt.Sprintf("internal/service", modelNameLower), os.ModePerm)
	// os.MkdirAll(fmt.Sprintf("internal/handler", modelNameLower), os.ModePerm)
	// os.MkdirAll(fmt.Sprintf("internal/infrastructure/routes", modelNameLower), os.ModePerm)

	// Generate semua file
	filesToGenerate := map[string]string{
		"entity":          "templates/entity_template.go.tpl",
		"repository":      "templates/repository_template.go.tpl",
		"repository_impl": "templates/repository_impl_template.go.tpl",
		"service":         "templates/service_template.go.tpl",
		"handler":         "templates/handler_template.go.tpl",
		"route":           "templates/route_template.go.tpl",
	}

	for _, templatePath := range filesToGenerate {
		generateFile(templatePath, modelName, modelNameLower, modelNames, tableName, projectName)
	}
}

func generateFile(templatePath, modelName, modelNameLower, modelNames, tableName, projectName string) {
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Failed to read template file %s: %v", templatePath, err)
	}

	tpl, err := template.New("template").Parse(string(content))
	if err != nil {
		log.Fatalf("Failed to parse template %s: %v", templatePath, err)
	}

	filePath := generateFilePath(templatePath, modelNameLower)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	data := map[string]string{
		"ModelName":      modelName,
		"ModelNameLower": modelNameLower,
		"ModelNames":     modelNames,
		"TableName":      tableName,
		"ProjectName":    projectName,
	}

	err = tpl.Execute(file, data)
	if err != nil {
		log.Fatalf("Failed to execute template %s: %v", templatePath, err)
	}
}

func generateFilePath(templatePath, modelNameLower string) string {
	switch {
	case strings.Contains(templatePath, "entity"):
		return fmt.Sprintf("internal/domain/%s/entity/%s.go", modelNameLower, modelNameLower)
	case strings.Contains(templatePath, "repository_template"):
		return fmt.Sprintf("internal/domain/%s/repository/%s_repository.go", modelNameLower, modelNameLower)
	case strings.Contains(templatePath, "repository_impl_template"):
		return fmt.Sprintf("internal/repository/%s_repository_impl.go", modelNameLower)
	case strings.Contains(templatePath, "service"):
		return fmt.Sprintf("internal/service/%s_service.go", modelNameLower)
	case strings.Contains(templatePath, "handler"):
		return fmt.Sprintf("internal/handler/%s_handler.go", modelNameLower)
	case strings.Contains(templatePath, "route"):
		return fmt.Sprintf("internal/infrastructure/routes/%s_route.go", modelNameLower)
	default:
		log.Fatal("Unknown template path")
	}

	return ""
}
