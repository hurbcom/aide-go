package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	actual := Join(", ", 654321987, "bar", 654.654)
	assert.Equal(t, `654321987, bar, 654.654`, actual)

	actual = Join(", ", int64(654321987), "bar")
	assert.Equal(t, `654321987, bar`, actual)

	actual = Join(", ", int64(654321987), int64(52354))
	assert.Equal(t, `654321987, 52354`, actual)

	actual = Join(", ", "foo")
	assert.Equal(t, `foo`, actual)

	actual = Join(", ", []string{"foo", "bar"})
	assert.Equal(t, `foo, bar`, actual)

	actual = Join(", ", []int{65485, 19734})
	assert.Equal(t, `65485, 19734`, actual)

	actual = Join("_", []string{"foo", "bar", ""}, nil, 123)
	assert.Equal(t, `foo_bar_123`, actual)

	var pStr *string
	str := "foo"

	actual = Join(", ", 654321987, nil, 654.654, "", pStr, &str)
	assert.Equal(t, `654321987, 654.654, foo`, actual)
}

func TestTruncate(t *testing.T) {
	x := "123456789abcdef"
	assert.Equal(t, x, Truncate(x, 16))
	assert.Equal(t, x[:15], Truncate(x, 15))
	assert.Equal(t, "12", Truncate(x, 2))

	expected := "{123456789abcdef}"
	json1 := `{
        123456789abcdef
    }`
	assert.Equal(t, expected, Truncate(json1, len(json1)))
}
