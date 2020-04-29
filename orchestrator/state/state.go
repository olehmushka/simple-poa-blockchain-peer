package state

type State struct {}

var instance *State

func GetState() *State {
  if instance == nil {
    instance = &State{}
  }
  return instance
}
