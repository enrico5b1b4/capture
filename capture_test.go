package capture_test

import (
	"testing"

	"github.com/enrico5b1b4/capture"
	"github.com/stretchr/testify/assert"
)

func TestCapture_Parse_String(t *testing.T) {
	type reminder struct {
		Who     string `regexpGroup:"who"`
		Message string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) to (?P<message>.*)`,
		"remind John to buy milk",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, "buy milk", myReminder.Message)
}

func TestCapture_Parse_PointerString(t *testing.T) {
	type reminder struct {
		Who     string  `regexpGroup:"who"`
		Message *string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) to (?P<message>.*)`,
		"remind John to buy milk",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, "buy milk", *myReminder.Message)
}

func TestCapture_Parse_NilPointerString(t *testing.T) {
	type reminder struct {
		Who     string  `regexpGroup:"who"`
		Message *string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) ?(?P<message>.*)?`,
		"remind John",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Nil(t, myReminder.Message)
}

func TestCapture_Parse_ZeroValueString(t *testing.T) {
	type reminder struct {
		Who     string `regexpGroup:"who"`
		Message string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) to ?(?P<message>.*)?`,
		"remind John to",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, "", myReminder.Message)
}

func TestCapture_Parse_Int(t *testing.T) {
	type reminder struct {
		Who     string `regexpGroup:"who"`
		Day     int    `regexpGroup:"day"`
		Month   string `regexpGroup:"month"`
		Year    int    `regexpGroup:"year"`
		Message string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) on the (?P<day>\d{1,2})(?:(st|nd|rd|th))? of (?P<month>\w+) (?P<year>\d{4}) to (?P<message>.*)`,
		"remind John on the 31st of october 2030 to buy milk",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, 31, myReminder.Day)
	assert.Equal(t, "october", myReminder.Month)
	assert.Equal(t, 2030, myReminder.Year)
	assert.Equal(t, "buy milk", myReminder.Message)
}

func TestCapture_Parse_PointerInt(t *testing.T) {
	type reminder struct {
		Who     string `regexpGroup:"who"`
		Day     *int   `regexpGroup:"day"`
		Month   string `regexpGroup:"month"`
		Year    int    `regexpGroup:"year"`
		Message string `regexpGroup:"message"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) on the (?P<day>\d{1,2})(?:(st|nd|rd|th))? of (?P<month>\w+) (?P<year>\d{4}) to (?P<message>.*)`,
		"remind John on the 31st of october 2030 to buy milk",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, 31, *myReminder.Day)
	assert.Equal(t, "october", myReminder.Month)
	assert.Equal(t, 2030, myReminder.Year)
	assert.Equal(t, "buy milk", myReminder.Message)
}

func TestCapture_Parse_NilPointerInt(t *testing.T) {
	type reminder struct {
		Who string `regexpGroup:"who"`
		Day *int   `regexpGroup:"day"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) ?(?P<day>\d{1,2})? (?P<message>.*)`,
		"remind John something",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Nil(t, myReminder.Day)
}

func TestCapture_Parse_ZeroValueInt(t *testing.T) {
	type reminder struct {
		Who string `regexpGroup:"who"`
		Day int    `regexpGroup:"day"`
	}

	myReminder := &reminder{}
	err := capture.Parse(
		`remind (?P<who>\w+) ?(?P<day>\d{1,2})? (?P<message>.*)`,
		"remind John something",
		myReminder,
	)

	assert.NoError(t, err)
	assert.Equal(t, "John", myReminder.Who)
	assert.Equal(t, 0, myReminder.Day)
}

func TestCapture_Parse_IntError(t *testing.T) {
	type WakeMe struct {
		Month int `regexpGroup:"month"`
	}

	myWakeMessage := &WakeMe{}
	err := capture.Parse(
		`wake me up when (?P<month>september|october|november|december) ends`,
		"wake me up when september ends",
		myWakeMessage,
	)

	assert.Error(t, err)
}

func TestCapture_Parse_Bool(t *testing.T) {
	type message struct {
		Field string `regexpGroup:"field"`
		Value bool   `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) to (?P<value>true|false)`,
		"set A to true",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Equal(t, true, myMessage.Value)
}

func TestCapture_Parse_PointerBool(t *testing.T) {
	type message struct {
		Field string `regexpGroup:"field"`
		Value *bool  `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) to (?P<value>true|false)`,
		"set A to true",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Equal(t, true, *myMessage.Value)
}

func TestCapture_Parse_NilPointerBool(t *testing.T) {
	type message struct {
		Field string `regexpGroup:"field"`
		Value *bool  `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) ?(?P<value>true|false)?`,
		"set A",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Nil(t, myMessage.Value)
}

func TestCapture_Parse_ZeroValueBool(t *testing.T) {
	type message struct {
		Field string `regexpGroup:"field"`
		Value bool   `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) ?(?P<value>true|false)?`,
		"set A",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.False(t, myMessage.Value)
}

func TestCapture_Parse_BoolError(t *testing.T) {
	type WakeMe struct {
		Month bool `regexpGroup:"month"`
	}

	myWakeMessage := &WakeMe{}
	err := capture.Parse(
		`wake me up when (?P<month>september|october|november|december) ends`,
		"wake me up when september ends",
		myWakeMessage,
	)

	assert.Error(t, err)
}

func TestCapture_Parse_Float64(t *testing.T) {
	type message struct {
		Field string  `regexpGroup:"field"`
		Value float64 `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) to (?P<value>[+-]?([0-9]*[.])?[0-9]+)`,
		"set A to 3.14",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Equal(t, 3.14, myMessage.Value)
}

func TestCapture_Parse_PointerFloat64(t *testing.T) {
	type message struct {
		Field string   `regexpGroup:"field"`
		Value *float64 `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) to (?P<value>[+-]?([0-9]*[.])?[0-9]+)`,
		"set A to 3.14",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Equal(t, 3.14, *myMessage.Value)
}

func TestCapture_Parse_NilPointerFloat64(t *testing.T) {
	type message struct {
		Field string   `regexpGroup:"field"`
		Value *float64 `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) ?(?P<value>[+-]?([0-9]*[.])?[0-9]+)?`,
		"set A",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Nil(t, myMessage.Value)
}

func TestCapture_Parse_ZeroValueFloat64(t *testing.T) {
	type message struct {
		Field string  `regexpGroup:"field"`
		Value float64 `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) ?(?P<value>[+-]?([0-9]*[.])?[0-9]+)?`,
		"set A",
		myMessage,
	)

	assert.NoError(t, err)
	assert.Equal(t, "A", myMessage.Field)
	assert.Equal(t, float64(0), myMessage.Value)
}

func TestCapture_Parse_Float64Error(t *testing.T) {
	type message struct {
		Field string  `regexpGroup:"field"`
		Value float64 `regexpGroup:"value"`
	}

	myMessage := &message{}
	err := capture.Parse(
		`set (?P<field>\w+) to (?P<value>.*)`,
		"set A to blue",
		myMessage,
	)

	assert.Error(t, err)
}

func TestCapture_Parse_NotStructError(t *testing.T) {
	message := "test"

	err := capture.Parse(
		`wake me up when (?P<month>september|october|november|december) ends`,
		"wake me up when september ends",
		&message,
	)

	assert.Error(t, err)
}
