package main

import "fmt"

type Namespace struct {
	name string
}

//todo use network connection instead of exec
func (ns Namespace) create() {
	execute(fmt.Sprintf("oc create ns %s", ns.name))
}

func (ns Namespace) remove() {
	execute(fmt.Sprintf("oc delete ns %s", ns.name))
}
