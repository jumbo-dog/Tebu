package reflection

import (
	"reflect"
)

func AttributeValues(r, unit interface{}) {
	unitValue := reflect.ValueOf(r).Elem()
	unitType := unitValue.Type()
	unitToFill := reflect.ValueOf(unit).Elem()

	if !unitToFill.IsValid() {
		return
	}
	unitToFillType := unitToFill.Type()

	for i := 0; i < unitValue.NumField(); i++ {
		consumerUnitName := unitType.Field(i).Name
		srcFieldStructValue := unitValue.Field(i)
		fieldToBeInserted := unitValue.FieldByName(consumerUnitName)

		for i := 0; i < unitToFill.NumField(); i++ {
			targetFieldName := unitToFillType.Field(i).Name
			targetValue := unitToFill.Field(i)

			if consumerUnitName == targetFieldName {
				switch targetValue.Kind() {
				case reflect.String:
					if srcFieldStructValue.String() == "" {
						fieldToBeInserted.Set(targetValue)
					}
				case reflect.Int64, reflect.Uint8, reflect.Uint16:
					if srcFieldStructValue.Interface() == reflect.Zero(targetValue.Type()).Interface() {
						fieldToBeInserted.Set(targetValue)
					}
				case reflect.Map:
					if targetValue.IsNil() || srcFieldStructValue.IsNil() {
						if srcFieldStructValue.IsNil() {
							fieldToBeInserted.Set(targetValue)
						}
					} else {
						mergedMap := reflect.MakeMap(targetValue.Type())
						targetMapKeys := targetValue.MapKeys()
						for _, key := range targetMapKeys {
							mergedMap.SetMapIndex(key, targetValue.MapIndex(key))
						}
						srcMapKeys := srcFieldStructValue.MapKeys()
						for _, key := range srcMapKeys {
							mergedMap.SetMapIndex(key, srcFieldStructValue.MapIndex(key))
						}
						fieldToBeInserted.Set(mergedMap)
					}
				}
			}
		}
		if srcFieldStructValue.Kind() == reflect.Ptr {
			continue
		}
	}
}

func InitializeStructIfNil(r, unit interface{}, fieldName string) {
	field := reflect.ValueOf(r).Elem().FieldByName(fieldName)

	if field.IsNil() {
		field.Set(reflect.ValueOf(unit))
	}
}