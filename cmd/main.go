package main

import (
	"fmt"
	"github.com/mortdeus/links"
)

var (
	root    = new(link.LinkedObject)
	linkers = []link.Linker{
		new(CustomLink),
		new(link.LinkedObject),
		new(CustomEmbeddedLink),
		&CustomEmbeddedIfaceLink{new(CustomLink)},
		&CustomEmbeddedIfaceLink{new(link.LinkedObject)},
		&CustomEmbeddedIfaceLink{new(CustomEmbeddedLink)},
	}
)

func main() {
	for i := range linkers {
		link.LinkerInfo(linkers[i])
	}
	link.Chain(root, linkers...)
	fmt.Println("[Chainlink Graph]\n")
	for child, i := root.Child(), 0; child != nil; child, i = child.(link.Linker).Child(), i+1 {
		fmt.Printf("(%v) %v\n\t\t|\n\t\tV\n", i, child)
	}

}
