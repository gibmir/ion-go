# ion-go

Project represents json-rpc 2.0 protocol implementation. It is a part of ion project

## Features:

* Functional remote procedure call API
* API generation with ionc. [Schema example](https://github.com/gibmir/ion-go/tree/master/demo/api/service.json)
* Generic http client and server

## Example

 * install [ionc](https://github.com/gibmir/ion-go/tree/master/ionc)
 * install Make
 * update [schema](https://github.com/gibmir/ion-go/tree/master/demo/api/service.json)
 * run from [demo root](https://github.com/gibmir/ion-go/tree/master/demo) 
```sh
make generate
```
 * update [client](https://github.com/gibmir/ion-go/tree/master/demo/cmd/client/client.go)
 * update [server](https://github.com/gibmir/ion-go/tree/master/demo/cmd/server/server.go)
 * run server
 * run client
