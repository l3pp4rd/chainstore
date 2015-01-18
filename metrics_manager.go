package chainstore

import (
	"fmt"
	"time"

	"github.com/rcrowley/go-metrics"
)

// @TODO: is this really necessary? could be just an example
type metricsManager struct {
	namespace string
	registry  metrics.Registry
	chain     Store
}

func Metricable(namespace string, registry metrics.Registry, store Store) Store {
	return &metricsManager{
		namespace: namespace,
		registry:  registry,
		chain:     store,
	}
}

func (m *metricsManager) Open() (err error) {
	_, err = m.measure("Open", func() ([]byte, error) {
		err := m.chain.Open()
		return nil, err
	})
	return
}

func (m *metricsManager) Close() (err error) {
	_, err = m.measure("Close", func() ([]byte, error) {
		err := m.chain.Close()
		return nil, err
	})
	return
}

func (m *metricsManager) Put(key string, val []byte) (err error) {
	_, err = m.measure("Put", func() ([]byte, error) {
		err := m.chain.Put(key, val)
		return nil, err
	})
	return
}

func (m *metricsManager) Get(key string) (val []byte, err error) {
	val, err = m.measure("Get", func() ([]byte, error) {
		val, err := m.chain.Get(key)
		return val, err
	})
	return
}

func (m *metricsManager) Del(key string) (err error) {
	_, err = m.measure("Del", func() ([]byte, error) {
		err := m.chain.Del(key)
		return nil, err
	})
	return
}

func (m *metricsManager) measure(method string, fn func() ([]byte, error)) ([]byte, error) {
	ns := fmt.Sprintf("%s.%s", m.namespace, method)
	metric := metrics.GetOrRegisterTimer(ns, m.registry)
	t := time.Now()
	val, err := fn()
	metric.UpdateSince(t)
	return val, err
}
