package models

import "github.com/astaxie/beego/orm"

type RecipeContent struct {
	Id          int
	Name        string
	ImgUrl      string
	ImgUrlLarge string
	VideoId     string
	Preparation string
	Ingredients string
	ListId      int
	VideoPath   string
}

func (recipeContent *RecipeContent) TableName() string {
	return "recipe_content"
}

func GetRecipeContentById(id int) RecipeContent {
	o := orm.NewOrm()
	var recipeContent RecipeContent
	_ = o.QueryTable("recipe_content").Filter("list_id", id).One(&recipeContent)
	return recipeContent
}
