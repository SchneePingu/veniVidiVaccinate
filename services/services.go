package services

import (
    "fmt"
    "encoding/json"
    "net/http"
)

func GetIdForService(doctorId int, title string) (int, error) {
    var serviceId int

    response, err := requestServices(doctorId)
    if err != nil {
        return serviceId, err
    }

   services, err := parseResponseForServices(response)
   if err != nil {
        return serviceId, err
   }

   service, err := getServiceByTitle(services, title)
   if err != nil {
        return serviceId, err
   }

   return service.Id, nil
}

func requestServices(doctorId int) (*http.Response, error) {
    urlForRequestingServices := createUrlForRequestingServices(doctorId)

    client := http.Client{}
    response, err := client.Get(urlForRequestingServices)
    if err != nil {
        return response, cannotFindServicesForDoctorError(doctorId)
    }

   return response, nil
}

func createUrlForRequestingServices(doctorId int) string {
    return fmt.Sprintf(urlTemplateForRequestingServices, doctorId)
}

func parseResponseForServices(response *http.Response) ([]service, error) {
    var services []service
    defer response.Body.Close()

    err := json.NewDecoder(response.Body).Decode(&services)
    if err != nil {
        return services, cannotReadServicesError()
    }

    return services, nil
}

func getServiceByTitle(services []service, title string) (service, error) {
    for _, service := range services {
        if service.Title == title {
            return service, nil
        }
    }

    return service{}, cannotFindServiceWithTitleError(title)
}