package lib

import (
	"fmt"

	"dario.cat/mergo"
	"laatoo.io/sdk/utils"
)

func MergeMaps(obj1, obj2 utils.StringMap) (utils.StringMap, error) {
	if obj1 == nil {
		return obj2, nil
	}
	if obj2 == nil {
		return obj1, nil
	}
	res := make(map[string]interface{}, len(obj1))
	for k, v := range obj1 {
		res[k] = v
	}
	options := []func(*mergo.Config){mergo.WithOverride}

	mergo.Merge(&res, obj2, options...)
	return utils.StringMap(res), nil
}

func ConvertToStringsMap(val map[string]interface{}) utils.StringsMap {
	if val == nil {
		return nil
	}
	strMp := make(utils.StringsMap)
	for k, v := range val {
		strV := fmt.Sprintf("%v", v)
		strMp[k] = strV
	}
	return strMp
}
