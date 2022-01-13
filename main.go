package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"strings"

	log "github.com/anderskvist/GoHelpers/log"
)

func main() {
	log.Info("Fetching data from danskretursystem.dk")

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormField("action")
	io.Copy(fw, strings.NewReader("getMapData"))

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://danskretursystem.dk/wp/wp-admin/admin-ajax.php", &b)
	r.Header.Set("Content-Type", w.FormDataContentType())

	debug, _ := httputil.DumpRequestOut(r, true)
	log.Debug(fmt.Sprintf("%s", debug))

	resp, _ := client.Do(r)
	log.Info(fmt.Sprintf("%s", resp.Status))
}
