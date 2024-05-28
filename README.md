# gohx

## Requirements

To develop this stack you need to install below packages:

- [golang](https://go.dev/dl/) >= 1.22
- [templ](https://templ.guide/quick-start/installation) >= v0.2
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) >= v1.26
- [bun](https://bun.sh/docs/installation) >= 1.1

For development purpose, these packages will help you to boost your local development process:

- [air](https://github.com/cosmtrek/air?tab=readme-ov-file#installation) >= v1.51
- [taskfile](https://taskfile.dev/installation/) >= v3.35

## Getting started

Once you clone the repository, you need to setup your project by installing dependencies by following command:

```sh
task setup
# or manually
go mod tidy
templ generate
bun install --frozen-lockfile
bun run build
```

Then you need to adjust the `.env` file for local development by following command:

```sh
cp .env.example .env
```

Once everything has been setup properly, you can run the application by:

```sh
task dev
# or simply run
go run cmd/server/*.go
```

## License

This template is [licensed](./LICENSE) under the MIT License.
