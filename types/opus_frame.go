package types

import (
	"time"

	"github.com/norasector/turbine-common/types/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaggedAudioFrameOpus struct {
	TalkGroup                *TalkGroup
	Audio                    *SegmentBinaryBytes
	SampleLengthMicroseconds int
	Timestamp                time.Time
}

func (t *TaggedAudioFrameOpus) ToProtobuf() *pb.TaggedOpusFrame {

	return &pb.TaggedOpusFrame{
		Tgid:                     uint32(t.TalkGroup.ID),
		FrameNumber:              uint64(t.Audio.SegmentNumber),
		SampleLengthMicroseconds: uint32(t.SampleLengthMicroseconds),
		Data:                     t.Audio.Data,
		Ts:                       timestamppb.New(t.Timestamp),
		SystemId:                 uint32(t.TalkGroup.SystemID),
		SrcId:                    uint32(t.TalkGroup.SourceID),
	}
}

func OpusFrameFromProtobuf(a *pb.TaggedOpusFrame) *TaggedAudioFrameOpus {
	return &TaggedAudioFrameOpus{
		TalkGroup: &TalkGroup{
			SystemID: int(a.SystemId),
			ID:       int(a.Tgid),
			SourceID: int(a.SrcId),
		},
		Audio: &SegmentBinaryBytes{
			SegmentNumber: int(a.FrameNumber),
			Data:          a.Data,
		},
		SampleLengthMicroseconds: int(a.SampleLengthMicroseconds),
		Timestamp:                a.Ts.AsTime(),
	}
}
