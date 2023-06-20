gen:
	rm -rf ./api
	mkdir api
	mkdir api/chi
	oapi-codegen -package api -generate "types,server" -o api/api.gen.go openapi.yaml
	oapi-codegen -package chi_api -generate "types,client" -o api/chi/api.gen.go chi_openapi.yaml
