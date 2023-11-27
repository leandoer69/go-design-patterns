package data

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	println("Отправка XML данных!")
}
