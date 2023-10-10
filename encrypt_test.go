package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	key := ParseKey("303132333435363738393031323" +
		"3343536373839303132333435363738393031")

	t.Run("sometext", func(t *testing.T) {
		plain := "sometext"

		enc := key.Encrypt(plain)
		dec := key.Decrypt(enc)

		require.Equal(t, dec, plain)
	})

	t.Run("empty", func(t *testing.T) {
		plain := ""

		enc := key.Encrypt(plain)
		dec := key.Decrypt(enc)

		require.Equal(t, dec, plain)
	})
}
