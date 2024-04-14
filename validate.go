/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

type number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
	~int8 | ~int16 | ~int32 | ~int64
	~int | ~uint | ~float32 | ~float64
}

func number_validator[T number](num T) string {
	return ""
}
