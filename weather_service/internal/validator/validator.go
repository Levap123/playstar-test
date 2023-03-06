package validator

type Validator struct {
	LatitudeMin  int
	LatitudeMax  int
	LongitudeMin int
	LongitudeMax int
}

func New() Validator {
	return Validator{
		LatitudeMin:  -90,
		LatitudeMax:  90,
		LongitudeMin: -180,
		LongitudeMax: 180,
	}
}

func (v *Validator) ValidateCoordinates(latitude, longitude int) error {
	if v.LatitudeMin > latitude || v.LatitudeMax < latitude {
		return ErrIncorrectLat
	}

	if v.LongitudeMin > longitude || v.LongitudeMax < longitude {
		return ErrIncorrectLon
	}

	return nil
}
