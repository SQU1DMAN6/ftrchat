package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func Frontend_CategoryNewCategory(w io.Writer, p FrontEndParams) error {
	return template.CategoryNewCategory.ExecuteTemplate(w, "basecategories.html", p)
}
