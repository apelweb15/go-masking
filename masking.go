package masking

import (
	"encoding/json"
	"fmt"
	"github.com/haifenghuang/orderedmap"
	"reflect"
	"regexp"
	"strings"
)

var defaultConfig = &Config{
	MaskingKeys:       DefaultMaskingKey,
	MaskingPercentage: DefaultMaskingPercentage,
	MaskingPosition:   DefaultMaskingPosition,
	MaskingCharacter:  DefaultMaskingCharacter,
}
var MaskConfig *Config = defaultConfig

func MaskString(input string, configs ...*Config) string {

	config := MaskConfig

	if len(configs) > 0 {
		config = configs[0]
	}

	if input == "" {
		return ""
	}

	maskingChar := config.MaskingCharacter
	if maskingChar == "" {
		maskingChar = DefaultMaskingCharacter
	}

	maskingPercentage := config.MaskingPercentage
	if maskingPercentage <= 0 {
		maskingPercentage = DefaultMaskingPercentage
	}

	totalLen := len(input)
	maskLen := (totalLen * maskingPercentage) / 100
	if maskLen >= totalLen {
		return strings.Repeat(maskingChar, totalLen)
	}

	switch config.MaskingPosition {
	case MaskLeft:
		visible := input[maskLen:]
		return strings.Repeat(maskingChar, maskLen) + visible

	case MaskRight:
		visible := input[:totalLen-maskLen]
		return visible + strings.Repeat(maskingChar, maskLen)

	case MaskCenter:
		visibleSide := (totalLen - maskLen) / 2
		left := input[:visibleSide]
		right := input[visibleSide+maskLen:]
		return left + strings.Repeat(maskingChar, maskLen) + right

	default:
		return strings.Repeat(maskingChar, totalLen)
	}
}

func MaskSensitive(data interface{}, configs ...*Config) interface{} {
	visited := map[uintptr]bool{}
	result := maskRecursive(reflect.ValueOf(data), visited, configs...)
	rv := reflect.ValueOf(result)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		return nil
	}

	return rv.Interface()
}

func maskRecursive(val reflect.Value, visited map[uintptr]bool, configs ...*Config) interface{} {
	config := MaskConfig
	if len(configs) > 0 {
		config = configs[0]
	}
	if !val.IsValid() {
		return nil
	}

	if val.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		if !val.IsNil() {
			err := val.Interface().(error)
			return err.Error()
		}
		return nil
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		ptr := val.Pointer()
		if visited[ptr] {
			return val.Interface()
		}
		visited[ptr] = true

		newPtr := reflect.New(val.Elem().Type())
		maskedElem := maskRecursive(val.Elem(), visited, config)
		safeSet(newPtr.Elem(), maskedElem)
		return newPtr.Interface()
	}

	switch val.Kind() {

	case reflect.Interface:
		return maskRecursive(val.Elem(), visited, config)

	case reflect.Struct:
		newStruct := reflect.New(val.Type()).Elem()
		t := val.Type()
		for i := 0; i < val.NumField(); i++ {

			field := t.Field(i)
			if field.PkgPath != "" {
				continue
			}
			if !newStruct.Field(i).CanSet() {
				continue
			}

			fieldName := field.Name
			fieldVal := val.Field(i)

			if isSensitiveKey(fieldName, config.MaskingKeys...) && fieldVal.Kind() == reflect.String {
				newStruct.Field(i).SetString(MaskString(fieldVal.String(), config))
				continue
			}

			if isSensitiveKey(fieldName, config.MaskingKeys...) && isInteger(fieldVal.Kind()) {
				newStruct.Field(i).SetInt(0)
				continue
			}
			if isSensitiveKey(fieldName, config.MaskingKeys...) && isFloat(fieldVal.Kind()) {
				newStruct.Field(i).SetFloat(0.0)
				continue
			}

			maskedVal := maskRecursive(fieldVal, visited, config)
			safeSet(newStruct.Field(i), maskedVal)

		}
		return newStruct.Interface()

	case reflect.Map:
		if val.IsNil() {
			return nil
		}
		newMap := reflect.MakeMap(val.Type())
		for _, key := range val.MapKeys() {
			keyStr := fmt.Sprintf("%v", key.Interface())
			valField := val.MapIndex(key)

			v := valField
			for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
				if v.IsNil() {
					break
				}
				v = v.Elem()
			}
			if isSensitiveKey(keyStr, config.MaskingKeys...) && v.Kind() == reflect.String {
				masked := MaskString(v.String(), config)
				newMap.SetMapIndex(key, reflect.ValueOf(masked))
				continue
			}

			if isSensitiveKey(keyStr, config.MaskingKeys...) && isInteger(v.Kind()) {
				newMap.SetMapIndex(key, reflect.ValueOf(0))
				continue
			}

			if isSensitiveKey(keyStr, config.MaskingKeys...) && isFloat(v.Kind()) {
				newMap.SetMapIndex(key, reflect.ValueOf(0.0))
				continue
			}

			maskedVal := maskRecursive(valField, visited, config)
			fv := reflect.ValueOf(maskedVal)
			if fv.IsValid() && fv.Type().AssignableTo(val.Type().Elem()) {
				newMap.SetMapIndex(key, fv)
			} else {
				newMap.SetMapIndex(key, reflect.Zero(val.Type().Elem()))
			}
		}
		return newMap.Interface()

	case reflect.Slice, reflect.Array:
		if val.Kind() == reflect.Slice && val.IsNil() {
			return nil
		}
		newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			maskedElem := maskRecursive(val.Index(i), visited, config)
			safeSet(newSlice.Index(i), maskedElem)
		}
		return newSlice.Interface()

	case reflect.String:

		omp := orderedmap.New()
		if err := json.Unmarshal([]byte(val.String()), &omp); err == nil {
			processOrderedMap(omp, config)
			b, err := json.Marshal(omp)
			if err == nil {
				return string(b)
			}
		}
		return val.String()

	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return val.Interface()

	case reflect.UnsafePointer, reflect.Chan, reflect.Func, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("<unsupported:%s>", val.Kind().String())

	default:
		return val.Interface()
	}
}

func processOrderedMap(omp *orderedmap.OrderedMap, config *Config) {
	for _, key := range omp.Keys() {
		val, _ := omp.Get(key)
		switch v := val.(type) {
		case string:
			if isSensitiveKey(key, config.MaskingKeys...) {
				omp.Set(key, MaskString(v, config))
			}

		case *orderedmap.OrderedMap:
			processOrderedMap(v, config)
		case map[string]interface{}:
			nested := orderedmap.New()
			for nk, nv := range v {
				nested.Set(nk, nv)
			}
			processOrderedMap(nested, config)
			omp.Set(key, nested)
		case []interface{}:
			for i, elem := range v {
				switch e := elem.(type) {
				case *orderedmap.OrderedMap:
					processOrderedMap(e, config)
				case map[string]interface{}:
					nested := orderedmap.New()
					for nk, nv := range e {
						nested.Set(nk, nv)
					}
					processOrderedMap(nested, config)
					v[i] = nested
				}
			}
			omp.Set(key, v)

		default:
		}
	}
}

func isSensitiveKey(key string, sensitiveKeys ...string) bool {
	if len(sensitiveKeys) == 0 {
		sensitiveKeys = DefaultMaskingKey
	}
	find := replaceNonAlpha(strings.ToLower(key))
	for _, s := range sensitiveKeys {
		if strings.Contains(find, replaceNonAlpha(strings.ToLower(s))) {
			return true
		}
	}
	return false
}

func replaceNonAlpha(s string) string {
	re := regexp.MustCompile(`[^A-Za-z]+`)
	return re.ReplaceAllString(s, "")
}

func safeSet(dst reflect.Value, val interface{}) {
	fv := reflect.ValueOf(val)
	if !fv.IsValid() {
		dst.Set(reflect.Zero(dst.Type()))
		return
	}

	if fv.Type().AssignableTo(dst.Type()) {
		dst.Set(fv)
		return
	}

	if fv.Type().ConvertibleTo(dst.Type()) {
		dst.Set(fv.Convert(dst.Type()))
		return
	}
	dst.Set(reflect.Zero(dst.Type()))
}
