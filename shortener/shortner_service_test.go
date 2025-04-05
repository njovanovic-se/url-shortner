package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://medium.com/@nikjov92/effectively-separate-ef-core-migrations-into-a-separate-project-net-8-2d84843d17da"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	assert.Equal(t, shortLink_1, "DyJwQAEs")
}
