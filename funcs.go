package lib

import (
	"dario.cat/mergo"
	"laatoo.io/sdk/utils"
)

func MergeMaps(obj1, obj2 utils.StringMap) utils.StringMap {
	if obj1 == nil {
		return obj2
	}
	if obj2 == nil {
		return obj1
	}
	res := make(map[string]interface{})
	mergo.Merge(&res, map[string]interface{}(obj1))
	mergo.Merge(&res, map[string]interface{}(obj2))
	return utils.StringMap(res)
}
