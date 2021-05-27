package timeslots

import (
    "errors"
    "fmt"
)

func cannotFindTimeslotsForDoctorAndServiceError(doctorId int, serviceId int) error {
    return errors.New(fmt.Sprintf("cannot find timeslots for doctor with ID '%d' and service with ID '%d'", doctorId))
}

func cannotReadTimeslotsError() error {
    return errors.New("cannot read timeslots")
}