package template

var (
	// Frontend
	FrontendIndex    = ParseBackEnd("themes/backend/index/index.html", "themes/backend/index/redirect.html")
	FrontendOther    = ParseBackEnd("themes/backend/index/other.html")
	FrontendWhatever = ParseBackEnd("themes/backend/dontaskmewhatthisis/whatever.html")
	LoginMain        = ParseBackEndLogin("themes/backend/login/login.html", "themes/backend/login/alert.html")
	RegisterMain     = ParseBackEndLogin("themes/backend/register/register.html", "themes/backend/register/alert.html")
	SuccessRegister  = ParseBackEndLogin("themes/backend/success/successRegister.html")
	ChatMainHTML     = ParseBackEndChat("themes/backend/chat/index.html")
	BlogNewBlog      = ParseBackEndBlog("themes/backend/blog/newblog.html")
	BlogListBlogs    = ParseBackEndBlog("themes/backend/blog/blogindex.html")
)
