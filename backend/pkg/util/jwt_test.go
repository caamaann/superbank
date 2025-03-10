package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	
	secret := "test-secret"
	userID := "1"
	role := "admin"
	expiration := time.Hour
	
	
	token, err := GenerateJWT(secret, userID, role, expiration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	
	
	gotUserID, gotRole, err := ValidateJWT(secret, token)
	assert.NoError(t, err)
	assert.Equal(t, userID, gotUserID)
	assert.Equal(t, role, gotRole)
	
	
	_, _, err = ValidateJWT(secret, "invalid-token")
	assert.Error(t, err)
	
	
	_, _, err = ValidateJWT("wrong-secret", token)
	assert.Error(t, err)
}