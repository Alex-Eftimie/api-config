package apiconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Configuration
	Custom       string
	MyCustomBool bool
	Subkey       struct {
		Opt int
	}
}

var TC *TestConfig

func TestLoadConfig(t *testing.T) {

	TC = &TestConfig{
		Configuration: *NewConfig(""),
	}

	LoadConfig(TC)

	assert.Equal(t, "testtoken", TC.AuthToken(), "The token should be: testtoken")
}

func TestSyncConfig(t *testing.T) {

	TC.Token = "TestSync"

	TC.Sync()

	assert.Equal(t, "TestSync", TC.AuthToken(), "The token should be: TestSync")

	TC.Token = "testtoken"

	TC.Sync()

	assert.Equal(t, "testtoken", TC.AuthToken(), "The token should be: testtoken")

}

func TestCustomOption(t *testing.T) {
	assert.Equal(t, "My Option", TC.Custom, "The Custom json key does not match")
}

func TestSubkeyOption(t *testing.T) {
	assert.Equal(t, 123, TC.Subkey.Opt, "Config.Subkey.Opt does not match")
}
