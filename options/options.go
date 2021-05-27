package options

import (
    "flag"
    "log"
)

var (
    DoctorId *int
    IsBioNTechPfizerVaccination *bool
    IsAstraZenecaVaccination *bool
    IsJohnsonAndJohnsonVaccination *bool
    UpdateInterval *int
)

func ParseCommandLineOptions() {
    DoctorId = flag.Int(nameOfDoctorOption, 0, descriptionOfDoctorOption)
    IsBioNTechPfizerVaccination = flag.Bool(nameOfBioNTechPfizerOption, false, descriptionOfBioNTechPfizerOption)
    IsAstraZenecaVaccination = flag.Bool(nameOfAstraZenecaOption, false, descriptionOfAstraZenecaOption)
    IsJohnsonAndJohnsonVaccination = flag.Bool(nameOfJohnsonAndJohnsonOption, false, descriptionOfJohnsonAndJohnsonOption)
    UpdateInterval = flag.Int(nameOfUpdateIntervalOption, 30, descriptionOfUpdateIntervalOption)

    flag.Parse()

    err:= checkIfRequiredOptionsAreGiven()
    if err != nil {
        log.Fatalf(err.Error())
    }
}

func checkIfRequiredOptionsAreGiven() error {
    requiredOptions := []string{nameOfDoctorOption}

    givenOptions := make(map[string]bool)
    flag.Visit(func(f *flag.Flag) { givenOptions[f.Name] = true })
    for _, optionName := range requiredOptions {
        if !givenOptions[optionName] {
            return missingRequiredOptionError(optionName)
        }
    }

    return nil
}