package dsl

import "minecraftremote/src/httprouteradapter"

type dsl interface {
}

func GivenALinuxRemote() httprouteradapter.HTTPRouterAdapter {
	return httprouteradapter.HTTPRouterAdapter{}
}
