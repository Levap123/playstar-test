package validator

type Validator struct {
	latitudeMin  float32
	latitudeMax  float32
	longitudeMin float32
	longitudeMax float32
}

func New() Validator {
	return Validator{
		latitudeMin:  -90.0,
		latitudeMax:  90.0,
		longitudeMin: -180.0,
		longitudeMax: 180.0,
	}
}

func (v *Validator) ValidateCoordinates(latitude, longitude float32) error {
	if v.latitudeMin > latitude || v.latitudeMax < latitude {
		return ErrIncorrectLat
	}

	if v.longitudeMin > longitude || v.longitudeMax < longitude {
		return ErrIncorrectLon
	}

	return nil
}
