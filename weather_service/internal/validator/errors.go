package validator

import "errors"

var ErrIncorrectLat = errors.New("expected values are in [-90, 90] range")
var ErrIncorrectLon = errors.New("expected values are in [-180, 180] range")