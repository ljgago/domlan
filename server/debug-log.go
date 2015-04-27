package server

import (
  "log"
)

var DEBUG_ENABLE bool = true

func DEBUG(t ...interface{}) {
  if DEBUG_ENABLE {
    log.Println(t...)
  } else {
    return
  }
}