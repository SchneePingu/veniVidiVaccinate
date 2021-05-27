package options

import (
    "errors"
    "fmt"
)

func missingRequiredOptionError(optionName string) error {
    return errors.New(fmt.Sprintf("required option '%s' is missing", optionName))
}