package viewbackend

import (
	"io"

	"github.com/SQU1DMAN6/ftrchat/view/template"
)

func Frontend_BlogNewBlog(w io.Writer, p FrontEndParams) error {
	return template.BlogNewBlog.ExecuteTemplate(w, "baseblog.html", p)
}

func Frontend_BlogMain(w io.Writer, p FrontEndParams) error {
	return template.BlogListBlogs.ExecuteTemplate(w, "baseblog.html", p)
}

func Frontend_BlogView(w io.Writer, p FrontEndParams) error {
	return template.BlogViewBlog.ExecuteTemplate(w, "baseblog.html", p)
}
