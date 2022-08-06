module github.com/gibmir/ion-go/http-client

go 1.18

require github.com/gibmir/ion-go/api v0.0.0

require (
	github.com/gibmir/ion-go/pool v0.0.0
	github.com/sirupsen/logrus v1.8.1
)

require golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect

replace github.com/gibmir/ion-go/client => ../client

replace github.com/gibmir/ion-go/api => ../api

replace github.com/gibmir/ion-go/pool => ../pool
