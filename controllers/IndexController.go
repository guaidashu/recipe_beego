package controllers

import "github.com/yy/recipe/models"

type IndexController struct {
	BaseController
}

func (i *IndexController) Get() {
	params := make(map[string]interface{})
	params["limit"] = 8
	imgPrefix := "https://s3-us-west-1.amazonaws.com/static-soource/recipe/images/small/"
	recipeList, _ := models.GetRecipeList(params)
	recipeListDessert, _ := models.GetRecipeListByCategory(4, params)
	recipeListEntrees, _ := models.GetRecipeListByCategory(12, params)
	recipeListSides, _ := models.GetRecipeListByCategory(13, params)
	i.Data["RecipeList"] = recipeList
	i.Data["ImgPrefix"] = imgPrefix
	i.Data["RecipeListDessert"] = recipeListDessert
	i.Data["RecipeListEntrees"] = recipeListEntrees
	i.Data["RecipeListSides"] = recipeListSides
	i.Data["Title"] = "Home"
	i.Layout = "common/layout.html"
	i.TplName = "index/index.html"
}
