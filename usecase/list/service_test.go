package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/entity"
	listMocks "Jumia_todoList/mocks/usecase/list"
	"Jumia_todoList/test/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListServiceCreation(t *testing.T) {

	service, listMockRepository := setupListService()

	serviceBuildFunc := func(_ map[string]interface{}) (res []interface{}) {
		l := model.ListCreateInput{Name: "Test"}
		id, err := service.Create(l)
		res = append(res, id)
		res = append(res, err)
		return
	}

	mockingFunc := func() {
		l := &entity.List{Name: "Test"}
		listMockRepository.On("Create", l).Return(1, nil).Once()
	}

	errorMockingFunc := func() {
		l := &entity.List{Name: "Test"}
		listMockRepository.On("Create", l).Return(0, fmt.Errorf("error int the repository layer")).Once()
	}
	cases := []helper.CaseService{
		{
			Name:             "ValidListCreation",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					id := res[0]
					return assert.Equal(t, id, 1)
				},
				func(res ...interface{}) bool {
					err := res[1]
					return assert.Equal(t, err, nil)
				},
			},
		},
		{
			Name:             "FailureListCreation",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					id := res[0]
					return assert.NotEqual(t, id, 1)
				},
				func(res ...interface{}) bool {
					err := res[1]
					return assert.Equal(t, err, fmt.Errorf("error int the repository layer"))
				},
			},
		},
	}

	helper.CaseServiceEntryIterator(t, cases)

}

func TestListServiceUpdate(t *testing.T) {

	service, listMockRepository := setupListService()

	serviceBuildFunc := func(_ map[string]interface{}) (res []interface{}) {
		l := model.ListUpdateInput{ID: 1, Name: "Test"}
		err := service.Update(l)
		res = append(res, err)
		return
	}

	mockingFunc := func() {
		l := &entity.List{ID: 1, Name: "Test"}
		listMockRepository.On("Update", l, 1).Return(nil).Once()
	}

	errorMockingFunc := func() {
		l := &entity.List{ID: 1, Name: "Test"}
		listMockRepository.On("Update", l, 1).Return(fmt.Errorf("error int the repository layer")).Once()
	}
	cases := []helper.CaseService{
		{
			Name:             "ValidListUpdate",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					err := res[0]
					return assert.Equal(t, err, nil)
				},
			},
		},
		{
			Name:             "FailureListUpdate",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					err := res[0]
					return assert.Equal(t, err, fmt.Errorf("error int the repository layer"))
				},
			},
		},
	}

	helper.CaseServiceEntryIterator(t, cases)

}

func TestListServiceDelete(t *testing.T) {

	service, listMockRepository := setupListService()

	serviceBuildFunc := func(_ map[string]interface{}) (res []interface{}) {
		l := model.ListRemoveInput{ID: 1}
		err := service.Delete(l)
		res = append(res, err)
		return
	}

	mockingFunc := func() {
		listMockRepository.On("Delete", 1).Return(nil).Once()
	}

	errorMockingFunc := func() {
		listMockRepository.On("Delete", 1).Return(fmt.Errorf("error int the repository layer")).Once()
	}
	cases := []helper.CaseService{
		{
			Name:             "ValidListDelete",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					err := res[0]
					return assert.Equal(t, err, nil)
				},
			},
		},
		{
			Name:             "FailureListDelete",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					err := res[0]
					return assert.Equal(t, err, fmt.Errorf("error int the repository layer"))
				},
			},
		},
	}

	helper.CaseServiceEntryIterator(t, cases)

}

func TestListServiceGet(t *testing.T) {

	service, listMockRepository := setupListService()

	serviceBuildFunc := func(_ map[string]interface{}) (res []interface{}) {
		l := model.GetAllListInput{Limit: 10, Offset: 0}
		err := service.GetAllLists(l)
		res = append(res, err)
		return
	}

	mockingFunc := func() {
		list := make([]entity.List, 0)
		list = append(list, entity.List{ID: 1, Name: "Test"})
		listMockRepository.On("GetAllLists", 10, 0).Return(list, nil).Once()
	}

	errorMockingFunc := func() {
		list := make([]entity.List, 0)
		listMockRepository.On("GetAllLists", 10, 0).Return(list, fmt.Errorf("error int the repository layer")).Once()
	}

	cases := []helper.CaseService{
		{
			Name:             "ValidListGetAll",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					list := make([]entity.List, 0)
					list = append(list, entity.List{ID: 1, Name: "Test"})
					return assert.Equal(t, list, list)
				},
			},
		},
		{
			Name:             "FailureListGetAll",
			ServiceBuildFunc: serviceBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(...interface{}) bool{
				func(res ...interface{}) bool {
					list := res[0]
					return assert.Equal(t, list, make([]entity.List, 0))
				},
			},
		},
	}

	helper.CaseServiceEntryIterator(t, cases)

}

func setupListService() (*Service, *listMocks.Repository) {

	repository := new(listMocks.Repository)
	service := &Service{}
	service.Load(repository, logging.Log)

	return service, repository
}
