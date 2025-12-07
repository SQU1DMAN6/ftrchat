package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func SuccessRegister(w io.Writer, p FrontEndParams) error {
	return template.SuccessRegister.ExecuteTemplate(w, "baselogin.html", p)
}
