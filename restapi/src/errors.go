package src

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExist = errors.New("task already exist")
