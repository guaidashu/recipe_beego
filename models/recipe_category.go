package models

import (
	"github.com/astaxie/beego/orm"
)

type RecipeCategory struct {
	Id       int
	Name     string
	PageNum  string
	Keyword  string
	NavType  int
	ParentId int
	Status   int
}

type RecipeSearch struct {
	Id string
	Name string
}

func (recipeCategory *RecipeCategory) TableName() string {
	return "recipe_type"
}

func GetRecipeCategory() []RecipeCategory {
	o := orm.NewOrm()
	var recipeCategoryList []RecipeCategory
	_, _ = o.Raw("select * from recipe_type where name!='All Recipes'").QueryRows(&recipeCategoryList)
	return recipeCategoryList
}

func GetRecipeCategoryById(id int) RecipeCategory {
	o := orm.NewOrm()
	var recipeCategory RecipeCategory
	_ = o.QueryTable("recipe_type").Filter("id", id).One(&recipeCategory)
	return recipeCategory
}
