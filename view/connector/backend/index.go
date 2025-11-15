package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func Frontend_Home(w io.Writer, p FrontEndParams) error {
	return template.FrontendIndex.ExecuteTemplate(w, "base.html", p)
}
