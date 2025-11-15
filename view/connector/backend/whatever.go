package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func FrontendWhatever(w io.Writer, p FrontEndParams) error {
	return template.FrontendWhatever.ExecuteTemplate(w, "base.html", p)
}
