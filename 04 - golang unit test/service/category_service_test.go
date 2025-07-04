package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	// ===== Program Mock =====
	// menyimulasikan pengambilan Category dengan ID 1, dimana tidak ada di database
	categoryRepository.Mock.On(
		"FindById",
		"1",
	).Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category) // Harus Nil, karena kita simulasikan kalo tidak ada Category dengan ID 1
	assert.NotNil(t, err)   // Harus ada error
}

func TestCategoryService_GetFound(t *testing.T) {
	var category = entity.Category{
		Id:   "2",
		Name: "Handphone",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)                          // Harusnya tidak ada error
	assert.NotNil(t, result)                    // Harusnya ada hasilnya berupa category
	assert.Equal(t, category.Id, result.Id)     // Harusnya sama
	assert.Equal(t, category.Name, result.Name) // Harusnya juga sama
}
