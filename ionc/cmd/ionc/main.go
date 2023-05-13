package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/gibmir/ion-go/ionc/internal/generator"
	"github.com/gibmir/ion-go/ionc/internal/reader"
	"github.com/sirupsen/logrus"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "./service.json", "path to json schema")
	var out string
	flag.StringVar(&out, "out", "./", "out directory")
	flag.Parse()

	logrus.Infof("starts ionc generation from [%s] to [%s]", path, out)

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatalf("can't read file [%s]. %v ", path, err)
	}
	logrus.Debugf("successfully read file from [%s]", path)

	apiJson := make(map[string]interface{})
	err = json.Unmarshal(fileBytes, &apiJson)
	if err != nil {
		logrus.Fatalf("can't unmarshal schema [%s]. %v", path, err)
	}
	logrus.Debugf("successfully unmarshal schema from [%s]", path)

	schema, err := reader.ReadSchema(path, apiJson)
	if err != nil {
		logrus.Fatalf("can't read schema [%s]. %v", path, err)
	}
	logrus.Debugf("successfully unmarshal schema from [%s]", path)

	err = generator.GenerateTemplate(schema, out)
	if err != nil {
		logrus.Fatalf("can't generate code for schema [%s]. %v", path, err)
	}
	logrus.Debugf("successfully generate code for schema from [%s]", path)

	logrus.Infof("ionc generation complete from [%s] to [%s]", path, out)
}
