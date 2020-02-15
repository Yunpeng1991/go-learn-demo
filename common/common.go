package common

func IsNull(param interface{}) bool {
	if nil == param {
		return true
	}
	return false
}
