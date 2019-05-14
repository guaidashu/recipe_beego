package controllers

import "github.com/yy/recipe/models"

type SearchController struct {
	BaseController
}

func (s *SearchController) Category() {
	category, _ := s.GetInt("category")
	page, _ := s.GetInt("page")
	if (page == 0) {
		page = 1
	}
	size := 12
	params := make(map[string]interface{})
	params["limit"] = size
	params["offset"] = models.GetCurrentDataPos(page, size)
	recipeList, num := models.GetRecipeListByCategory(category, params)
	recipeCategory := models.GetRecipeCategoryById(category)
	Page := models.PageUtil(num, page, size, recipeList)
	s.Data["RecipeCategory"] = recipeCategory
	s.Data["RecipeList"] = recipeList
	s.Data["Page"] = Page
	s.Data["SearchType"] = "category"
	s.Data["Title"] = recipeCategory.Name
	s.Data["ImgPrefix"] = "https://s3-us-west-1.amazonaws.com/static-soource/recipe/images/small/"
	s.Layout = "common/layout.html"
	s.TplName = "search/category.html"
}

func (s *SearchController) Search() {
	search := s.GetString("search")
	page, _ := s.GetInt("page")
	if (page == 0){
		page = 1
	}
	size := 12
	params := make(map[string]interface{})
	params["limit"] = size
	params["offset"] = models.GetCurrentDataPos(page, size)
	var recipeCategory models.RecipeSearch
	recipeCategory.Name = search
	recipeCategory.Id = search
	recipeList, num := models.GetRecipeListByTitle(search, params)
	Page := models.PageUtil(num, page, size, recipeList)
	s.Data["RecipeCategory"] = recipeCategory
	s.Data["RecipeList"] = recipeList
	s.Data["Page"] = Page
	s.Data["SearchType"] = "search"
	s.Data["Title"] = recipeCategory.Name
	s.Data["ImgPrefix"] = "https://s3-us-west-1.amazonaws.com/static-soource/recipe/images/small/"
	s.Layout = "common/layout.html"
	s.TplName = "search/category.html"
}

func (s *SearchController) Top() {
	page, _ := s.GetInt("page")
	if (page == 0){
		page = 1
	}
	size := 12
	params := make(map[string]interface{})
	params["limit"] = size
	params["offset"] = models.GetCurrentDataPos(page, size)
	var recipeCategory models.RecipeSearch
	recipeCategory.Name = "top"
	recipeCategory.Id = "top"
	recipeList, num := models.GetRecipeListByPageViews(params)
	Page := models.PageUtil(num, page, size, recipeList)
	s.Data["RecipeCategory"] = recipeCategory
	s.Data["RecipeList"] = recipeList
	s.Data["Page"] = Page
	s.Data["SearchType"] = "top"
	s.Data["Title"] = "Top"
	s.Data["ImgPrefix"] = "https://s3-us-west-1.amazonaws.com/static-soource/recipe/images/small/"
	s.Layout = "common/layout.html"
	s.TplName = "search/category.html"
}
