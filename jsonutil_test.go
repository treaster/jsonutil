package jsonutil_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/treaster/jsonutil"
)

func TestSafeUnmarshal(t *testing.T) {
	type S struct {
		Field1 int
		Field2 string
		// Field3 is unknown
	}

	{
		// Verify that the SafeUnmarshal can Unmarshal JSON with no unknown fields.

		input := []byte(`{
			"Field1": 10,
			"Field2": "abc"
		}`)

		var output S
		err := jsonutil.SafeUnmarshal(input, &output)
		require.NoError(t, err)
	}

	{
		input := []byte(`{
			"Field1": 10,
			"Field2": "abc",
			"Field3": "def"
		}`)

		// Verify that the built-in JSON Unmarshal will not produce an error
		{
			var output S
			err := json.Unmarshal(input, &output)
			require.NoError(t, err)
		}

		// Verify that the SafeUnmarshal does produce an error due to the extra
		// field in the input.
		{
			var output S
			err := jsonutil.SafeUnmarshal(input, &output)
			require.Error(t, err)
		}
	}
}

func TestMustMarshal(t *testing.T) {
	{
		// Verify that the MustMarshal can marshal clean JSON successfully.
		input := json.RawMessage(`{
			"Field1": 10,
			"Field2": "abc"
		}`)

		_ = jsonutil.MustMarshal(input)
	}

	{
		// Verify that the SafeUnmarshal panics when presented with an unmarshalable object.
		input := json.RawMessage(`{
			"Field1": 10,
			"Field2": "abc" this is an error
		}`)

		require.Panics(t, func() {
			_ = jsonutil.MustMarshal(input)
		})
	}
}
