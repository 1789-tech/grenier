package driver

import (
	"fmt"
	"sort"
	"sync"
)

// Factory builds a Driver from its parsed manifest. Drivers register a factory
// under their manifest ID at init time.
type Factory func(m Manifest) (Driver, error)

var (
	regMu    sync.RWMutex
	registry = map[string]Factory{}
)

// Register makes a driver available by manifest ID. Called from driver package
// init functions.
func Register(id string, f Factory) {
	regMu.Lock()
	defer regMu.Unlock()
	registry[id] = f
}

// Load builds the driver registered for the manifest's ID.
func Load(m Manifest) (Driver, error) {
	regMu.RLock()
	f, ok := registry[m.ID]
	regMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("no driver registered for manifest id %q", m.ID)
	}
	return f(m)
}

// Registered lists the IDs of all registered drivers.
func Registered() []string {
	regMu.RLock()
	defer regMu.RUnlock()
	ids := make([]string, 0, len(registry))
	for id := range registry {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
