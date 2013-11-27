package link

import (
	"errors"
	"fmt"
)

func Chain(root Linker, links ...Linker) error {
	if len(links) < 1 {
		return errors.New("There are no links to chain together.")
	}
	for i, l := range links {

		switch {
		case i == 0:
			if err := root.Link(l); err != nil {
				return err
			}
			fallthrough
		case i < len(links)-1:
			if err := l.Link(links[i+1]); err != nil {
				return err
			}
		}
	}
	return nil
}

// makes any type embedding LinkedObject implement this interface
type Linker interface {
	Link(interface{}) error
	Parent() interface{}
	Child() interface{}
}

type NotLinkerErr struct {
	fmt.Stringer
}

func (err NotLinkerErr) Error() string {
	return fmt.Sprintf("%v is not a valid linker.", err.Stringer)
}

type CantSetErr struct {
	fmt.Stringer
}

func (err CantSetErr) Error() string {
	return fmt.Sprintf("%v can not be set using reflection.", err.Stringer)
}
func LinkerInfo(l Linker) {
	fmt.Println(l, "\n")
}
