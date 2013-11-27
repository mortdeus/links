package main

import (
	"fmt"
	"github.com/mortdeus/links"
	"reflect"
)

type CustomEmbeddedIfaceLink struct {
	link.Linker
}

func (l *CustomEmbeddedIfaceLink) Link(i interface{}) error { return l.Linker.Link(i) }
func (l *CustomEmbeddedIfaceLink) Parent() interface{}      { return l.Linker.Parent() }
func (l *CustomEmbeddedIfaceLink) Child() interface{}       { return l.Linker.Child() }
func (l *CustomEmbeddedIfaceLink) String() string {
	return fmt.Sprint(
		reflect.TypeOf(l).String(), "\n",
		"EmbeddedType:\t", reflect.TypeOf(l.Linker).String())
}

type CustomEmbeddedLink struct {
	CustomLink
}

func (l *CustomEmbeddedLink) String() string {
	return fmt.Sprint(
		reflect.TypeOf(l).String(), "\n",
		"EmbeddedType:\t", l.CustomLink.String())
}

type CustomLink struct {
	parent, child interface{}
}

func (l *CustomLink) Link(i interface{}) error { l.child = i; return nil }
func (l *CustomLink) Parent() interface{}      { return l.parent }
func (l *CustomLink) Child() interface{}       { return l.child }
func (l *CustomLink) String() string           { return fmt.Sprint(reflect.TypeOf(l).String()) }
