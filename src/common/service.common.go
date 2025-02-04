package common

import (
	"url-shortner/src/util"
)

type ServiceStruct[T any] struct{}

func (s *ServiceStruct[T]) CreateOneRecord(createDto *T) error {
	return util.DB.Create(&createDto).Error
}

func (s *ServiceStruct[T]) FindAllRecords() ([]T, error) {
	var data []T
	err := util.DB.Find(&data).Error

	return data, err
}

func (s *ServiceStruct[T]) FindOneRecordById(id int) (T, error) {
	var data T
	err := util.DB.Where("id = ?", id).First(&data).Error

	return data, err
}

func (s *ServiceStruct[T]) FindOneRecordByQuery(queryMap map[string]interface{}) (T, error) {
	var data T
	err := util.DB.Where(queryMap).First(&data).Error

	return data, err
}

func (s *ServiceStruct[T]) FindLastRecord() (T, error) {
	var data T
	err := util.DB.Last(&data).Error

	return data, err
}

func (s *ServiceStruct[T]) DeleteOneRecordById(id int) error {
	var data T
	return util.DB.Delete(&data, id).Error
}

func (s *ServiceStruct[T]) UpdateOneRecordById(id int, updateBody T) error {
	var data T
	return util.DB.Model(&data).Where("id = ?", id).Updates(updateBody).Error
}
