package main

import (
	"go_design_patterns/Adapter/01/pkg/data"
)

func main() {
	xmlDoc := data.JsonDocument{}
	jsonToXmlAdapter := data.JsonDocumentAdapter{&xmlDoc}
	jsonToXmlAdapter.SendXmlData()
}
