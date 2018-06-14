// Made by Martin Alejandro Flores Ramirez
// martinfloresram@gmail.com
// Documents Querying
// https://github.com/og2

package quiver

import (
	"strconv"
	"strings"
	"unicode"
)

type M map[string]interface{}

func AND(args ...string) M {

	if len(args) <= 0 {
		return nil
	}
	parentMap := M{
		"$and": []M{},
	}

	for _, arg := range args {
		conditionMap := M{}
		e := SpaceMap(arg)
		eTemp := e
		if thisHas(e, "=>") { //there's an AND
			objA := strings.Split(e, "=>")
			conditionMap[objA[0]] = M{"$elemMatch": M{}} //caracteristicas
			eTemp = objA[1]
		}
		if thisHas(eTemp, "&&") { //there's an AND
			splitCondition := strings.Split(eTemp, "&&")
			expMap := M{}
			for _, condition := range splitCondition {
				if thisHas(condition, "==") { //isArray
					obj := strings.Split(condition, "==")
					if isNumber(obj[1]) {
						expMap[obj[0]] = stringToInteger(obj[1])
					} else {
						expMap[obj[0]] = obj[1]
					}
				}
				if thisHas(condition, ">=") { //isArray
					obj := strings.Split(condition, ">=")
					if isNumber(obj[1]) {
						expMap[obj[0]] = M{"$gte": stringToInteger(obj[1])}
					} else {
						expMap[obj[0]] = M{"$gte": obj[1]}
					}
				}
				if thisHas(condition, "<=") { //isArray
					obj := strings.Split(condition, "<=")
					if isNumber(obj[1]) {
						expMap[obj[0]] = M{"$lte": stringToInteger(obj[1])}
					} else {
						expMap[obj[0]] = M{"$lte": obj[1]}
					}
				}
			}
			conditionMap[strings.Split(e, "=>")[0]].(M)["$elemMatch"] = expMap
			parentMap["$and"] = append(parentMap["$and"].([]M), conditionMap)
		} else {
			if thisHas(eTemp, "==") { //isArray
				obj := strings.Split(eTemp, "==")
				if isNumber(obj[1]) {
					parentMap["$and"] = append(parentMap["$and"].([]M), M{obj[0]: stringToInteger(obj[1])})
				} else {
					parentMap["$and"] = append(parentMap["$and"].([]M), M{obj[0]: obj[1]})
				}
			}
		}
	}
	return parentMap
}
func Query(obj interface{}) interface{} {

	switch obj.(type) {
	case int:
		break
	case float64:
		break
	case string:
		break
	case M:
		obj = obj.(M)
		break
	default:
		break
	}

	return obj
}

func thisHas(str string, target string) bool {
	if strings.Contains(str, target) {
		return true
	}
	return false
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func stringToInteger(s string) int {
	output, err := strconv.Atoi(s)
	if err != nil {
		output = 0
	}
	return output
}
