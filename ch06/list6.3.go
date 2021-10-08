package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8">
		<title>Example</title>
	</head>
	<body>
<hr/>
		<p>
		{{ .Date | dateFormat "Jan 2, 2006" }}
		</p>
<hr/>	
</body>
</html>`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

/**

 */
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

/**

 */
func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t, err := template.New("date").Funcs(funcMap).Parse(tpl) // Создание нового екземпляра template
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}
	// Передача карти с дополнительными функциями механизму шаблонов

	data := &struct {
		Date time.Time
	}{
		Date: time.Now(),
	}
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
