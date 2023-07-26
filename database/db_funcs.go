package database

import (
	"reflect"

	"gorm.io/gorm"
)

type TestInstance struct {
	gorm.Model
	Name         string `json:"name"`
	Unique_field string `gorm:"unique;not null; default:null" json:"email"`
}

func Save(m IModel) *gorm.DB {
	return DB.Save(m)
}

func GetOne(m IModel, conds ...interface{}) *gorm.DB {
	is_list := reflect.TypeOf(conds).Kind() == reflect.Slice || reflect.TypeOf(conds).Kind() == reflect.Array

	if is_list {
		return DB.Where(conds[0], conds[1:]...).First(m)
	} else {
		return DB.First(&m, conds[0])
	}

}

func Update(m IModel, values interface{}) *gorm.DB {
	return DB.Model(m).Updates(values)
}

func Delete(m IModel) *gorm.DB {
	return DB.Delete(m)
}

func MakeMigrations(modelList []interface{}) []error {
	var error_list []error
	for _, model := range modelList {
		error_list = append(error_list, DB.AutoMigrate(&model))
	}
	return error_list
}
