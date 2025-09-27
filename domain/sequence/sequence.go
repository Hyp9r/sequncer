package sequence

type Sequence struct {
	Name                 string
	OpenTrackingEnabled  bool
	ClickTrackingEnabled bool
}

func NewSequence(name string, openTrackingEnabled bool, clickTrackingEnabled bool) *Sequence {
	return &Sequence{
		Name:                 name,
		OpenTrackingEnabled:  openTrackingEnabled,
		ClickTrackingEnabled: clickTrackingEnabled,
	}
}
