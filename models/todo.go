package models

import (
	"GoWebCode/Bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//To_de增删改查都放在这里
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todolist []*Todo, err error) {
	if err := dao.DB.Find(todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATode(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
