package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Pablo Henrique", "Pablo@gmail.com", "Pablo123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.Equal(t, "Pablo Henrique", user.Name)
	assert.Equal(t, "Pablo@gmail.com", user.Email)
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := NewUser("Pablo Henrique", "Pablo@gmail.com", "Pablo123")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("Pablo123"))
	assert.False(t, user.CheckPassword("pablo12234"))
	assert.NotEqual(t, "Pablo123", user.Password)
}
