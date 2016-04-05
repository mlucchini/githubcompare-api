package auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGithubClientId(t *testing.T) {
	clientId := GithubClientId()
	assert.NotNil(t, clientId)
}

func TestGithubClientSecret(t *testing.T) {
	clientSecret := GithubClientSecret()
	assert.NotNil(t, clientSecret)
}