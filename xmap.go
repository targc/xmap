package xmap

import "strings"

type XMap map[string]interface{}

func (m XMap) M(key string) XMap {
	val, ok := m[key].(map[string]interface{})

	if !ok {
		return make(XMap)
	}

	return val
}

func (m XMap) I(key string) interface{} {
	val, ok := m[key].(interface{})

	if !ok {
		return nil
	}

	return val
}

func (m XMap) Get(key string) interface{} {
	keys := strings.Split(key, ".")

	if len(keys) == 0 {
		return nil
	}

	if len(keys) == 1 {
		return m.I(keys[0])
	}

	tmp := make(XMap)

	for i, k := range keys {

		if i == 0 {
			tmp = m.M(k)
			continue
		}

		if tmp == nil {
			return nil
		}

		if i == len(keys)-1 {
			return tmp.I(k)
		}

		tmp = tmp.M(k)
	}

	return nil
}

func (m XMap) GetS(key string) string {
	val := m.Get(key)

	if val == nil {
		return ""
	}

	valS, ok := val.(string)

	if !ok {
		return ""
	}

	return valS
}
