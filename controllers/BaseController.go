package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yy/recipe/models"
)

type HeaderData struct {
	CategoryList []models.RecipeCategory
	LastNewRecipe []*models.RecipeList
}

type BaseController struct {
	beego.Controller
	Header HeaderData
}

func (c *BaseController) Prepare() {

	c.initHeaderData()

	c.Data["Header"] = c.Header
}

func (c *BaseController) initHeaderData() {
	// 初始化Category 列表
	c.Header.CategoryList = models.GetRecipeCategory()
	params := make(map[string]interface{})
	params["limit"] = 2
	c.Header.LastNewRecipe, _ = models.GetRecipeList(params)
}