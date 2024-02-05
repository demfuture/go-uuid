package uuid

import (
	"testing"
)

func TestBase64(t *testing.T) {

	t.Run("Url base64 encode", func(t *testing.T) {
		id := Must(FromString("69359037-9599-48e7-b8f2-48393c019135"))

		encoded := id.Base64URL()

		expected := "aTWQN5WZSOe48kg5PAGRNQ"

		if encoded != expected {
			t.Errorf("Expected %s but got %s", expected, encoded)
		}
	})

	t.Run("Url base64 decode", func(t *testing.T) {
		id := "aTWQN5WZSOe48kg5PAGRNQ"

		encoded, err := FromBase64URL(id)
		if err != nil {
			t.Error(err)
		}

		expected := Must(FromString("69359037-9599-48e7-b8f2-48393c019135"))

		if encoded != expected {
			t.Errorf("Expected %s but got %s", expected, encoded)
		}
	})

	t.Run("Continuous base 64 encode", func(t *testing.T) {
		id := Must(FromString("69356037-9599-48e7-b8f2-48393c019135"))

		encoded := id.ContinuousBase64()

		expected := "aTVgN5WZSOe48kg5PAGRNQ"

		if encoded != expected {
			t.Errorf("Expected %s but got %s", expected, encoded)
		}
	})

	t.Run("Continuous base64 decode", func(t *testing.T) {
		id := "Ej5FZ90ibEtOkVkJmVUQAAA"

		encoded, err := FromContinuousBase64(id)
		if err != nil {
			t.Error(err)
		}

		expected := Must(FromString("123e4567-e89b-12d3-a456-426655440000"))

		if encoded != expected {
			t.Errorf("Expected %s but got %s", expected, encoded)
		}
	})

}
