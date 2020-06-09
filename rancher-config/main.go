package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

func newTemplate() (*template.Template, error) {
	t := template.New("base")
	var funcMap template.FuncMap = map[string]interface{}{}
	funcMap["include"] = func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := t.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	tmpl, err := t.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).ParseGlob("templates/*")
	return tmpl, err
}

func readConfigJSON() (templateData, error) {
	td := templateData{}
	f, err := os.Open("config/config.json")
	if err != nil {
		return td, err
	}
	defer f.Close()

	d := json.NewDecoder(f)
	err = d.Decode(&td)
	if err != nil {
		return td, err
	}
	return td, nil
}

func main() {

	tmpl, err := newTemplate()
	if err != nil {
		log.Fatal(err)
	}
	data, err := readConfigJSON()
	if err != nil {
		log.Fatal(err)
	}
	myTmpl := singleTemplate{
		tmpl: tmpl,
		data: data,
	}

	listenOn := ":80"
	http.HandleFunc("/", myTmpl.serveTemplate)

	log.Printf("Listening on %s", listenOn)
	err = http.ListenAndServe(listenOn, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type templateData struct {
	IP                string
	Zabbix            string `json:"zabbix"`
	DNS               string
	Mounts            [][]string `json:"mounts"`
	SSHAuthorizedKeys []string   `json:"ssh_authorized_keys"`
	RegistryMirror    struct {
		URL string `json:"url"`
	} `json:"registry-mirror"`
}

type singleTemplate struct {
	tmpl *template.Template
	data templateData
}

func (t singleTemplate) serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := newTemplate() //this line refreshes the template at every request
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	tp := filepath.Clean(r.URL.Path)
	tp = strings.Trim(tp, "/")
	fp := filepath.Join("templates", tp)

	t.data.IP = strings.Split(r.RemoteAddr, ":")[0]

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	err = tmpl.ExecuteTemplate(w, tp, t.data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
