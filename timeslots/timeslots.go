package timeslots

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func IsTimeslotForDoctorAndServiceAvailable(doctorId, serviceId int) (bool, error) {
    timeslots, err := getTimeslotsForDoctorAndService(doctorId, serviceId)
    if err != nil {
        return false, nil
    }

    isTimeslotAvailable := !isListOfTimeslotsEmpty(timeslots)

    return isTimeslotAvailable, nil
}

func getTimeslotsForDoctorAndService(doctorId, serviceId int) ([]timeslot, error) {
    response, err := requestTimeslotsForService(doctorId, serviceId)
    if err != nil {
        return nil, err
    }

    return parseResponseForTimeslots(response)
}

func requestTimeslotsForService(doctorId, serviceId int) (*http.Response, error){
    urlForRequestingTimeslots := createUrlForRequestingTimeslots(doctorId, serviceId)

    client := http.Client{}
    response, err := client.Get(urlForRequestingTimeslots)
    if err != nil {
        return response, cannotFindTimeslotsForDoctorAndServiceError(doctorId, serviceId)
    }

    return response, nil
}

func createUrlForRequestingTimeslots(doctorId, serviceId int) string {
    return fmt.Sprintf(urlTemplateForRequestingTimeslots, doctorId, serviceId)
}

func parseResponseForTimeslots(response *http.Response) ([]timeslot, error) {
    var timeslots []timeslot
    defer response.Body.Close()

    err := json.NewDecoder(response.Body).Decode(&timeslots)
    if err != nil {
        return timeslots, cannotReadTimeslotsError()
    }

    return timeslots, nil
}

func isListOfTimeslotsEmpty(timeslots []timeslot) bool {
    return len(timeslots) == 0
}