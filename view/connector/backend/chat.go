package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func Frontend_ChatMain(w io.Writer, p FrontEndParams) error {
	return template.ChatMainHTML.ExecuteTemplate(w, "base.html", p)
}
