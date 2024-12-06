.PHONY: all
all: slides.html slides.pdf

slides.html: slides.md
	docker run --rm --entrypoint=/bin/bash --volume="${CURDIR}:/src/" --workdir=/src  marpteam/marp-cli:v4.0.3 -c 'marp-cli.js --allow-local-files /src/slides.md'

slides.pdf: slides.md
	docker run --rm --entrypoint=/bin/bash --volume="${CURDIR}:/src/" --workdir=/src  marpteam/marp-cli:v4.0.3 -c 'marp-cli.js --allow-local-files --pdf /src/slides.md'

# Set or reset the docker image registry so we don't commit the registry URL to the repo
IMAGE_REGISTRY ?= example
demo:
	sed -i 's#image: example#image: $(IMAGE_REGISTRY)#' example/manifest.yaml
	sed -i 's#ref="example",#ref="$(IMAGE_REGISTRY)",#' Tiltfile

reset:
	sed -i 's#image: example#image: example#' example/manifest.yaml
	sed -i 's#ref="example",#ref="example",#' Tiltfile
