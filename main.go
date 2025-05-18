package main

import(
	"net/http"
	"html/template"
	"log"
	"fmt"
	"aws-resource-finder/aws"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data [][]string) {
	t, err := template.ParseFiles("./template/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		fmt.Println("Template error:", err)
		return
	}
	t.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("./template/images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/aws-resource-finder", func(w http.ResponseWriter, r *http.Request) {
		resourceType := r.URL.Query().Get("resourceType")
		var finderQuery string
		if resourceType == ""{
			resourceType = "AWS::Lambda::Function"
		}
		finderQuery = fmt.Sprintf("SELECT * WHERE resourceType = '%s'", resourceType)
        resultFromResourceConfig := aws.GettingResourceFromAwsConfigMapInventory(finderQuery)
		renderTemplate(w, "result", resultFromResourceConfig)
	})
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}