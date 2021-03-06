package services

import (
	"net/http"
	"os"
	"zarbat_test/internal/logging"
	l "zarbat_test/internal/logging"
)

func DialHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** DialHandler")
	xml, err := os.ReadFile("xml/dial.xml")
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}

func NumberHandler(w http.ResponseWriter, r *http.Request) {
	l.Debug.Println("******************************** NumberHandler")
	xml, err := os.ReadFile("xml/number.xml")
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	l.Debug.Println(string(xml))
	w.Write([]byte(xml))
}
