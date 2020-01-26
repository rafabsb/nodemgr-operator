package controller

import (
	"github.com/rafabsb/nodemgr-operator/pkg/controller/nodemgr"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodemgr.Add)
}
