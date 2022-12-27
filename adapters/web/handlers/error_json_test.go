package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
