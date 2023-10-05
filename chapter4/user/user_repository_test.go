package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := NewMockUserRepository(ctrl)

	expectedUser := UserDetails{ID: "1", Name: "John Doe"}
	mockRepository.EXPECT().GetUserByID("1").Return(expectedUser, nil)

	// Call the function being tested with the mock dependency
	user, err := mockRepository.GetUserByID("1")

	// Assert the expected results
	assert.Equal(t, expectedUser, user)
	assert.Nil(t, err)
}
