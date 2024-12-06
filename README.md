# In-Cluster Development Environments using Tilt

Author: Nicholas Capo
Presented: 2024-12-05

## Contents

`slides.md`: Slide contents written in Markdown to be converted to HTML/PDF using [marp](https://marp.app/)
`slides.html`/`slides.pdf`: Generated slides
`Makefile`: Commands to generated slides

`example/`: Example application to be run during the demo

## Demo Preparation

Docker registry URL's or other domain names have not bee committed to the repo.
To insert them in the right place run this command with your values:

```shell
make demo IMAGE_REGISTRY=example.com/registry EXAMPLE_DOMAIN=example.com
```

Running the command without extra arguments resets the default values

```shell
make demo
```
