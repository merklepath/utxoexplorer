package webui

import (
	"fmt"
	"net/http"

	"github.com/merklepath/utxoexplorer/templates"
)

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.New()
	err := tmpl.Render(w, "index.tmpl", map[string]interface{}{
		"Title": "Home",
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
