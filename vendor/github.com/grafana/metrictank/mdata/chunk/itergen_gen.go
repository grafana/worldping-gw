package chunk

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *IterGen) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "T0":
			z.T0, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "IntervalHint":
			z.IntervalHint, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "B":
			z.B, err = dc.ReadBytes(z.B)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *IterGen) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "T0"
	err = en.Append(0x83, 0xa2, 0x54, 0x30)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.T0)
	if err != nil {
		return
	}
	// write "IntervalHint"
	err = en.Append(0xac, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.IntervalHint)
	if err != nil {
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.B)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IterGen) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "T0"
	o = append(o, 0x83, 0xa2, 0x54, 0x30)
	o = msgp.AppendUint32(o, z.T0)
	// string "IntervalHint"
	o = append(o, 0xac, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x69, 0x6e, 0x74)
	o = msgp.AppendUint32(o, z.IntervalHint)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendBytes(o, z.B)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IterGen) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "T0":
			z.T0, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "IntervalHint":
			z.IntervalHint, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "B":
			z.B, bts, err = msgp.ReadBytesBytes(bts, z.B)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *IterGen) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint32Size + 13 + msgp.Uint32Size + 2 + msgp.BytesPrefixSize + len(z.B)
	return
}
