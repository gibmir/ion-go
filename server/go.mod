module github.com/gibmir/ion-go/server

go 1.18

require (
	github.com/gibmir/ion-go/api v0.0.0
	github.com/gibmir/ion-go/pool v0.0.0
	github.com/gibmir/ion-go/processor v0.0.0
	github.com/sirupsen/logrus v1.9.0
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace github.com/gibmir/ion-go/api => ../api

replace github.com/gibmir/ion-go/pool => ../pool

replace github.com/gibmir/ion-go/processor => ../processor
