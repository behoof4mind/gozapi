package host

import "gozapi/zabbix"

type Host struct {
	Params HostParams
}

type HostParams struct {

}

func (h *Host) Get(z *zabbix.Zabbix) {

}