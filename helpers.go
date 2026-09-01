package inttegro

// String returns a pointer to the provided string value.
//
// Useful for setting optional string fields that require pointers.
//
// Example:
//
//	params := OrderCreateParams{
//	    Number: inttegro.String("ORD-2023-00123"),
//	}
func String(v string) *string { return &v }

// Bool returns a pointer to the provided boolean value.
//
// Useful for setting optional boolean flags that require pointers.
// Many Inttegro API parameters use *bool to distinguish between
// explicit false and unset (nil).
//
// Example:
//
//	params := OrderCreateParams{
//	    ExecutePayment: inttegro.Bool(true),
//	    Finalize:       inttegro.Bool(true),
//	}
func Bool(v bool) *bool { return &v }

// Int returns a pointer to the provided int value.
//
// Useful for setting optional integer fields that require pointers.
func Int(v int) *int { return &v }

// Int64 returns a pointer to the provided int64 value.
//
// Useful for setting optional int64 fields that require pointers.
func Int64(v int64) *int64 { return &v }
