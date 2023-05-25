package html

type HTMLparserInterface interface {
	Create(templateName string, data interface{}) (string, error)
}
