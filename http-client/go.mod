module github.com/gibmir/ion-go/http-client

go 1.18

require (
	github.com/gibmir/ion-go/api v0.0.0
	github.com/gibmir/ion-go/pool v0.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gibmir/ion-go/client => ../client

replace github.com/gibmir/ion-go/api => ../api

replace github.com/gibmir/ion-go/pool => ../pool
