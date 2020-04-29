package baseconfig

import (
  "os"
  "testing"

  "github.com/stretchr/testify/assert"
)

type settingKeys struct {
  IntProperty    int32  `name:"TEST_INT_PROPERTY"`
  StringProperty string `name:"TEST_STRING_PROPERTY"`
}

type brokenSettingsKeys struct {
  IntProperty int32 `someOtherTag:"BLA-BLA"`
}

type optionalSettingKeys struct {
  IntProperty    int32  `name:"TEST_INT_PROPERTY" optional:"true"`
  StringProperty string `name:"TEST_STRING_PROPERTY"`
}

const (
  intPropertyKey         = "TEST_INT_PROPERTY"
  stringPropertyKey      = "TEST_STRING_PROPERTY"
  expectedIntEnvValue    = "25"
  expectedStringEnvValue = "ENV_STRING"
)

func Test_ShouldProperlyReadConfigValues(t *testing.T) {
  keys := settingKeys{}
  expectedKeys := settingKeys{
    IntProperty:    25,
    StringProperty: "ENV_STRING",
  }

  if err := os.Setenv(intPropertyKey, expectedIntEnvValue); err != nil {
    t.Errorf("Can't setup %v variable", intPropertyKey)
  }
  defer os.Unsetenv(intPropertyKey)

  if err := os.Setenv(stringPropertyKey, expectedStringEnvValue); err != nil {
    t.Errorf("Can't setup %v variable", stringPropertyKey)
  }
  defer os.Unsetenv(stringPropertyKey)

  ReadConfigValues(&keys)

  assert.Equal(t, expectedKeys, keys)
}

func Test_ShouldPanicIfEnvVariableIsNotSet(t *testing.T) {
  keys := settingKeys{}

  if err := os.Setenv(intPropertyKey, expectedIntEnvValue); err != nil {
    t.Errorf("Can't setup %v variable", intPropertyKey)
  }
  defer os.Unsetenv(intPropertyKey)

  assert.PanicsWithValue(
    t,
    "Variables are not set: `TEST_STRING_PROPERTY`",
    func() {
      ReadConfigValues(&keys)
    },
  )
}

func Test_ShouldPanicIfNameTagIsNotSpecified(t *testing.T) {
  keys := brokenSettingsKeys{}

  assert.PanicsWithValue(
    t,
    "`name` tag is not set for `settingsKeys` field `IntProperty`",
    func() {
      ReadConfigValues(&keys)
    },
  )
}

func Test_ShouldSkipIfOptional(t *testing.T) {
  keys := optionalSettingKeys{}
  expectedKeys := optionalSettingKeys{
    StringProperty: "ENV_STRING",
  }

  if err := os.Setenv(stringPropertyKey, expectedStringEnvValue); err != nil {
    t.Errorf("Can't setup %v variable", stringPropertyKey)
  }
  defer os.Unsetenv(stringPropertyKey)

  ReadConfigValues(&keys)

  assert.Equal(t, expectedKeys, keys)
}
