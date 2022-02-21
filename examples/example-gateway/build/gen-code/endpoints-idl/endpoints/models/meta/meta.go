// Code generated by thriftrw v1.29.2. DO NOT EDIT.
// @generated

package meta

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

type Dgx struct {
	S1 string `json:"s1,required"`
	I2 int32  `json:"i2,required"`
	B3 *bool  `json:"b3,omitempty"`
}

// ToWire translates a Dgx struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Dgx) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.S1), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueI32(v.I2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.B3 != nil {
		w, err = wire.NewValueBool(*(v.B3)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Dgx struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dgx struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dgx
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dgx) FromWire(w wire.Value) error {
	var err error

	s1IsSet := false
	i2IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.S1, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				s1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.I2, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				i2IsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.B3 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !s1IsSet {
		return errors.New("field S1 of Dgx is required")
	}

	if !i2IsSet {
		return errors.New("field I2 of Dgx is required")
	}

	return nil
}

// Encode serializes a Dgx struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Dgx struct could not be encoded.
func (v *Dgx) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.S1); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TI32}); err != nil {
		return err
	}
	if err := sw.WriteInt32(v.I2); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.B3 != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 3, Type: wire.TBool}); err != nil {
			return err
		}
		if err := sw.WriteBool(*(v.B3)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Dgx struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Dgx struct could not be generated from the wire
// representation.
func (v *Dgx) Decode(sr stream.Reader) error {

	s1IsSet := false
	i2IsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.S1, err = sr.ReadString()
			if err != nil {
				return err
			}
			s1IsSet = true
		case fh.ID == 2 && fh.Type == wire.TI32:
			v.I2, err = sr.ReadInt32()
			if err != nil {
				return err
			}
			i2IsSet = true
		case fh.ID == 3 && fh.Type == wire.TBool:
			var x bool
			x, err = sr.ReadBool()
			v.B3 = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !s1IsSet {
		return errors.New("field S1 of Dgx is required")
	}

	if !i2IsSet {
		return errors.New("field I2 of Dgx is required")
	}

	return nil
}

// String returns a readable string representation of a Dgx
// struct.
func (v *Dgx) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("S1: %v", v.S1)
	i++
	fields[i] = fmt.Sprintf("I2: %v", v.I2)
	i++
	if v.B3 != nil {
		fields[i] = fmt.Sprintf("B3: %v", *(v.B3))
		i++
	}

	return fmt.Sprintf("Dgx{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Dgx match the
// provided Dgx.
//
// This function performs a deep comparison.
func (v *Dgx) Equals(rhs *Dgx) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.S1 == rhs.S1) {
		return false
	}
	if !(v.I2 == rhs.I2) {
		return false
	}
	if !_Bool_EqualsPtr(v.B3, rhs.B3) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Dgx.
func (v *Dgx) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("s1", v.S1)
	enc.AddInt32("i2", v.I2)
	if v.B3 != nil {
		enc.AddBool("b3", *v.B3)
	}
	return err
}

// GetS1 returns the value of S1 if it is set or its
// zero value if it is unset.
func (v *Dgx) GetS1() (o string) {
	if v != nil {
		o = v.S1
	}
	return
}

// GetI2 returns the value of I2 if it is set or its
// zero value if it is unset.
func (v *Dgx) GetI2() (o int32) {
	if v != nil {
		o = v.I2
	}
	return
}

// GetB3 returns the value of B3 if it is set or its
// zero value if it is unset.
func (v *Dgx) GetB3() (o bool) {
	if v != nil && v.B3 != nil {
		return *v.B3
	}

	return
}

// IsSetB3 returns true if B3 is not nil.
func (v *Dgx) IsSetB3() bool {
	return v != nil && v.B3 != nil
}

type Fred struct {
	ContentType string `json:"contentType,required"`
	Auth        string `json:"auth,required"`
}

// ToWire translates a Fred struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Fred) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.ContentType), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueString(v.Auth), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Fred struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Fred struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Fred
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Fred) FromWire(w wire.Value) error {
	var err error

	contentTypeIsSet := false
	authIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ContentType, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				contentTypeIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Auth, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				authIsSet = true
			}
		}
	}

	if !contentTypeIsSet {
		return errors.New("field ContentType of Fred is required")
	}

	if !authIsSet {
		return errors.New("field Auth of Fred is required")
	}

	return nil
}

// Encode serializes a Fred struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Fred struct could not be encoded.
func (v *Fred) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.ContentType); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Auth); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Fred struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Fred struct could not be generated from the wire
// representation.
func (v *Fred) Decode(sr stream.Reader) error {

	contentTypeIsSet := false
	authIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.ContentType, err = sr.ReadString()
			if err != nil {
				return err
			}
			contentTypeIsSet = true
		case fh.ID == 2 && fh.Type == wire.TBinary:
			v.Auth, err = sr.ReadString()
			if err != nil {
				return err
			}
			authIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !contentTypeIsSet {
		return errors.New("field ContentType of Fred is required")
	}

	if !authIsSet {
		return errors.New("field Auth of Fred is required")
	}

	return nil
}

// String returns a readable string representation of a Fred
// struct.
func (v *Fred) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("ContentType: %v", v.ContentType)
	i++
	fields[i] = fmt.Sprintf("Auth: %v", v.Auth)
	i++

	return fmt.Sprintf("Fred{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Fred match the
// provided Fred.
//
// This function performs a deep comparison.
func (v *Fred) Equals(rhs *Fred) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.ContentType == rhs.ContentType) {
		return false
	}
	if !(v.Auth == rhs.Auth) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Fred.
func (v *Fred) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("contentType", v.ContentType)
	enc.AddString("auth", v.Auth)
	return err
}

// GetContentType returns the value of ContentType if it is set or its
// zero value if it is unset.
func (v *Fred) GetContentType() (o string) {
	if v != nil {
		o = v.ContentType
	}
	return
}

// GetAuth returns the value of Auth if it is set or its
// zero value if it is unset.
func (v *Fred) GetAuth() (o string) {
	if v != nil {
		o = v.Auth
	}
	return
}

type Garply struct {
	UUID  *string `json:"UUID,omitempty"`
	Token *string `json:"token,omitempty"`
}

// ToWire translates a Garply struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Garply) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.UUID != nil {
		w, err = wire.NewValueString(*(v.UUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Token != nil {
		w, err = wire.NewValueString(*(v.Token)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Garply struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Garply struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Garply
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Garply) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UUID = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Token = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Garply struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Garply struct could not be encoded.
func (v *Garply) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.UUID != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.UUID)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.Token != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Token)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Garply struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Garply struct could not be generated from the wire
// representation.
func (v *Garply) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.UUID = &x
			if err != nil {
				return err
			}

		case fh.ID == 2 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Token = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Garply
// struct.
func (v *Garply) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.UUID != nil {
		fields[i] = fmt.Sprintf("UUID: %v", *(v.UUID))
		i++
	}
	if v.Token != nil {
		fields[i] = fmt.Sprintf("Token: %v", *(v.Token))
		i++
	}

	return fmt.Sprintf("Garply{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Garply match the
// provided Garply.
//
// This function performs a deep comparison.
func (v *Garply) Equals(rhs *Garply) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.UUID, rhs.UUID) {
		return false
	}
	if !_String_EqualsPtr(v.Token, rhs.Token) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Garply.
func (v *Garply) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.UUID != nil {
		enc.AddString("UUID", *v.UUID)
	}
	if v.Token != nil {
		enc.AddString("token", *v.Token)
	}
	return err
}

// GetUUID returns the value of UUID if it is set or its
// zero value if it is unset.
func (v *Garply) GetUUID() (o string) {
	if v != nil && v.UUID != nil {
		return *v.UUID
	}

	return
}

// IsSetUUID returns true if UUID is not nil.
func (v *Garply) IsSetUUID() bool {
	return v != nil && v.UUID != nil
}

// GetToken returns the value of Token if it is set or its
// zero value if it is unset.
func (v *Garply) GetToken() (o string) {
	if v != nil && v.Token != nil {
		return *v.Token
	}

	return
}

// IsSetToken returns true if Token is not nil.
func (v *Garply) IsSetToken() bool {
	return v != nil && v.Token != nil
}

type Grault struct {
	UUID  *string `json:"UUID,omitempty"`
	Token *string `json:"token,omitempty"`
}

// ToWire translates a Grault struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Grault) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.UUID != nil {
		w, err = wire.NewValueString(*(v.UUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Token != nil {
		w, err = wire.NewValueString(*(v.Token)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Grault struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Grault struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Grault
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Grault) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UUID = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Token = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Grault struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Grault struct could not be encoded.
func (v *Grault) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.UUID != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.UUID)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.Token != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Token)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Grault struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Grault struct could not be generated from the wire
// representation.
func (v *Grault) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.UUID = &x
			if err != nil {
				return err
			}

		case fh.ID == 2 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Token = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Grault
// struct.
func (v *Grault) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.UUID != nil {
		fields[i] = fmt.Sprintf("UUID: %v", *(v.UUID))
		i++
	}
	if v.Token != nil {
		fields[i] = fmt.Sprintf("Token: %v", *(v.Token))
		i++
	}

	return fmt.Sprintf("Grault{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Grault match the
// provided Grault.
//
// This function performs a deep comparison.
func (v *Grault) Equals(rhs *Grault) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.UUID, rhs.UUID) {
		return false
	}
	if !_String_EqualsPtr(v.Token, rhs.Token) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Grault.
func (v *Grault) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.UUID != nil {
		enc.AddString("UUID", *v.UUID)
	}
	if v.Token != nil {
		enc.AddString("token", *v.Token)
	}
	return err
}

// GetUUID returns the value of UUID if it is set or its
// zero value if it is unset.
func (v *Grault) GetUUID() (o string) {
	if v != nil && v.UUID != nil {
		return *v.UUID
	}

	return
}

// IsSetUUID returns true if UUID is not nil.
func (v *Grault) IsSetUUID() bool {
	return v != nil && v.UUID != nil
}

// GetToken returns the value of Token if it is set or its
// zero value if it is unset.
func (v *Grault) GetToken() (o string) {
	if v != nil && v.Token != nil {
		return *v.Token
	}

	return
}

// IsSetToken returns true if Token is not nil.
func (v *Grault) IsSetToken() bool {
	return v != nil && v.Token != nil
}

type Thud struct {
	SomeResHeader *string `json:"someResHeader,omitempty"`
}

// ToWire translates a Thud struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Thud) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.SomeResHeader != nil {
		w, err = wire.NewValueString(*(v.SomeResHeader)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Thud struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Thud struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Thud
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Thud) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.SomeResHeader = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Thud struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Thud struct could not be encoded.
func (v *Thud) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.SomeResHeader != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.SomeResHeader)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Thud struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Thud struct could not be generated from the wire
// representation.
func (v *Thud) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.SomeResHeader = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Thud
// struct.
func (v *Thud) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.SomeResHeader != nil {
		fields[i] = fmt.Sprintf("SomeResHeader: %v", *(v.SomeResHeader))
		i++
	}

	return fmt.Sprintf("Thud{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Thud match the
// provided Thud.
//
// This function performs a deep comparison.
func (v *Thud) Equals(rhs *Thud) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.SomeResHeader, rhs.SomeResHeader) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Thud.
func (v *Thud) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.SomeResHeader != nil {
		enc.AddString("someResHeader", *v.SomeResHeader)
	}
	return err
}

// GetSomeResHeader returns the value of SomeResHeader if it is set or its
// zero value if it is unset.
func (v *Thud) GetSomeResHeader() (o string) {
	if v != nil && v.SomeResHeader != nil {
		return *v.SomeResHeader
	}

	return
}

// IsSetSomeResHeader returns true if SomeResHeader is not nil.
func (v *Thud) IsSetSomeResHeader() bool {
	return v != nil && v.SomeResHeader != nil
}

type TokenOnly struct {
	Token string `json:"Token,required"`
}

// ToWire translates a TokenOnly struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *TokenOnly) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Token), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a TokenOnly struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TokenOnly struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v TokenOnly
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *TokenOnly) FromWire(w wire.Value) error {
	var err error

	TokenIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Token, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				TokenIsSet = true
			}
		}
	}

	if !TokenIsSet {
		return errors.New("field Token of TokenOnly is required")
	}

	return nil
}

// Encode serializes a TokenOnly struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a TokenOnly struct could not be encoded.
func (v *TokenOnly) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Token); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a TokenOnly struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a TokenOnly struct could not be generated from the wire
// representation.
func (v *TokenOnly) Decode(sr stream.Reader) error {

	TokenIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Token, err = sr.ReadString()
			if err != nil {
				return err
			}
			TokenIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !TokenIsSet {
		return errors.New("field Token of TokenOnly is required")
	}

	return nil
}

// String returns a readable string representation of a TokenOnly
// struct.
func (v *TokenOnly) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Token: %v", v.Token)
	i++

	return fmt.Sprintf("TokenOnly{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this TokenOnly match the
// provided TokenOnly.
//
// This function performs a deep comparison.
func (v *TokenOnly) Equals(rhs *TokenOnly) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Token == rhs.Token) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TokenOnly.
func (v *TokenOnly) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("Token", v.Token)
	return err
}

// GetToken returns the value of Token if it is set or its
// zero value if it is unset.
func (v *TokenOnly) GetToken() (o string) {
	if v != nil {
		o = v.Token
	}
	return
}

type UUIDOnly struct {
	UUID string `json:"UUID,required"`
}

// ToWire translates a UUIDOnly struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *UUIDOnly) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.UUID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a UUIDOnly struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a UUIDOnly struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v UUIDOnly
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *UUIDOnly) FromWire(w wire.Value) error {
	var err error

	UUIDIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.UUID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				UUIDIsSet = true
			}
		}
	}

	if !UUIDIsSet {
		return errors.New("field UUID of UUIDOnly is required")
	}

	return nil
}

// Encode serializes a UUIDOnly struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a UUIDOnly struct could not be encoded.
func (v *UUIDOnly) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.UUID); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a UUIDOnly struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a UUIDOnly struct could not be generated from the wire
// representation.
func (v *UUIDOnly) Decode(sr stream.Reader) error {

	UUIDIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.UUID, err = sr.ReadString()
			if err != nil {
				return err
			}
			UUIDIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !UUIDIsSet {
		return errors.New("field UUID of UUIDOnly is required")
	}

	return nil
}

// String returns a readable string representation of a UUIDOnly
// struct.
func (v *UUIDOnly) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("UUID: %v", v.UUID)
	i++

	return fmt.Sprintf("UUIDOnly{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this UUIDOnly match the
// provided UUIDOnly.
//
// This function performs a deep comparison.
func (v *UUIDOnly) Equals(rhs *UUIDOnly) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.UUID == rhs.UUID) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of UUIDOnly.
func (v *UUIDOnly) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("UUID", v.UUID)
	return err
}

// GetUUID returns the value of UUID if it is set or its
// zero value if it is unset.
func (v *UUIDOnly) GetUUID() (o string) {
	if v != nil {
		o = v.UUID
	}
	return
}
