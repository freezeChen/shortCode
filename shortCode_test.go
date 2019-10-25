package shortCode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultShortCode(t *testing.T) {
	shortCode, err := NewShortCode(CodeLength(6))
	assert.NoError(t, err)

	var id int64 = 592044735
	code, err := shortCode.Encode(id)
	fmt.Println(code)
	assert.NoError(t, err)

	sid, err := shortCode.Decode(code)
	assert.NoError(t, err)

	assert.Equal(t, id, sid)

}

func TestCustomShortCode(t *testing.T) {
	shortCode, err := NewShortCode(SeedsIndex(2), CodeLength(6))
	assert.NoError(t, err)

	var code = "2p5ppp"

	decode, err := shortCode.Decode(code)
	assert.NoError(t, err)

	encode, err := shortCode.Encode(decode)
	assert.NoError(t, err)

	assert.Equal(t, code, encode)

}
