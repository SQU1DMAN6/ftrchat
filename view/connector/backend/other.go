package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func Frontend_Other(w io.Writer, p FrontEndParams) error {
	return template.FrontendOther.ExecuteTemplate(w, "base.html", p)
}
