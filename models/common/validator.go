package common

import valid "github.com/iteny/hmgo/govalidator"

func init() {
	valid.SetFieldsRequiredByDefault(true)
}
