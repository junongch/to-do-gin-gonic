package todo

import (
	"github.com/junongch/to-do-gin-gonic/config"
)

func GetAllTodosSer(todos *[]Todo) (err error) {
	if err = config.DB.Find(todos).Error; err != nil {
		return err
	}

	return nil
}

func GetTodoSer(todo *Todo, id string)(err error) {
	if err = config.DB.First(todo, id).Error; err != nil {
		return err
	}

	return nil
}

func CreateTodoSer(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTodoSer(todo *Todo) (err error) {
	if err = config.DB.Model(todo).Updates(*todo).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTodoSer(id string) (err error) {
	if err = config.DB.Delete(&Todo{}, id).Error; err != nil {
		return err
	}

	return nil
}
