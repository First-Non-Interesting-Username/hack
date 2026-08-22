The project is developed using devenv.

To run commands inside devenv, run `devenv shell <command>`.

If go tooling isn't available on the system, use above command, do not install it.

When writing Go code, avoid complexity and prefer readable code.

If not adding new, very big features, all go files should be under `cmd/` directory, in `cmd` package.

When possible, prefer using cobra/viper options instead of basic go functions.

Compilation:

```bash
go build
```

Running:

```bash
go run
```

The key in `config.toml.example` is revoked.

For docs, use simple language.
