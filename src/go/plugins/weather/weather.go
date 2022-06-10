package weather

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.zabbix.com/ap/plugin-support/plugin"
)

type Plugin struct {
	plugin.Base
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	if len(params) != 1 {
		return nil, errors.New("porra loca!!")
	}

	res, err := http.Get(fmt.Sprintf("https://wttr.in/~%s?format=%%t", params[0]))
	if err != nil {
		return nil, err
	}

	temp, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return string(temp)[0 : len(temp)-4], nil
}
