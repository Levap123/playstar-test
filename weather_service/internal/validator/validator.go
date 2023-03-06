package validator

type Validator struct {
	LatitudeMin  float32
	LatitudeMax  float32
	LongitudeMin float32
	LongitudeMax float32
}

func New() Validator {
	return Validator{
		LatitudeMin:  -90.0,
		LatitudeMax:  90.0,
		LongitudeMin: -180.0,
		LongitudeMax: 180.0,
	}
}

func (v *Validator) ValidateCoordinates(latitude, longitude float32) error {
	if v.LatitudeMin > latitude || v.LatitudeMax < latitude {
		return ErrIncorrectLat
	}

	if v.LongitudeMin > longitude || v.LongitudeMax < longitude {
		return ErrIncorrectLon
	}

	return nil
}
