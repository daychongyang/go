package jenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	rawURL     = "http://root:123456@localhost:8080/jenkins"
	jenkinsURL = "http://localhost:8080/jenkins/"
)

func TestGetInfo(t *testing.T) {
	jenkins, err := New(rawURL)
	assert.Nil(t, err)

	info, err := jenkins.Info()

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, jenkinsURL, info.URL)
}