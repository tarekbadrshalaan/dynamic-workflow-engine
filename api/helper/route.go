package helper

import (
	"dwf/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route :
type Route struct {
	Method    string            //HTTP method
	Path      string            //url endpoint
	Handle    httprouter.Handle //Controller function which dispatches the right HTML page and/or data for each route
	IsLogging bool
}

// Handler :
func (r *Route) Handler() httprouter.Handle {
	if r.IsLogging {
		r.WithLogger()
	}

	return httprouter.Handle(r.Handle)
}

// WithLogger :
func (r *Route) WithLogger() *Route {
	r.Handle = Logmid(r.Handle)
	return r
}

// Logmid : logging midleware
func Logmid(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger.Infof("[%s] on: %s", r.Method, r.URL)
		next(w, r, ps)
	}
}

// WriteResponseJSON :
func WriteResponseJSON(w http.ResponseWriter, v interface{}, stateCode int) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(stateCode)
	w.Write(data)
}

// WriteResponseError :
func WriteResponseError(w http.ResponseWriter, err error, stateCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	msg := fmt.Sprintf(`{"errorText":"%v"}`, err)
	http.Error(w, msg, stateCode)
}

// ReadJSON :
func ReadJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
