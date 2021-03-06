package link

import (
	"fmt"
	"reflect"
)

// LinkedObject should be embedded
type LinkedObject struct {
	provider *Provider

	// reference to the websocket
	connID uint64

	// id for the linked object, generated by client
	id uint64

	parent, child Linker

	container Linker
}
type Provider int

func (l *LinkedObject) Link(i interface{}) error {
	_, ok := i.(Linker)
	if !ok {
		return NotLinkerErr{i.(fmt.Stringer)}
	} else {
		switch lobj := i.(type) {
		case *LinkedObject:
			lobj.parent, l.child = l, lobj
			return nil
		default:
			linkval := reflect.ValueOf(i)
			if linkval.Kind() == reflect.Ptr {
				linkval = linkval.Elem()
			}
			fval := linkval.FieldByNameFunc(func(s string) bool {
				switch s {
				case "LinkedObject", "*LinkedObject":
					if f, ok := linkval.Type().FieldByName(s); ok && f.Anonymous {
						return true
					}
				}
				return false
			})
			if fval.IsValid() {
				if fval.Kind() != reflect.Ptr && fval.CanAddr() {
					l.child = fval.Addr().Interface().(Linker)
				} else {
					l.child = fval.Interface().(Linker)
				}
				l.child.(*LinkedObject).parent = l
				return nil
			}
			l.child = &LinkedObject{parent: l, container: i.(Linker)}
			return nil
		}
	}
}
func (l *LinkedObject) Parent() interface{} {
	if l.container != nil {
		return l.container.Parent()
	}
	return l.parent
}
func (l *LinkedObject) Child() interface{} {
	if l.container != nil {
		return l.container.Child()
	}
	return l.child
}
func (l *LinkedObject) String() string {
	if l.container != nil {
		return reflect.TypeOf(l.container).String()
	}
	return reflect.TypeOf(l).String()
}
