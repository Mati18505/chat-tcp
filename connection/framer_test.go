package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderMarshall(t *testing.T) {
	header := header{
		size: 15,
	}
	bytes := marshallHeader(header)
	unmarshalled := unmarshallHeader(bytes)

	assert.Equal(t, header, unmarshalled)
}