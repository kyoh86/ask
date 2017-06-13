// Code generated by ask-gen. DO NOT EDIT!!

package ask

import "strconv"

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

// String sets a string variable, "v" to accept user input
func (s Service) StringVar(v *string) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			*v = input
			return nil
		})
	})
}

// String sets a string variable, "v" to accept user input
func StringVar(v *string) Doer {
	return StringVar(v)
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

// Bool sets a bool variable, "v" to accept user input
func (s Service) BoolVar(v *bool) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseBool(input)
			if err != nil {
				return err
			}
			*v = bool(p)
			return nil
		})
	})
}

// Bool sets a bool variable, "v" to accept user input
func BoolVar(v *bool) Doer {
	return BoolVar(v)
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

// Uint sets a uint variable, "v" to accept user input
func (s Service) UintVar(v *uint) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, strconv.IntSize)
			if err != nil {
				return err
			}
			*v = uint(p)
			return nil
		})
	})
}

// Uint sets a uint variable, "v" to accept user input
func UintVar(v *uint) Doer {
	return UintVar(v)
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

// Uint8 sets a uint8 variable, "v" to accept user input
func (s Service) Uint8Var(v *uint8) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 8)
			if err != nil {
				return err
			}
			*v = uint8(p)
			return nil
		})
	})
}

// Uint8 sets a uint8 variable, "v" to accept user input
func Uint8Var(v *uint8) Doer {
	return Uint8Var(v)
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

// Uint16 sets a uint16 variable, "v" to accept user input
func (s Service) Uint16Var(v *uint16) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 16)
			if err != nil {
				return err
			}
			*v = uint16(p)
			return nil
		})
	})
}

// Uint16 sets a uint16 variable, "v" to accept user input
func Uint16Var(v *uint16) Doer {
	return Uint16Var(v)
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

// Uint32 sets a uint32 variable, "v" to accept user input
func (s Service) Uint32Var(v *uint32) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 32)
			if err != nil {
				return err
			}
			*v = uint32(p)
			return nil
		})
	})
}

// Uint32 sets a uint32 variable, "v" to accept user input
func Uint32Var(v *uint32) Doer {
	return Uint32Var(v)
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

// Uint64 sets a uint64 variable, "v" to accept user input
func (s Service) Uint64Var(v *uint64) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseUint(input, 10, 64)
			if err != nil {
				return err
			}
			*v = uint64(p)
			return nil
		})
	})
}

// Uint64 sets a uint64 variable, "v" to accept user input
func Uint64Var(v *uint64) Doer {
	return Uint64Var(v)
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

// Int sets a int variable, "v" to accept user input
func (s Service) IntVar(v *int) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, strconv.IntSize)
			if err != nil {
				return err
			}
			*v = int(p)
			return nil
		})
	})
}

// Int sets a int variable, "v" to accept user input
func IntVar(v *int) Doer {
	return IntVar(v)
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

// Int8 sets a int8 variable, "v" to accept user input
func (s Service) Int8Var(v *int8) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 8)
			if err != nil {
				return err
			}
			*v = int8(p)
			return nil
		})
	})
}

// Int8 sets a int8 variable, "v" to accept user input
func Int8Var(v *int8) Doer {
	return Int8Var(v)
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

// Int16 sets a int16 variable, "v" to accept user input
func (s Service) Int16Var(v *int16) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 16)
			if err != nil {
				return err
			}
			*v = int16(p)
			return nil
		})
	})
}

// Int16 sets a int16 variable, "v" to accept user input
func Int16Var(v *int16) Doer {
	return Int16Var(v)
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

// Int32 sets a int32 variable, "v" to accept user input
func (s Service) Int32Var(v *int32) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return err
			}
			*v = int32(p)
			return nil
		})
	})
}

// Int32 sets a int32 variable, "v" to accept user input
func Int32Var(v *int32) Doer {
	return Int32Var(v)
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

// Int64 sets a int64 variable, "v" to accept user input
func (s Service) Int64Var(v *int64) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				return err
			}
			*v = int64(p)
			return nil
		})
	})
}

// Int64 sets a int64 variable, "v" to accept user input
func Int64Var(v *int64) Doer {
	return Int64Var(v)
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

// Float32 sets a float32 variable, "v" to accept user input
func (s Service) Float32Var(v *float32) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseFloat(input, 32)
			if err != nil {
				return err
			}
			*v = float32(p)
			return nil
		})
	})
}

// Float32 sets a float32 variable, "v" to accept user input
func Float32Var(v *float32) Doer {
	return Float32Var(v)
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

// Float64 sets a float64 variable, "v" to accept user input
func (s Service) Float64Var(v *float64) Doer {
	return DoFunc(func() error {
		return s.AskFunc(func(input string) error {
			p, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return err
			}
			*v = float64(p)
			return nil
		})
	})
}

// Float64 sets a float64 variable, "v" to accept user input
func Float64Var(v *float64) Doer {
	return Float64Var(v)
}
