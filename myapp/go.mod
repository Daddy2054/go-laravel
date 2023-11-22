module myapp

go 1.21.4

replace github.com/daddy2054/celeritas => ../celeritas

require github.com/daddy2054/celeritas v0.0.0-00010101000000-000000000000

require (
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
