def main():
    secret_settings(disable_scrub=True)
    allow_k8s_contexts("dev")
    example()

def example():
    docker_build(
        context="example",
        dockerfile="example/Dockerfile",
        ignore=["manifest.yaml"],
        target="development",
        ref="example",
        platform="linux/amd64",
        live_update=[
            fall_back_on("example/air.toml"),
            sync("example/go.mod", "/go/src/example/"),
            sync("example/go.sum", "/go/src/example/"),
            sync("example/main.go", "/go/src/example/"),
            # run("kill -SIGHUP 1"),
        ],
    )

    k8s_yaml("example/manifest.yaml")

    k8s_resource(
        'example',
        labels=["services"],
        port_forwards=[port_forward(40002, 40000, "debugger")],
        links=[link("https://dev-ncapo.example.com/")],
    )

main()
