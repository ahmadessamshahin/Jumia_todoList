package list

import (
	"Jumia_todoList/config/database"
	env "Jumia_todoList/config/environment"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	createdList *entity.List
)

type SuiteTester struct {
	suite.Suite

	listRepository *Instance
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(SuiteTester))
}

func (suite *SuiteTester) BeforeTest(suiteName, testName string) {}

func (suite *SuiteTester) AfterTest(suiteName, testName string) {}

func (s *SuiteTester) SetupSuite() {
	databaseEnv := env.Database{
		UserName: "postgres",
		Password: "123456",
		Name:     "jumia_test",
		Port:     "5432",
		Host:     "localhost",
	}
	s.listRepository = &Instance{Logger: logging.Log, ORM: database.Connect(databaseEnv)}
	fmt.Println("🧹🧹🧹 Cleaning Table ")
	s.listRepository.ClearTable()
}

func (suite *SuiteTester) SetupTest() {

}

func (suite *SuiteTester) Test1CreateList() {
	createdList = &entity.List{Name: "Test"}
	id, err := suite.listRepository.Create(createdList)
	suite.Assert().Nil(err)
	createdList.ID = int64(id)
}

func (suite *SuiteTester) Test2FailCreateList() {
	_, err := suite.listRepository.Create(&entity.List{Name: "Test"})
	suite.Assert().NotNil(err)
}

func (suite *SuiteTester) Test3UpdateList() {
	updatedList := &entity.List{Name: "Test1", ID: 1}

	err := suite.listRepository.Update(updatedList, 1)
	suite.Assert().Nil(err)

	createdList.Name = updatedList.Name
}

func (suite *SuiteTester) Test4UpdateListFailure() {

	id, err := suite.listRepository.Create(&entity.List{Name: "New"})
	suite.Assert().Nil(err)

	updatedList := &entity.List{Name: "Test1", ID: int64(id)}

	err = suite.listRepository.Update(updatedList, id)

	suite.Assert().NotNil(err)

}

func (suite *SuiteTester) Test5GetList() {
	list, err := suite.listRepository.GetAllLists(10, 0)
	suite.Assert().Nil(err)
	suite.Assert().Equal(list[0].ID, createdList.ID)

}

func (suite *SuiteTester) Test6DeleteList() {
	err := suite.listRepository.Delete(int(createdList.ID))
	suite.Assert().Nil(err)
}

func (suite *SuiteTester) Test7DeleteListDoesnotExists() {
	err := suite.listRepository.Delete(100)
	suite.Assert().Nil(err)
}
