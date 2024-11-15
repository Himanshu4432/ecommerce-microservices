// internal/handlers/handlers.go\npackage handlers\n\nimport (\n\t"net/http"\n)\n\nfunc SampleHandler(w http.ResponseWriter, r *http.Request) {\n\tw.Write([]byte("Hello from payment-service!"))\n}\n
