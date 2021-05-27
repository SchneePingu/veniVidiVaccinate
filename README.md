### What is veniVidiVaccinate?
`veniVidiVaccinate` is a deamon to check the availability of a certain vaccine for a certain doctor.
The purpose is to notify you as soon as a vaccination is available.

### Installation of veniVidiVaccinate

Please download the `veniVidiVaccinate` binary from the latest release.

### Usage of veniVidiVaccinate

`veniVidiVaccinate` runs on the command line and provides options for the vaccines from BioNTech/Pfizer, AstraZeneca and Johnson & Johnson.
Checking the availability of these vaccines requires your doctor's ID. To obtain it, go to [jameda](https://www.jameda.de/) and search for the doctor, who provides you with vaccination.
The ID is then displayed in the URL, in the format `https://www.jameda.de/.../ID_1/`.

If the requested vaccine is available for your doctor, `veniVidiVaccinate` - as long as it is executed - notifies you with a recurring pop-up notification,
such that you can book a vaccination on [jameda](https://www.jameda.de/), manually.

```
veniVidiVaccinate -help

Usage of ./veniVidiVaccinate:
  -astrazeneca
    	enable search for AstraZeneca vaccine (default false)
  -biontech
    	enable search for BioNTech/Pfizer vaccine (default false)
  -doctor int
    	ID of doctor on jameda (required option)
  -interval int
    	interval between two successive updates given in minutes (default 30)
  -johnson
    	enable search for Johnson & Johnson vaccine (default false)
```