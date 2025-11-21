package main

import "fmt"

func GetTypeName(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case complex128:
		return "complex128"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	s := []any{any(1), any(3.13), any("hello"), any(1 + 2i)}
	for _, e := range s {
		fmt.Println(e, ":", GetTypeName(e))
	}
}
