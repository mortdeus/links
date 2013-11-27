package main

import (
	"fmt"
	"github.com/mortdeus/links"
	"reflect"
)

type CustomEmbeddedIfaceLink struct {
	link.Linker
}

func (l *CustomEmbeddedIfaceLink) String() string {
	return fmt.Sprintf("%v {%v}", reflect.TypeOf(l).String(), reflect.TypeOf(l.Linker).String())
}

type CustomEmbeddedLink struct {
	link.LinkedObject
}

func (l *CustomEmbeddedLink) String() string {
	return fmt.Sprintf("%v {%v}", reflect.TypeOf(l).String(), reflect.TypeOf(l.LinkedObject).String())
}

type CustomLink struct {
	parent, child interface{}
}

func (l *CustomLink) Link(i interface{}) error {
	l.child = i
	switch il := i.(type) {
	case *CustomLink:
		il.parent = l

	case *link.LinkedObject, *CustomEmbeddedLink, *CustomEmbeddedIfaceLink:
		return il.(link.Linker).Link(i)

	}

	return nil
}
func (l *CustomLink) Parent() interface{} { return l.parent }
func (l *CustomLink) Child() interface{}  { return l.child }
func (l *CustomLink) String() string      { return reflect.TypeOf(l).String() }
