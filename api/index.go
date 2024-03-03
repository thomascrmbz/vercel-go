package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html/>
<html>
    <head>
        <title>Go Runtime</title>
        <script>
            window.va =
                window.va ||
                function () {
                    (window.vaq = window.vaq || []).push(arguments);
                };
        </script>
        <script defer src="/_vercel/insights/script.js"></script>
    </head>
    <body>
        <h1>Hello From Go</h1>
    </body>
</html>`)
}
