package helpers

import (
    "encoding/json"
    "io/ioutil"
)

type Configuration struct {
    Database `json:"database"`
}

type Database struct {
    ConnectionString string `json:"connection_string"`
    Type             string `json:"type"`
}

func LoadConfiguration() Configuration {
    data, _ := ioutil.ReadFile("conf/mapping.json")
    var configuration Configuration
    json.Unmarshal(data, &configuration)
    return configuration
}
