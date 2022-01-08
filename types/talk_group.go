package types

import "time"

type TalkGroup struct {
	SystemID   int
	ID         int
	SourceID   int
	Frequency  int
	LastUpdate time.Time
}

type TaggedAudioSampleFloat32 struct {
	TalkGroup *TalkGroup
	Audio     *SegmentFloat32
}
