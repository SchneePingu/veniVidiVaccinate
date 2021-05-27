module venividivaccinate

go 1.16

require (
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20210519211817-2312de329ae4 // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	options v0.0.0-00010101000000-000000000000
	vaccinations v0.0.0-00010101000000-000000000000
)

replace services => ./services

replace timeslots => ./timeslots

replace options => ./options

replace notifications => ./notifications

replace vaccinations => ./vaccinations
