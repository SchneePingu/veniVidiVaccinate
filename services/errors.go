package services

import (
    "fmt"
    "errors"
)

func cannotFindServicesForDoctorError(doctorId int) error {
    return errors.New(fmt.Sprintf("cannot find services for doctor with ID '%d'", doctorId))
}

func cannotReadServicesError() error {
    return errors.New("cannot read services")
}

func cannotFindServiceWithTitleError(title string) error {
    return errors.New(fmt.Sprintf("cannot find service with title '%s'", title))
}