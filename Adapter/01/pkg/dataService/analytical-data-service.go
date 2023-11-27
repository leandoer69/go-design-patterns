package dataService

// XML format
type AnalyticalDataService interface {
	SendXmlData()
}

type XMLDocument struct {
}

func (doc XMLDocument) SendXmlData() {
	println("Отпарвка XML документа!")
}
