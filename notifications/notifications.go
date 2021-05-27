package notifications

import (
    "time"
    "github.com/gen2brain/beeep"
)

func SendRepeatingPopUpNotificationForAvailableService(service string) {
    for true {
        beeep.Alert(service, popUpNotificationMessage, popUpNotificationType)
        time.Sleep(popUpNotificationIntervalInSeconds * time.Second)
    }
}