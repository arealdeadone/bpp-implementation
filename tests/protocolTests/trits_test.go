package protocolTests

import (
	"bpp/protocol"
	"testing"
)

func TestGenerateRandomKey_Length(t *testing.T){
	var length = uint64(32)
	generatedKey := protocol.GenerateRandomKey(length)

	if uint64(len(generatedKey)) != length {
		t.Errorf("Expected length of the string to be %d but got %d\r\n", length, len(generatedKey))
	}
}


