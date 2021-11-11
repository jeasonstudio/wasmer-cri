package utils_test

import (
	"testing"

	"github.com/jeasonstudio/wasmer-cri/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id := utils.GenerateID()
	assert.NotNil(t, id)
	assert.Equal(t, len(id), 64)
}
