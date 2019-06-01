/**
 * registry.go - balancers registry
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

package balance

import (
	"github.com/yyyar/gobetween/balance/middleware"
	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/core"
)

/**
 * Registry of available Balancers
 */
var registry = make(map[string]func(config.BalanceConfig) interface{})

/**
 * Initialize type registry
 */
func init() {
	registry["leastconn"] = NewLeastconnBalancer
	registry["roundrobin"] = NewRoundrobinBalancer
	registry["weight"] = NewWeightBalancer
	registry["iphash"] = NewIphashBalancer
	registry["iphash1"] = NewIphash1Balancer
	registry["iphash2"] = NewIphash2Balancer
	registry["leastbandwidth"] = NewLeastbandwidthBalancer
}

/**
 * Create new Balancer based on strategy
 */
func New(sniConf *config.Sni, cfg config.BalanceConfig) core.Balancer {
	balancer := registry[cfg.Kind](cfg).(core.Balancer)

	if sniConf == nil {
		return balancer
	}

	return &middleware.SniBalancer{
		SniConf:  sniConf,
		Delegate: balancer,
	}
}
