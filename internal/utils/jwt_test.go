package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	jwt, err := GenerateJWT("shubhamdixit")
	log.Println(jwt)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(jwt), 0)
}
