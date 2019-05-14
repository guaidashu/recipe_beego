package models

import (
	"github.com/astaxie/beego/orm"
)

type RecipeList struct {
	Id           int
	Name         string
	Url          string
	ImgUrl       string
	ImgUrlLarge  string
	Introduce    string
	PageViews    int
	RecipeTypeId int
	Status       int
}

func (r *RecipeList) TableName() string {
	return "recipe_tmp_list"
}

func GetRecipeListWithPage() {

}

func GetRecipeListByCategory(category int, params map[string]interface{}) ([]*RecipeList, int) {
	o := orm.NewOrm()
	var recipeList []*RecipeList
	var limit int
	var offset int
	if _, ok := params["limit"]; ok {
		limit = params["limit"].(int)
	} else {
		limit = 8
	}
	if _, ok := params["offset"]; ok {
		offset = params["offset"].(int)
	} else {
		offset = 0
	}
	query := o.QueryTable("recipe_tmp_list").Filter("recipe_type_id", category)
	num, _ := query.Count()
	_, _ = query.Limit(limit, offset).All(&recipeList)
	return recipeList, int(num)
}

func GetRecipeList(params map[string]interface{}) ([]*RecipeList, int) {
	o := orm.NewOrm()
	var recipeList []*RecipeList
	var limit int
	var offset int
	if _, ok := params["limit"]; ok {
		limit = params["limit"].(int)
	} else {
		limit = 8
	}
	if _, ok := params["offset"]; ok {
		offset = params["offset"].(int)
	} else {
		offset = 0
	}
	num, _ := o.QueryTable("recipe_tmp_list").Limit(limit, offset).All(&recipeList)
	return recipeList, int(num)
}

func GetRecipeListById(id int) RecipeList {
	o := orm.NewOrm()
	var recipeList RecipeList
	_ = o.QueryTable("recipe_tmp_list").Filter("id", id).One(&recipeList)
	return recipeList
}


func GetRecipeListByTitle(name string, params map[string]interface{}) ([]*RecipeList, int) {
	o := orm.NewOrm()
	var recipeList []*RecipeList
	var limit int
	var offset int
	if _, ok := params["limit"]; ok {
		limit = params["limit"].(int)
	} else {
		limit = 8
	}
	if _, ok := params["offset"]; ok {
		offset = params["offset"].(int)
	} else {
		offset = 0
	}
	query := o.QueryTable("recipe_tmp_list").Filter("name__icontains", name)
	num, _ := query.Count()
	_, _ = query.Limit(limit, offset).All(&recipeList)
	return recipeList, int(num)
}

func GetRecipeListByPageViews(params map[string]interface{}) ([]*RecipeList, int) {
	o := orm.NewOrm()
	var recipeList []*RecipeList
	var limit int
	var offset int
	if _, ok := params["limit"]; ok {
		limit = params["limit"].(int)
	} else {
		limit = 8
	}
	if _, ok := params["offset"]; ok {
		offset = params["offset"].(int)
	} else {
		offset = 0
	}
	query := o.QueryTable("recipe_tmp_list").OrderBy("-page_views")
	num, _ := query.Count()
	_, _ = query.Limit(limit, offset).All(&recipeList)
	return recipeList, int(num)
}