package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"text/template"
)

func main() {

	var b bytes.Buffer

	f, err := os.OpenFile("index.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	io.Copy(&b, f)

	template := template.Must(template.New("index").Parse(b.String()))

	template.Execute(os.Stdout, map[string]interface{}{
		"username":     "gowi.coffee",
		"confirmToken": "12345",
	})
}
