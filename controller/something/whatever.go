package something

import (
	"net/http"

	viewBackend "github.com/SQU1DMAN6/ftrchat/view/connector/backend"
)

func Whatever(w http.ResponseWriter, r *http.Request) {
	p := viewBackend.FrontEndParams{
		Title:   "Flip the Rs",
		Message: "This is a site that is unquestionably unsuspicious.",
	}
	viewBackend.FrontendWhatever(w, p)
}
