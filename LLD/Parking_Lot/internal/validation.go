package internal

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

const phoneNumberRegex = `^(\+91[\-\s]?)?[0]?(91)?[789]\d{9}$`
const vehicleNumberRegex = `^[A-Z]{2}[0-9]{2}[A-Z]{2}[0-9]{4}$`

func (v *Vehicle) Validate() error {
	err := validation.ValidateStruct(v,
		validation.Field(&v.VehicleType, validation.Required),
		validation.Field(&v.VehicleSubType, validation.Required),
		validation.Field(&v.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(phoneNumberRegex))),
		validation.Field(&v.VehicleNumber, validation.Required, validation.Match(regexp.MustCompile(vehicleNumberRegex))),
	)

	if err == nil {
		return nil
	}

	return err
}
