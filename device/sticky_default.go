//go:build !linux

package device

import (
	"github.com/sagernet/amneziawg-go/conn"
	"github.com/sagernet/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
