// Generated SBE (Simple Binary Encoding) message codec

package model

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Message struct {
	MsgId   uint64
	MsgType uint8
	Content []uint8
}

func (m *Message) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.MsgId); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MsgType); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(m.Content))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Content); err != nil {
		return err
	}
	return nil
}

func (m *Message) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.MsgIdInActingVersion(actingVersion) {
		m.MsgId = m.MsgIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.MsgId); err != nil {
			return err
		}
	}
	if !m.MsgTypeInActingVersion(actingVersion) {
		m.MsgType = m.MsgTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MsgType); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.ContentInActingVersion(actingVersion) {
		var ContentLength uint32
		if err := _m.ReadUint32(_r, &ContentLength); err != nil {
			return err
		}
		if cap(m.Content) < int(ContentLength) {
			m.Content = make([]uint8, ContentLength)
		}
		if err := _m.ReadBytes(_r, m.Content); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *Message) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MsgIdInActingVersion(actingVersion) {
		if m.MsgId < m.MsgIdMinValue() || m.MsgId > m.MsgIdMaxValue() {
			return fmt.Errorf("Range check failed on m.MsgId (%v < %v > %v)", m.MsgIdMinValue(), m.MsgId, m.MsgIdMaxValue())
		}
	}
	if m.MsgTypeInActingVersion(actingVersion) {
		if m.MsgType < m.MsgTypeMinValue() || m.MsgType > m.MsgTypeMaxValue() {
			return fmt.Errorf("Range check failed on m.MsgType (%v < %v > %v)", m.MsgTypeMinValue(), m.MsgType, m.MsgTypeMaxValue())
		}
	}
	return nil
}

func MessageInit(m *Message) {
	return
}

func (*Message) SbeBlockLength() (blockLength uint16) {
	return 9
}

func (*Message) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Message) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (*Message) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Message) SbeSemanticType() (semanticType []byte) {
	return []byte("request")
}

func (*Message) MsgIdId() uint16 {
	return 2
}

func (*Message) MsgIdSinceVersion() uint16 {
	return 0
}

func (m *Message) MsgIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MsgIdSinceVersion()
}

func (*Message) MsgIdDeprecated() uint16 {
	return 0
}

func (*Message) MsgIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message) MsgIdMinValue() uint64 {
	return 0
}

func (*Message) MsgIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Message) MsgIdNullValue() uint64 {
	return math.MaxUint64
}

func (*Message) MsgTypeId() uint16 {
	return 3
}

func (*Message) MsgTypeSinceVersion() uint16 {
	return 0
}

func (m *Message) MsgTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MsgTypeSinceVersion()
}

func (*Message) MsgTypeDeprecated() uint16 {
	return 0
}

func (*Message) MsgTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message) MsgTypeMinValue() uint8 {
	return 0
}

func (*Message) MsgTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Message) MsgTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*Message) ContentMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message) ContentSinceVersion() uint16 {
	return 0
}

func (m *Message) ContentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ContentSinceVersion()
}

func (*Message) ContentDeprecated() uint16 {
	return 0
}

func (Message) ContentCharacterEncoding() string {
	return "null"
}

func (Message) ContentHeaderLength() uint64 {
	return 4
}
