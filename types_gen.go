// Code generated by ask-gen. DO NOT EDIT!!

package ask

import (
	"errors"
	"strconv"
)

// VarType indicates a type of a variable.
type VarType string

func (t VarType) String() string {
	return string(t)
}

const (
	// VarTypeString indicates the string
	VarTypeString = VarType("string")

	// VarTypeBool indicates the bool
	VarTypeBool = VarType("bool")

	// VarTypeYesNo indicates the bool
	VarTypeYesNo = VarType("yesNo")

	// VarTypeUint indicates the uint
	VarTypeUint = VarType("uint")

	// VarTypeUint8 indicates the uint8
	VarTypeUint8 = VarType("uint8")

	// VarTypeUint16 indicates the uint16
	VarTypeUint16 = VarType("uint16")

	// VarTypeUint32 indicates the uint32
	VarTypeUint32 = VarType("uint32")

	// VarTypeUint64 indicates the uint64
	VarTypeUint64 = VarType("uint64")

	// VarTypeInt indicates the int
	VarTypeInt = VarType("int")

	// VarTypeInt8 indicates the int8
	VarTypeInt8 = VarType("int8")

	// VarTypeInt16 indicates the int16
	VarTypeInt16 = VarType("int16")

	// VarTypeInt32 indicates the int32
	VarTypeInt32 = VarType("int32")

	// VarTypeInt64 indicates the int64
	VarTypeInt64 = VarType("int64")

	// VarTypeFloat32 indicates the float32
	VarTypeFloat32 = VarType("float32")

	// VarTypeFloat64 indicates the float64
	VarTypeFloat64 = VarType("float64")
)

// VarTypes is the all types.
func VarTypes() []VarType {
	return []VarType{
		VarTypeString,

		VarTypeBool,

		VarTypeYesNo,

		VarTypeUint,

		VarTypeUint8,

		VarTypeUint16,

		VarTypeUint32,

		VarTypeUint64,

		VarTypeInt,

		VarTypeInt8,

		VarTypeInt16,

		VarTypeInt32,

		VarTypeInt64,

		VarTypeFloat32,

		VarTypeFloat64,
	}
}

// Interface will get a value that type of 't'.
func (s Service) Interface(t VarType) (interface{}, error) {
	switch t {
	case VarTypeString:
		return s.String()

	case VarTypeBool:
		return s.Bool()

	case VarTypeYesNo:
		return s.YesNo()

	case VarTypeUint:
		return s.Uint()

	case VarTypeUint8:
		return s.Uint8()

	case VarTypeUint16:
		return s.Uint16()

	case VarTypeUint32:
		return s.Uint32()

	case VarTypeUint64:
		return s.Uint64()

	case VarTypeInt:
		return s.Int()

	case VarTypeInt8:
		return s.Int8()

	case VarTypeInt16:
		return s.Int16()

	case VarTypeInt32:
		return s.Int32()

	case VarTypeInt64:
		return s.Int64()

	case VarTypeFloat32:
		return s.Float32()

	case VarTypeFloat64:
		return s.Float64()
	}
	return nil, errors.New("invalid type")
}

// Interface will get a value that type of 't'.
func Interface(t VarType) (interface{}, error) {
	return static.Interface(t)
}

// String takes string value from user input
func (s Service) String() (*string, error) {
	var v string
	if err := s.StringVar(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// String takes string value from user input
func String() (*string, error) {
	return static.String()
}

// StringVar sets a string variable, "v" to accept user input
func (s Service) StringVar(v *string) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			*v = input
			return nil
		})
	})
}

// StringVar sets a string variable, "v" to accept user input
func StringVar(v *string) Doer {
	return static.StringVar(v)
}

// Bool takes bool value from user input
func (s Service) Bool() (*bool, error) {
	var v bool
	if err := s.BoolVar(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Bool takes bool value from user input
func Bool() (*bool, error) {
	return static.Bool()
}

// BoolVar sets a bool variable, "v" to accept user input
func (s Service) BoolVar(v *bool) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseBool(input)
			if err != nil {
				return err
			}
			*v = bool(p)
			return nil
		})
	})
}

// BoolVar sets a bool variable, "v" to accept user input
func BoolVar(v *bool) Doer {
	return static.BoolVar(v)
}

// YesNo takes bool value from user input
func (s Service) YesNo() (*bool, error) {
	var v bool
	if err := s.YesNoVar(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// YesNo takes bool value from user input
func YesNo() (*bool, error) {
	return static.YesNo()
}

// YesNoVar sets a bool variable, "v" to accept user input
func (s Service) YesNoVar(v *bool) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := parseYesNo(input)
			if err != nil {
				return err
			}
			*v = bool(p)
			return nil
		})
	})
}

// YesNoVar sets a bool variable, "v" to accept user input
func YesNoVar(v *bool) Doer {
	return static.YesNoVar(v)
}

// Uint takes uint value from user input
func (s Service) Uint() (*uint, error) {
	var v uint
	if err := s.UintVar(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Uint takes uint value from user input
func Uint() (*uint, error) {
	return static.Uint()
}

// UintVar sets a uint variable, "v" to accept user input
func (s Service) UintVar(v *uint) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, strconv.IntSize)
			if err != nil {
				return err
			}
			*v = uint(p)
			return nil
		})
	})
}

// UintVar sets a uint variable, "v" to accept user input
func UintVar(v *uint) Doer {
	return static.UintVar(v)
}

// Uint8 takes uint8 value from user input
func (s Service) Uint8() (*uint8, error) {
	var v uint8
	if err := s.Uint8Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Uint8 takes uint8 value from user input
func Uint8() (*uint8, error) {
	return static.Uint8()
}

// Uint8Var sets a uint8 variable, "v" to accept user input
func (s Service) Uint8Var(v *uint8) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 8)
			if err != nil {
				return err
			}
			*v = uint8(p)
			return nil
		})
	})
}

// Uint8Var sets a uint8 variable, "v" to accept user input
func Uint8Var(v *uint8) Doer {
	return static.Uint8Var(v)
}

// Uint16 takes uint16 value from user input
func (s Service) Uint16() (*uint16, error) {
	var v uint16
	if err := s.Uint16Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Uint16 takes uint16 value from user input
func Uint16() (*uint16, error) {
	return static.Uint16()
}

// Uint16Var sets a uint16 variable, "v" to accept user input
func (s Service) Uint16Var(v *uint16) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 16)
			if err != nil {
				return err
			}
			*v = uint16(p)
			return nil
		})
	})
}

// Uint16Var sets a uint16 variable, "v" to accept user input
func Uint16Var(v *uint16) Doer {
	return static.Uint16Var(v)
}

// Uint32 takes uint32 value from user input
func (s Service) Uint32() (*uint32, error) {
	var v uint32
	if err := s.Uint32Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Uint32 takes uint32 value from user input
func Uint32() (*uint32, error) {
	return static.Uint32()
}

// Uint32Var sets a uint32 variable, "v" to accept user input
func (s Service) Uint32Var(v *uint32) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 32)
			if err != nil {
				return err
			}
			*v = uint32(p)
			return nil
		})
	})
}

// Uint32Var sets a uint32 variable, "v" to accept user input
func Uint32Var(v *uint32) Doer {
	return static.Uint32Var(v)
}

// Uint64 takes uint64 value from user input
func (s Service) Uint64() (*uint64, error) {
	var v uint64
	if err := s.Uint64Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Uint64 takes uint64 value from user input
func Uint64() (*uint64, error) {
	return static.Uint64()
}

// Uint64Var sets a uint64 variable, "v" to accept user input
func (s Service) Uint64Var(v *uint64) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 64)
			if err != nil {
				return err
			}
			*v = uint64(p)
			return nil
		})
	})
}

// Uint64Var sets a uint64 variable, "v" to accept user input
func Uint64Var(v *uint64) Doer {
	return static.Uint64Var(v)
}

// Int takes int value from user input
func (s Service) Int() (*int, error) {
	var v int
	if err := s.IntVar(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Int takes int value from user input
func Int() (*int, error) {
	return static.Int()
}

// IntVar sets a int variable, "v" to accept user input
func (s Service) IntVar(v *int) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, strconv.IntSize)
			if err != nil {
				return err
			}
			*v = int(p)
			return nil
		})
	})
}

// IntVar sets a int variable, "v" to accept user input
func IntVar(v *int) Doer {
	return static.IntVar(v)
}

// Int8 takes int8 value from user input
func (s Service) Int8() (*int8, error) {
	var v int8
	if err := s.Int8Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Int8 takes int8 value from user input
func Int8() (*int8, error) {
	return static.Int8()
}

// Int8Var sets a int8 variable, "v" to accept user input
func (s Service) Int8Var(v *int8) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 8)
			if err != nil {
				return err
			}
			*v = int8(p)
			return nil
		})
	})
}

// Int8Var sets a int8 variable, "v" to accept user input
func Int8Var(v *int8) Doer {
	return static.Int8Var(v)
}

// Int16 takes int16 value from user input
func (s Service) Int16() (*int16, error) {
	var v int16
	if err := s.Int16Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Int16 takes int16 value from user input
func Int16() (*int16, error) {
	return static.Int16()
}

// Int16Var sets a int16 variable, "v" to accept user input
func (s Service) Int16Var(v *int16) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 16)
			if err != nil {
				return err
			}
			*v = int16(p)
			return nil
		})
	})
}

// Int16Var sets a int16 variable, "v" to accept user input
func Int16Var(v *int16) Doer {
	return static.Int16Var(v)
}

// Int32 takes int32 value from user input
func (s Service) Int32() (*int32, error) {
	var v int32
	if err := s.Int32Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Int32 takes int32 value from user input
func Int32() (*int32, error) {
	return static.Int32()
}

// Int32Var sets a int32 variable, "v" to accept user input
func (s Service) Int32Var(v *int32) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return err
			}
			*v = int32(p)
			return nil
		})
	})
}

// Int32Var sets a int32 variable, "v" to accept user input
func Int32Var(v *int32) Doer {
	return static.Int32Var(v)
}

// Int64 takes int64 value from user input
func (s Service) Int64() (*int64, error) {
	var v int64
	if err := s.Int64Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Int64 takes int64 value from user input
func Int64() (*int64, error) {
	return static.Int64()
}

// Int64Var sets a int64 variable, "v" to accept user input
func (s Service) Int64Var(v *int64) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return err
			}
			*v = int64(p)
			return nil
		})
	})
}

// Int64Var sets a int64 variable, "v" to accept user input
func Int64Var(v *int64) Doer {
	return static.Int64Var(v)
}

// Float32 takes float32 value from user input
func (s Service) Float32() (*float32, error) {
	var v float32
	if err := s.Float32Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Float32 takes float32 value from user input
func Float32() (*float32, error) {
	return static.Float32()
}

// Float32Var sets a float32 variable, "v" to accept user input
func (s Service) Float32Var(v *float32) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseFloat(input, 32)
			if err != nil {
				return err
			}
			*v = float32(p)
			return nil
		})
	})
}

// Float32Var sets a float32 variable, "v" to accept user input
func Float32Var(v *float32) Doer {
	return static.Float32Var(v)
}

// Float64 takes float64 value from user input
func (s Service) Float64() (*float64, error) {
	var v float64
	if err := s.Float64Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// Float64 takes float64 value from user input
func Float64() (*float64, error) {
	return static.Float64()
}

// Float64Var sets a float64 variable, "v" to accept user input
func (s Service) Float64Var(v *float64) Doer {
	return doFunc(func() error {
		return s.AskParseFunc(func(input string) error {
			p, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return err
			}
			*v = float64(p)
			return nil
		})
	})
}

// Float64Var sets a float64 variable, "v" to accept user input
func Float64Var(v *float64) Doer {
	return static.Float64Var(v)
}
