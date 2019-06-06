PACKAGES := ./service/...

.PHONY: run generate client route openapi

run:
	gen run ${PACKAGES}

generate: client route openapi

client:
	mkdir -p client && gen client -o ./client/client_gen.go ${PACKAGES}

route:
	mkdir -p route && gen route -o ./route/route_gen.go ${PACKAGES}

openapi:
	mkdir -p openapi && gen openapi -o ./openapi/openapi.json ${PACKAGES}

