package controllers

type Context interface {
	Param(key string) string
	ShouldBindJSON(obj interface{}) (string error)
	JSON(code int, obj interface{})
}
