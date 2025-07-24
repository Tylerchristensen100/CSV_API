package docs

import (
	"net/http"
)

func README(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "https://github.com/Tylerchristensen100/CSV_API", http.StatusFound)
}
