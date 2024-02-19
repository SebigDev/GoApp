package vo

import (
	"CrashCourse/GoApp/internal/utils"
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type PinError struct {
	ErrorMsg string
}

func (e PinError) Error() string {
	return fmt.Sprintf(e.ErrorMsg)
}

type Value string

func (p *Value) String() string {
	return string(*p)
}

func NewValue(val string) (Value, error) {
	if utils.Length(val) == 0 {
		return "", PinError{ErrorMsg: "Pin value must be provided"}
	}
	if utils.Length(val) < 4 || utils.Length(val) > 4 {
		return "", PinError{ErrorMsg: "Pin value must have a length of 4"}
	}
	return Value(val), nil
}

type Pin struct {
	ValueHash    []byte
	RecoverValue string
}

func NewPin(pin Value) *Pin {
	return &Pin{
		ValueHash:    []byte(pin),
		RecoverValue: uuid.NewString(),
	}
}

func Verify(pin string, p Pin) error {
	cleanPin, err := NewValue(pin)
	if err != nil {
		return err
	}
	hashedPin := []byte(cleanPin)
	// for i, b := range p.ValueHash {
	// 	if b != hashedPin[i] {
	// 		return PinError{ErrorMsg: "You have provided an invalid PIN"}
	// 	}
	// }
	if slices.Equal(p.ValueHash, hashedPin) {
		return nil
	}
	return PinError{ErrorMsg: "You have provided an invalid PIN"}
}

func VerifyRecover(recoveryPin string, p Pin) error {
	if utils.Length(recoveryPin) == 0 {
		return PinError{ErrorMsg: "You have provided an invalid or empty recovery PIN"}
	}
	if recoveryPin != p.RecoverValue {
		return PinError{ErrorMsg: "You have provided an invalid recovery PIN"}
	}
	return nil
}