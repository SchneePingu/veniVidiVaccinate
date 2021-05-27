package vaccinations

import (
    "fmt"
    "log"
    "services"
    "timeslots"
    "notifications"
    "options"
)

func CheckIfVaccinationsAreAvailable() {
    if *options.IsBioNTechPfizerVaccination {
        go checkIfServiceIsAvailableForDoctor(*options.DoctorId, bioNTechPfizerVaccination)
    }

    if *options.IsAstraZenecaVaccination {
        go checkIfServiceIsAvailableForDoctor(*options.DoctorId, astraZenecaVaccination)
    }

    if *options.IsJohnsonAndJohnsonVaccination {
        go checkIfServiceIsAvailableForDoctor(*options.DoctorId, johnsonAndJohnsonVaccination)
    }
}

func checkIfServiceIsAvailableForDoctor(doctorId int, service string) {
    serviceId, err := services.GetIdForService(doctorId, service)
    if err != nil {
        logError(doctorId, service, err)
        return
    }

    timeslotAvailable, err := timeslots.IsTimeslotForDoctorAndServiceAvailable(doctorId, serviceId)
    if err != nil {
        logError(doctorId, service, err)
        return
    }

    if timeslotAvailable {
        notifications.SendRepeatingPopUpNotificationForAvailableService(service)
    }
}

func logError(doctorId int, service string, err error) {
    log.Printf(fmt.Sprintf(errorMessageTemplate, doctorId, service))
    log.Printf(err.Error())
}
