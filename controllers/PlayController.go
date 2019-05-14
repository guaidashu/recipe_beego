package controllers

import (
	"github.com/yy/recipe/models"
	"strings"
)

type PlayController struct {
	BaseController
}

func (b *PlayController) Play() {
	id, _ := b.GetInt("id")
	content := models.GetRecipeContentById(id)
	list := models.GetRecipeListById(id)
	params := make(map[string]interface{})
	params["limit"] = 1
	params["offset"] = 50
	lastNewRecipe, _ := models.GetRecipeList(params)
	b.Data["Content"] = content
	b.Data["LastNewRecipe"] = lastNewRecipe
	ingredients := strings.Split(content.Ingredients, ",")
	b.Data["List"] = list
	b.Data["Ingredients"] = ingredients
	b.Data["Title"] = content.Name
	b.Layout = "common/layout.html"
	b.TplName = "play/play.html"
}

func (b *PlayController) Test()  {
	b.TplName = "play/test.html"
}