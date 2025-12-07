package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func RegisterMain(w io.Writer, p FrontEndParams) error {
	return template.RegisterMain.ExecuteTemplate(w, "baselogin.html", p)
}
