package responses

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// HTMLResponse represents a response of the Application in HTML format.
type HTMLResponse struct {
	Status   int
	Template string
	Data     interface{}
}

// Write writes a ApplicationResposne to the given response writer encoded as JSON.
func (r *HTMLResponse) Write(w http.ResponseWriter) {
	content := fmt.Sprintf("templates/%s.html", r.Template)
	tmpl, err := template.ParseFiles(content, "templates/layout/header.html", "templates/layout/footer.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(r.Status)
	if err := tmpl.ExecuteTemplate(w, r.Template, r.Data); err != nil {
		log.Println(err)
	}
}
