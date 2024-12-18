package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var page = []byte(`
<!doctype html>
<html>
<head>
    <title>Tilt Example</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>
</head>

<body>
<div>
    <h1>Tilt Example</h1>
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64" enable-background="new 0 0 64 64">
    <path fill="#f7b600" d="m2 61l8.6-3-6.5-3z"/>
    <path fill="#ffdd7d" d="m26.9 36.4l-12.1-12.2-2 5.6z"/>
    <path fill="#f7b600" d="m12.8 29.8l-2.2 6.3 26.8 12.5 1.3-.4-11.8-11.8z"/>
    <path fill="#ffdd7d" d="m8.5 42.4l20 9.3 8.9-3.1-26.8-12.5z"/>
    <path fill="#f7b600" d="m6.3 48.7l13.2 6.2 9-3.2-20-9.3z"/>
    <path fill="#ffdd7d" d="m6.3 48.7l-2.2 6.3 6.5 3 8.9-3.1z"/>
    <path d="m31.9 31.2c6.7 6.6 10.2 14 7.8 16.4-2.5 2.4-9.9-1-16.7-7.7-6.7-6.6-10.2-14-7.8-16.4 2.5-2.4 9.9 1.1 16.7 7.7" fill="#493816"/>
    <path d="m23.5 14.5c-1.6-2.3.1-3.3 2.3-2.9-2.1-2.5-.8-4.3 2.5-3.6 1 .2-.4 1.9-1.3 1.9 2.7 2 1.2 4.2-1.7 3.7 2.6 3.5-1.8 2.6-3.8 2.8-.5 2.6 2.5 5.6 1.5 5.6-2.2 0-5.8-8.3.5-7.5" fill="#42ade2"/>
    <path d="m44.5 19.3c-1.5.7-5.7-5.9-.5-6-3-2.7-2.6-4 1.4-4.1-4.6-4.6 2.7-6.2 3.4-3.8.2.7-2.2-.6-3 .7-.9 1.5 5.6 5.4-1.1 5.1 2.5 2.5 2.6 3.7-1.3 4.1.5.8 2.1 3.6 1.1 4" fill="#ff8736"/>
    <path d="m46.2 34.9l1.5-1.3c0 0 1.4 2.1 2.4 2.9.8-3.6.6-5.7 4.7-3.3-2.3-6.2 1.5-3.9 5.2-2.2-.2-1.6 0-1.4 1.6-1.9 1.4 5.3-2.4 3.7-5.4 2 1.8 4.8-.1 4.5-3.9 2.9-.1 2-.7 4.3-1.9 4.5-1.4.4-4.2-3.6-4.2-3.6" fill="#ed4c5c"/>
    <path d="m35 20.1c-1.8 2.4-4.7 3.7-6.8 5.8-2.2 2.2-3.5 8.2-3.5 8.2s.8-6.3 2.9-8.7c1.9-2.2 4.7-3.8 6.2-6.3 2.6-4.6.2-10.6-3.2-14.1.7-.6 1.7-1.4 2.2-2 3.3 4.1 6.1 12 2.2 17.1" fill="#c28fef"/>
    <path d="m38.1 25.2c-2.6 1.9-4.5 4.7-6.3 7.3-1.6 2.3-6.7 5.2-6.7 5.2s4.8-3.3 6.3-5.7c1.8-3 3.6-6.1 6.4-8.3 5.6-4.3 13.7-3.9 20-1.6-.4.9-1.1 2.8-1.1 2.8s-13.3-3.6-18.6.3" fill="#ff8736"/>
    <g fill="#42ade2">
    <path d="m49.2 24.7c-1.7 2.2-2.5 4.9-3.8 7.4-1.2 2.3-2.8 4.5-5.1 5.7-2.6 1.3-8.3.9-8.3.9s5.7-.1 8.1-1.7c2.4-1.6 3.7-4.4 4.6-7 1.8-5 4-10.4 9.2-12.6.3.9 1 2.8 1 2.8s-2.9.8-5.7 4.5"/>
    <path transform="matrix(.707-.7072.7072.707-8.3165 8.458)" d="m4 12.3h4v4h-4z"/>
    </g>
    <path transform="matrix(.7071-.7071.7071.7071-13.4747 13.8633)" fill="#ff8736" d="m8 21.2h4v4h-4z"/>
    <path transform="matrix(.707-.7072.7072.707-1.905 15.0572)" fill="#ed4c5c" d="m15.2 7.8h4v4h-4z"/>
    <path transform="matrix(.7071-.7071.7071.7071-16.8081 46.7362)" fill="#c28fef" d="m46 41.7h4v4h-4z"/>
    <path transform="matrix(.7071-.7071.7071.7071-25.5139 45.1176)" fill="#ed4c5c" d="m39.7 51.4h4v4h-4z"/>
    <path transform="matrix(.7071-.7071.7071.7071-23.4619 54.546)" fill="#ff8736" d="m52.1 53.6h4v4h-4z"/>
    <g fill="#42ade2">
    <path transform="matrix(.7071-.7071.7071.7071-13.5212 52.7722)" d="m54.9 40.7h4v4h-4z"/>
    <path transform="matrix(.7071-.7071.7071.7071 6.223 40.6826)" d="m50.2 10.8h4v4h-4z"/>
    </g>
    <path transform="matrix(.7071-.7071.7071.7071-14.6842 24.2063)" fill="#ed4c5c" d="m19.9 27.8h4v4h-4z"/>
    </svg>
</div>
</body>
</html>
`)

func main() {
	var addr = ":3000"

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	http.HandleFunc("/", homepage)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })

	slog.Info("Listening", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("ListenAndServe", "error", err)
		os.Exit(1)
	}
}

func homepage(resp http.ResponseWriter, req *http.Request) {
	slog.Info("HTTP Request", "method", req.Method, "path", req.URL.Path)
	_, err := resp.Write(page)
	if err != nil {
		slog.Error("Error writing response", "error", err)
	}
}
