// package main

// import (
// 	"log"
// 	"os"
// 	"text/template"
// )

// var tpl *template.Template

//	func init() {
//		tpl = template.Must(template.ParseFiles("tpl.gohtml"))
//	}
//
//	func main() {
//		err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Parse the template file
	tmpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Correct data structure with a Title field
	data := struct {
		Title string
		Name  string
	}{
		Title: "Welcome Page",
		Name:  "John Doe",
	}

	// Execute the template with the correct data
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
