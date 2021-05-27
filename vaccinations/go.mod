module vaccinations

go 1.16

replace services => ../services

replace timeslots => ../timeslots

replace notifications => ../notifications

replace options => ../options

require (
	github.com/gen2brain/beeep v0.0.0-20200526185328-e9c15c258e28 // indirect
	notifications v0.0.0-00010101000000-000000000000
	options v0.0.0-00010101000000-000000000000
	services v0.0.0-00010101000000-000000000000
	timeslots v0.0.0-00010101000000-000000000000
)
