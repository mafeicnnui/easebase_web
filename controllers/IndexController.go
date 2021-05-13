package controllers

type IndexController struct {
	BaseController
}

//query tree page
func (c *IndexController) Get() {
	c.TplName = "index.html"
}
