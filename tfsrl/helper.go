package tfsrl

import (
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// replaceInKeys function
func replaceInKeys(i interface{}, old, new string) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		//log.Printf("[DEBUG] replaceInKeys: map[interface{}]interface{}: %v\n", x)
		nm := map[string]interface{}{}

		for k, v := range x {
			ks := k.(string)
			nm[strings.Replace(ks, old, new, -1)] = replaceInKeys(v, old, new)
		}
		//log.Printf("[DEBUG] replaceInKeys: map[interface{}]interface{}: %v\n", nm)
		return nm
	case []interface{}:
		//log.Printf("[DEBUG] replaceInKeys: []interface{}: %v\n", x)
		for i, v := range x {
			//log.Printf("[DEBUG] replaceInKeys: []interface{}: %v\n", v)
			x[i] = replaceInKeys(v, old, new)
			//x[i] = strings.Replace(v.(string), old, new, -1)
		}
	case map[string]interface{}:
		//log.Printf("[DEBUG] replaceInKeys: default: %v\n", i)
		nm := map[string]interface{}{}

		for k, v := range x {
			nm[strings.Replace(k, old, new, -1)] = replaceInKeys(v, old, new)
		}
		//log.Printf("[DEBUG] replaceInKeys: map[interface{}]interface{}: %v\n", nm)
		return nm
	}
	//log.Printf("[DEBUG] replaceInKeys: return: %v\n", i)
	return i
}

func getResourceListKey(rn, rk *string, d *schema.ResourceData) (key string, err error) {
	t := d.Get(*rn).([]interface{})
	log.Debugf("GetResourceListKey Data: %v", t)
	for _, i := range t {
		switch x := i.(type) {
		case map[string]interface{}:
			if v, ok := x[*rk]; ok {
				switch v.(type) {
				case string:
					key = v.(string)
				case int:
					key = strconv.Itoa(v.(int))
				}
				break
			}
		}
	}
	log.Debugf("GetResourceListKey Key Info: %s", key)
	return key, nil
}
