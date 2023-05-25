package pdf

type PDFGeneratorInterface interface {
	Create(htmlFile string) (string, error)
}
