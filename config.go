package apiconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/muhammadmuzzammil1998/jsonc"
)

// ConfigFile is the file where the settings are stored
var ConfigFile = "./config.json"

// ConfigurationInterface .
type ConfigurationInterface interface {

	// ConfigurationInterface needs to implement Locker
	sync.Locker

	// ConfigurationInterface needs to have a root AuthToken
	AuthToken() string
	setObj(interface{})
}

// Configuration is the base Configuration object
type Configuration struct {
	*sync.Mutex `json:"-"`
	Token       string `json:"AuthToken"`
	actualObj   interface{}
	configFile  string
}

// NewConfig returns a pointer to a filled new instance of Configuration
func NewConfig(file string) *Configuration {
	if len(file) == 0 {
		file = ConfigFile
	}
	return &Configuration{
		Mutex:      &sync.Mutex{},
		configFile: file,
	}
}

// AuthToken return the root authToken
func (c *Configuration) AuthToken() string {

	return c.Token
}
func (c *Configuration) setObj(obj interface{}) {
	c.actualObj = obj
}

// LoadConfig loads the config and return the fiilled object
func (c *Configuration) LoadConfig(Config ConfigurationInterface) ConfigurationInterface {

	jsonFile, err := os.Open(c.configFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	byteValue = jsonc.ToJSON(byteValue)

	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Fatalln(err)
	}
	Config.setObj(Config)
	return Config
}

// Sync Writes the config to disk
// Presumably after you've changed it but it does not do any checks
func (c *Configuration) Sync() {
	c.Lock()
	b, err := json.MarshalIndent(c.actualObj, "", "\t")
	c.Unlock()
	if err != nil {
		log.Panicf("Json Marshal Error: %s", err)
	}

	err = ioutil.WriteFile(c.configFile, b, 0644)

	if err != nil {
		log.Panicf("Failed to write config.json")
	}
}
