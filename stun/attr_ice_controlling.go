package stun

import (
	"github.com/pkg/errors"
)

type IceControlling struct {
	TieBreaker uint64
}

func (i *IceControlling) Pack(message *Message) error {
	v := make([]byte, 8)
	enc.PutUint64(v, i.TieBreaker)
	message.AddAttribute(AttrIceControlling, v)
	return nil
}

func (i *IceControlling) Unpack(message *Message, rawAttribute *RawAttribute) error {
	v := rawAttribute.Value

	if len(v) != 8 {
		return errors.Errorf("invalid TieBreaker length %d != %d (expected)", len(v), 8)
	}

	i.TieBreaker = enc.Uint64(v)

	return nil
}
