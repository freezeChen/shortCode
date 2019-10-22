package shortCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortCode(t *testing.T) {
	shortCode := NewShortCode()
	var id int64 = 1345678536
	// fmt.
	code, err := shortCode.Encode(id)
	if err != nil {
		t.Error(err)
	}
	sid, err := shortCode.Decode(code)
	if err != nil {
		t.Error(err)

	}

	assert.Equal(t, id, sid)

}

func TestCustomShortCode(t *testing.T) {
	// shortCode := NewShortCode(SeedsIndex(2), Length(6))
	// var id
	// shortCode.Encode()
}
