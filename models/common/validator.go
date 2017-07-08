package common

import valid "github.com/asaskevich/govalidator"

func init() {
	valid.SetFieldsRequiredByDefault(true)
}
