package config

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_InitConfigPanics(t *testing.T) {
  assert.Panics(t, InitConfig)
}
