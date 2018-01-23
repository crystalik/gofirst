package main

import (
        "net/http"
        "adapter"
        "encoding/xml"
)

func makeHandler (fn func (_ interface{}) []byte, dat interface{}) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		decoder := xml.NewDecoder(r.Body)
		error := decoder.Decode(&dat)
		if error != nil {
			panic(error)
		}
		defer r.Body.Close()
		w.Write(fn(dat))
	}
}

func main() {

}

func init () {
	LastTime := new(adapter.LastSessionTime)
	Guids := new(adapter.Keys)

	http.HandleFunc("/1c/adapter/getChanges", makeHandler(adapter.GetChanges, LastTime))
	http.HandleFunc("/1c/adapter/readIntegrationMessages", makeHandler(adapter.ReadIntegrationMessages, Guids))
        err := http.ListenAndServe(":9005", nil)
        if err != nil {
                panic(err.Error())
        }
}
