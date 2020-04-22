### Quickstart

> Requires docker and docker-compose install locally

```bash

./docker-helper.sh --up

# You should now have a docker shell...

# Init install/cache dependencies and install tools to bin
make init

# Building (generate, format, build, vet)
make

# Execute tests
make test

# Migrate up your local database
sql-migrate up

# Start the local built server
apiserver

```

### vscode

Same requirements as above, always develop *inside* the running `development` docker container. 

Run CMD+SHIFT+P `Go: Install/Update Tools` after starting vscode to autoinstall all golang vscode dependencies, then **reload your window**.