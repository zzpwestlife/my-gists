package main

func structToMap(st *myStruct) map[string]interface{} {
    flatMap := make(map[string]interface{})
    v := reflect.ValueOf(holder)
    flattenStruct("", v, flatMap)
    return flatMap
}

func flattenStruct(prefix string, v reflect.Value, m map[string]interface{}) {
    if v.Kind() == reflect.Ptr {
       v = v.Elem()
    }
    t := v.Type()
    for i := 0; i < v.NumField(); i++ {
       field := t.Field(i)
       value := v.Field(i)
       fieldName := field.Name
       // 排除自动生成的字段
       if slicex.Contains([]string{"state", "sizeCache", "unknownFields"}, fieldName) {
          continue
       }
       // 驼峰转蛇形
       fieldName = stringx.CamelToSnake(fieldName)
       if prefix != "" {
          fieldName = prefix + "." + fieldName
       }
       if value.Kind() == reflect.Ptr && !value.IsNil() {
          value = value.Elem()
       }
       if value.Kind() == reflect.Struct {
          flattenStruct(fieldName, value, m)
       } else if value.Kind() == reflect.Ptr && value.IsNil() {
          // nil 直接去掉 (如下图的 Cellphone, LocalName.MiddleName)
          continue
       } else if value.Kind() == reflect.Slice && value.Len() == 0 {
          // 空 slice 直接去掉 (如下图的 TaxResidencies)
          continue
       } else {
          m[fieldName] = value.Interface()
       }
    }
}