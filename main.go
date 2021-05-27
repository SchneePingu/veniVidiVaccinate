package main

import (
    "time"
    "options"
    "vaccinations"
)

func main() {
    options.ParseCommandLineOptions()

    for true {
        go vaccinations.CheckIfVaccinationsAreAvailable()
        waitForNMinutes(*options.UpdateInterval)
    }
}

func waitForNMinutes(minutes int) {
    time.Sleep(time.Duration(minutes) * time.Minute)
}