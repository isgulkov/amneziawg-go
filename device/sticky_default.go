//go:build !linux

package device

import (
	"github.com/isgulkov/amneziawg-go/conn"
	"github.com/isgulkov/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
