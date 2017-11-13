package dialex

import (
        "os"
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
        dial := New(os.Getenv("API_KEY"))
        out, err := dial.Transform("diff column", "en")
        assert.Nil(t, err)
        assert.Equal(t, out, "difference column", "should be equal to difference column")
}

func TestInvalidKey(t *testing.T) {
        dial := New("")
        assert.Panics(t, func() { dial.Transform("diff column", "en") })
}
