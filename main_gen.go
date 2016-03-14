package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Hotel) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "HotelCode":
			z.HotelCode, err = dc.ReadString()
			if err != nil {
				return
			}
		case "HotelStatus":
			z.HotelStatus, err = dc.ReadString()
			if err != nil {
				return
			}
		case "HotelCurrency":
			z.HotelCurrency, err = dc.ReadString()
			if err != nil {
				return
			}
		case "MinutesOffsetUtc":
			z.MinutesOffsetUtc, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "ProviderName":
			z.ProviderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Rooms":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(xsz) {
				z.Rooms = z.Rooms[:xsz]
			} else {
				z.Rooms = make([]Room, xsz)
			}
			for xvk := range z.Rooms {
				err = z.Rooms[xvk].DecodeMsg(dc)
				if err != nil {
					return
				}
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
func (z *Hotel) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Id"
	err = en.Append(0x87, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Id)
	if err != nil {
		return
	}
	// write "HotelCode"
	err = en.Append(0xa9, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HotelCode)
	if err != nil {
		return
	}
	// write "HotelStatus"
	err = en.Append(0xab, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HotelStatus)
	if err != nil {
		return
	}
	// write "HotelCurrency"
	err = en.Append(0xad, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HotelCurrency)
	if err != nil {
		return
	}
	// write "MinutesOffsetUtc"
	err = en.Append(0xb0, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x55, 0x74, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.MinutesOffsetUtc)
	if err != nil {
		return
	}
	// write "ProviderName"
	err = en.Append(0xac, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ProviderName)
	if err != nil {
		return
	}
	// write "Rooms"
	err = en.Append(0xa5, 0x52, 0x6f, 0x6f, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Rooms)))
	if err != nil {
		return
	}
	for xvk := range z.Rooms {
		err = z.Rooms[xvk].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Hotel) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Id"
	o = append(o, 0x87, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt(o, z.Id)
	// string "HotelCode"
	o = append(o, 0xa9, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.HotelCode)
	// string "HotelStatus"
	o = append(o, 0xab, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.HotelStatus)
	// string "HotelCurrency"
	o = append(o, 0xad, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79)
	o = msgp.AppendString(o, z.HotelCurrency)
	// string "MinutesOffsetUtc"
	o = append(o, 0xb0, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x55, 0x74, 0x63)
	o = msgp.AppendInt(o, z.MinutesOffsetUtc)
	// string "ProviderName"
	o = append(o, 0xac, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ProviderName)
	// string "Rooms"
	o = append(o, 0xa5, 0x52, 0x6f, 0x6f, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Rooms)))
	for xvk := range z.Rooms {
		o, err = z.Rooms[xvk].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Hotel) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "HotelCode":
			z.HotelCode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HotelStatus":
			z.HotelStatus, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HotelCurrency":
			z.HotelCurrency, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "MinutesOffsetUtc":
			z.MinutesOffsetUtc, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ProviderName":
			z.ProviderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rooms":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Rooms) >= int(xsz) {
				z.Rooms = z.Rooms[:xsz]
			} else {
				z.Rooms = make([]Room, xsz)
			}
			for xvk := range z.Rooms {
				bts, err = z.Rooms[xvk].UnmarshalMsg(bts)
				if err != nil {
					return
				}
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

func (z *Hotel) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 10 + msgp.StringPrefixSize + len(z.HotelCode) + 12 + msgp.StringPrefixSize + len(z.HotelStatus) + 14 + msgp.StringPrefixSize + len(z.HotelCurrency) + 17 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.ProviderName) + 6 + msgp.ArrayHeaderSize
	for xvk := range z.Rooms {
		s += z.Rooms[xvk].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *HotelDetailsResponseModel) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hotels":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Hotels) >= int(xsz) {
				z.Hotels = z.Hotels[:xsz]
			} else {
				z.Hotels = make([]Hotel, xsz)
			}
			for bzg := range z.Hotels {
				err = z.Hotels[bzg].DecodeMsg(dc)
				if err != nil {
					return
				}
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
func (z *HotelDetailsResponseModel) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Hotels"
	err = en.Append(0x81, 0xa6, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Hotels)))
	if err != nil {
		return
	}
	for bzg := range z.Hotels {
		err = z.Hotels[bzg].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *HotelDetailsResponseModel) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Hotels"
	o = append(o, 0x81, 0xa6, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Hotels)))
	for bzg := range z.Hotels {
		o, err = z.Hotels[bzg].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *HotelDetailsResponseModel) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hotels":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Hotels) >= int(xsz) {
				z.Hotels = z.Hotels[:xsz]
			} else {
				z.Hotels = make([]Hotel, xsz)
			}
			for bzg := range z.Hotels {
				bts, err = z.Hotels[bzg].UnmarshalMsg(bts)
				if err != nil {
					return
				}
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

func (z *HotelDetailsResponseModel) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for bzg := range z.Hotels {
		s += z.Hotels[bzg].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Room) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "InvCode":
			z.InvCode, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RackRate":
			z.RackRate, err = dc.ReadString()
			if err != nil {
				return
			}
		case "HideRackRate":
			z.HideRackRate, err = dc.ReadBool()
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
func (z *Room) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Id"
	err = en.Append(0x84, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Id)
	if err != nil {
		return
	}
	// write "InvCode"
	err = en.Append(0xa7, 0x49, 0x6e, 0x76, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.InvCode)
	if err != nil {
		return
	}
	// write "RackRate"
	err = en.Append(0xa8, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RackRate)
	if err != nil {
		return
	}
	// write "HideRackRate"
	err = en.Append(0xac, 0x48, 0x69, 0x64, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.HideRackRate)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Room) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Id"
	o = append(o, 0x84, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt(o, z.Id)
	// string "InvCode"
	o = append(o, 0xa7, 0x49, 0x6e, 0x76, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.InvCode)
	// string "RackRate"
	o = append(o, 0xa8, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.RackRate)
	// string "HideRackRate"
	o = append(o, 0xac, 0x48, 0x69, 0x64, 0x65, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.HideRackRate)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Room) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "InvCode":
			z.InvCode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RackRate":
			z.RackRate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HideRackRate":
			z.HideRackRate, bts, err = msgp.ReadBoolBytes(bts)
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

func (z *Room) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.InvCode) + 9 + msgp.StringPrefixSize + len(z.RackRate) + 13 + msgp.BoolSize
	return
}
