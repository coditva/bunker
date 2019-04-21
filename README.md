[![Build Status](https://travis-ci.com/coditva/bunker.svg?token=6cA7nqyyepjxqz4zK2sH&branch=master)](https://travis-ci.com/coditva/bunker)

# Bunker
_Proof-of-concept containerization engine_


## About
The project is divided into two binaries:

### `bunkerd`
This is the daemon which runs in the background with
[containerd](https://github.com/containerd/containerd) as it's back-end. It
exposes an API accessible via RPC over the socket mentioned in the
`internal/config.go` file. The API could be found in the `internal/api/` folder.
It is similar to the Docker API but is more minimal.

### `bunker`
This is the CLI for the bunker daemon. It parses user commands and sends them to
the bunkerd daemon over RPC.


## Building the project
To build both the binaries, simply issue the `make` command in the project root.
```bash
make    # build the binaries
```

To build the binaries selectively, you can specify the binary name as the target
to the make file.
```bash
make bunker     # build bunker cli
make bunkerd    # build bunker daemon
```

You can even use the build script in the `scripts/build` directory.
```bash
export TARGET=bunker
sh scripts/build/binary.sh
```


## Using Bunker
### Using `bunkerd` daemon
To start the bunkerd daemon, simply issue the following command (you will need
`sudo` rights):
```bash
sudo build/bunkerd start    # start the daemon
sudo build/bunkerd stop     # stop the daemon
```

### Using the `bunker` CLI
Use the binary in `build/` to run the CLI:
```bash
sudo build/bunker pull [image]      # pull an image
```


## API
The Bunker CLI is very similar to the most popular containerization engine
Docker, but is more minimal than that. The project is in its early stages and
at present the following commands are implemented in `bunker`:
- `pull`: Pull an image from the docker registry
- `images`: List images pulled from the registry


## Debugging
The log files for the daemon and CLI found in `/tmp` as `/tmp/bunkerd.log` and
`/tmp/bunker.log` respectively can be used to debug errors or track execution.


## License
MIT (c) 2019-present Utkarsh Maheshwari <github.com/coditva>
