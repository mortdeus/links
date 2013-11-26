package main

import "github.com/mortdeus/links"

var (
	l1 = new(CustomLink)
	l2 = new(link.LinkedObject)
	l3 = new(CustomEmbeddedLink)
)

func main() {
	link.LinkerInfo(l1)
	link.LinkerInfo(l2)
	link.LinkerInfo(l3)

	link.LinkerInfo(&CustomEmbeddedIfaceLink{l1})
	link.LinkerInfo(&CustomEmbeddedIfaceLink{l2})
	link.LinkerInfo(&CustomEmbeddedIfaceLink{l3})
}
