package errors

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func NewHTML(internal error, public string) *HTML {
	return &HTML{
		Internal: internal,
		Public:   public,
	}
}

type HTML struct {
	Internal error
	Public   string
}

func (h *HTML) Write(w http.ResponseWriter) {
	content := fmt.Sprintf("templates/errors.html")
	tmpl, err1 := template.ParseFiles(content, "templates/layout/header.html", "templates/layout/footer.html")
	if err1 != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Printf("[ERROR]: unexpected unhandled error: %v", h.Internal.Error())

	if err1 := tmpl.ExecuteTemplate(w, "errors", h.Public); err1 != nil {
		log.Println(err1)
	}
	return
}

func (h HTML) Error() string {
	return fmt.Sprintf("Internal: %s | Public: %s", h.Internal, h.Public)
}
