package service

import (
	"testing"

	"github.com/AndresJejen/GoPersonalServer/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entities.Post) (*entities.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.Post), args.Error(1)
}

func (mock *mockRepository) FindAll() ([]entities.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Post), args.Error(1)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)

	post := entities.Post{Title: "Title Saved", Text: "Texting"}

	// Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "Title Saved", result.Title)
	assert.Equal(t, "Texting", result.Text)
	assert.Nil(t, err)

}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var identifier int64 = 1

	post := entities.Post{ID: identifier, Title: "Title Saved", Text: "Texting"}

	// setup expectations
	mockRepo.On("FindAll").Return([]entities.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, "Title Saved", result[0].Title)
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Texting", result[0].Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The Post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {

	post := entities.Post{Title: "", Text: "Texting"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The Post Titile is empty", err.Error())
}
