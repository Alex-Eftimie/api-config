package apiconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// ConfigurationInterface .
type ConfigurationInterface interface {

	// ConfigurationInterface needs to implement Locker
	sync.Locker

	// ConfigurationInterface needs to have a root AuthToken
	AuthToken() string
}

// Configuration is the base Configuration object
type Configuration struct {
	*sync.Mutex
	Token string `json:"AuthToken"`
}

// NewConfig returns a pointer to a filled new instance of Configuration
func NewConfig() *Configuration {
	return &Configuration{
		Mutex: &sync.Mutex{},
	}
}

// AuthToken return the root authToken
func (c *Configuration) AuthToken() string {

	return c.Token
}

// ConfigFile is the file where the settings are stored
var ConfigFile = "./config.json"

// LoadConfig loads the config and return the fiilled object
func LoadConfig(Config ConfigurationInterface) ConfigurationInterface {

	jsonFile, err := os.Open(ConfigFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Fatalln(err)
	}
	return Config
}

// Sync Writes the config to disk
// Presumably after you've changed it but it does not do any checks
func (c *Configuration) Sync() {
	c.Lock()
	b, err := json.MarshalIndent(c, "", "\t")
	c.Unlock()
	if err != nil {
		log.Panicf("Json Marshal Error: %s", err)
	}

	err = ioutil.WriteFile(ConfigFile, b, 0644)

	if err != nil {
		log.Panicf("Failed to write config.json")
	}
}
