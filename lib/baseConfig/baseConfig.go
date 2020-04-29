package baseconfig

import (
  "reflect"
  "strconv"
  "strings"

  "github.com/spf13/viper"
)

func setFieldValue(field reflect.Value, nameTag string) {
  switch field.Kind() {

  case reflect.Int32:
    field.SetInt(int64(viper.GetInt32(nameTag)))

  case reflect.String:
    field.SetString(viper.GetString(nameTag))

  case reflect.Bool:
    field.SetBool(viper.GetBool(nameTag))

  default:
    panic("Unsupported value config type, type = " + field.Kind().String())
  }
}

func ReadConfigValues(keys interface{}) {
  var notSetVariables []string

  t := reflect.TypeOf(keys).Elem()
  if t == nil {
    panic("Unable to set Env variables")
  }

  s := reflect.ValueOf(keys).Elem()

  for i := 0; i < t.NumField(); i++ {
    fieldType := t.Field(i)

    nameTag := fieldType.Tag.Get("name")
    if nameTag == "" {
      panic("`name` tag is not set for `settingsKeys` field `" + fieldType.Name + "`")
    }

    _ = viper.BindEnv(nameTag)

    field := s.Field(i)
    if !field.IsValid() && !field.CanSet() {
      panic("Unable to set config struct keys")
    }

    if defaultTag := fieldType.Tag.Get("default"); defaultTag != "" {
      viper.SetDefault(nameTag, defaultTag)
    }

    optionalTag := fieldType.Tag.Get("optional")
    isOptional, _ := strconv.ParseBool(optionalTag)

    if !viper.IsSet(nameTag) && !isOptional {
      notSetVariables = append(notSetVariables, "`"+nameTag+"`")
      continue
    }

    setFieldValue(field, nameTag)
  }

  if len(notSetVariables) > 0 {
    panic("Variables are not set: " + strings.Join(notSetVariables, ", "))
  }
}
