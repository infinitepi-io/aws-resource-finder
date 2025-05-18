package main

import (
    "embed"
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "aws-resource-finder/aws"
)

// Version is set during build
var Version = "dev"

//go:embed template/* config/*
var content embed.FS

type ResourceType struct {
    Type    string `json:"type"`
    Display string `json:"display"`
}

type ResourceConfig struct {
    Resources []ResourceType `json:"resources"`
}

func loadResourceTypes() (ResourceConfig, error) {
    var config ResourceConfig
    file, err := content.ReadFile("config/aws-resources.json")
    if err != nil {
        return config, err
    }
    
    err = json.Unmarshal(file, &config)
    return config, err
}

type PageData struct {
    Resources    []ResourceType
    Results      [][]string
    SelectedType string
}

func main() {
    log.Printf("AWS Resource Finder version %s starting...", Version)
    resourceConfig, err := loadResourceTypes()
    if err != nil {
        log.Fatalf("Failed to load resource types: %v", err)
    }

    fs := http.FileServer(http.FS(content))
    http.Handle("/template/", fs)
    
    http.HandleFunc("/aws-resource-finder", func(w http.ResponseWriter, r *http.Request) {
        resourceType := r.URL.Query().Get("resourceType")
        
        if resourceType == "" && len(resourceConfig.Resources) > 0 {
            // Default to first resource type if none specified
            resourceType = resourceConfig.Resources[0].Type
        }
        
        finderQuery := fmt.Sprintf("SELECT * WHERE resourceType = '%s'", resourceType)
        resultFromResourceConfig := aws.GettingResourceFromAwsConfigMapInventory(finderQuery)
        
        data := PageData{
            Resources:    resourceConfig.Resources,
            Results:      resultFromResourceConfig,
            SelectedType: resourceType,
        }
        
        renderTemplate(w, "result", data)
    })

    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFS(content, fmt.Sprintf("template/%s.html", tmpl))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}