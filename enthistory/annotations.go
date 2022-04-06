package enthistory

type Annotation struct{}

const annotationName = "EntHistory"

func (Annotation) Name() string {
	return annotationName
}

// TrackField marks the field to be tracked by enthistory
func TrackField() Annotation {
	return Annotation{}
}
