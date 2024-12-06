.PHONY: all demo
all: slides.html slides.pdf demo

slides.html: slides.md
	docker run --rm --entrypoint=/bin/bash --volume="${CURDIR}:/src/" --workdir=/src  marpteam/marp-cli:v4.0.3 -c 'marp-cli.js --allow-local-files /src/slides.md'

slides.pdf: slides.md
	docker run --rm --entrypoint=/bin/bash --volume="${CURDIR}:/src/" --workdir=/src  marpteam/marp-cli:v4.0.3 -c 'marp-cli.js --allow-local-files --pdf /src/slides.md'

# Set or reset the docker image registry and domain name so we don't commit them to the repo
IMAGE_REGISTRY := example
EXAMPLE_DOMAIN := example.com
demo:
	sed -E -i 's#image: .+$$#image: $(IMAGE_REGISTRY)#' example/manifest.yaml
	sed -E -i 's#Host\(`dev-ncapo\..+`\)#Host(`dev-ncapo\.$(EXAMPLE_DOMAIN)`)#' example/manifest.yaml
	sed -E -i 's#ref=".+",#ref="$(IMAGE_REGISTRY)",#' Tiltfile
