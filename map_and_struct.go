package main

import (
    "encoding/json"
    "fmt"
    "github.com/golang/protobuf/jsonpb"
    "github.com/golang/protobuf/ptypes/struct"
)

func MapToStruct(inputMap map[string]interface{}) (*structpb.Struct, error) {
    structObj := &structpb.Struct{}
    structObj.Fields = make(map[string]*structpb.Value)

    for key, value := range inputMap {
        fieldValue, err := structpb.NewValue(value)
        if err != nil {
            return nil, err
        }
        structObj.Fields[key] = fieldValue
    }

    return structObj, nil
}

func StructToMap(inputStruct *structpb.Struct) (map[string]interface{}, error) {
    if inputStruct == nil {
        return nil, nil
    }

    resultMap := make(map[string]interface{})

    for key, value := range inputStruct.Fields {
        var decodedValue interface{}
        if err := json.Unmarshal(value.Kind.(*structpb.Value_JsonValue).JsonValue, &decodedValue); err != nil {
            return nil, err
        }
        resultMap[key] = decodedValue
    }

    return resultMap, nil
}


func main() {
     inputMap := map[string]interface{}{
        "foo": "bar",
        "number": 42,
        "nested": map[string]interface{}{
            "nestedKey": "nestedValue",
        },
    }

    structObj, err := MapToStruct(inputMap)
    if err != nil {
        fmt.Printf("Error converting map to Struct: %v\n", err)
        return
    }

    // Print the resulting Struct as JSON
    marshaller := jsonpb.Marshaler{}
    jsonString, _ := marshaller.MarshalToString(structObj)
    fmt.Println(jsonString)

  // ----------------------------------------------------
  
    // JSON representation of a google.protobuf.Struct
    jsonString := `{
        "foo": "bar",
        "number": 42,
        "nested": {
            "nestedKey": "nestedValue"
        }
    }`

    structObj := &structpb.Struct{}
    err := jsonpb.UnmarshalString(jsonString, structObj)
    if err != nil {
        fmt.Printf("Error unmarshaling JSON to Struct: %v\n", err)
        return
    }

    resultMap, err := StructToMap(structObj)
    if err != nil {
        fmt.Printf("Error converting Struct to map: %v\n", err)
        return
    }

    // Print the resulting map
    fmt.Println(resultMap)
}
