package utils

func GetPtrValue[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func GetValuePtr[T int | string | bool](value T) *T {
	return &value
}
