package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func LoginMain(w io.Writer, p FrontEndParams) error {
	return template.LoginMain.ExecuteTemplate(w, "baselogin.html", p)
}
