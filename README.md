# fetch-host

Optimal remote host fetching for UTCS students

## Prerequisites

[golang](https://go.dev/doc/install)

## Install

Clone this repo:

```
git clone https://github.com/migopp/fetch-host.git
```

Then build:

```
cd fetch-host/cmd/fetch-host/ && go install
```

This should install the binary to your `~/go/bin`. Add this directory to your `PATH` if it's not already there.

```
export PATH=$PATH:~/go/bin
```

## Configuration

You can config the resulting SSH command in `~/.config/fetch-host/config.json`.

A base configuration might be:

```json
{
  "utcsUsername": "migopp",
  "sshTemplate": "ssh %s@%s.cs.utexas.edu"
}
```

`utcsUsername` is your username. `sshTemplate` is a `go` format string. For exmaple, if you wanted SSH with x-forwarding enabled, you would change the template to `ssh -X %s@%s.cs.utexas.edu`.

If a config file doesn't exist when `fetch-host` executes, then it will generate one.

## Usage

```
fetch-host
```

The resulting command with the optimal host will be printed to the console.

### Instant SSH

You can set an alias in your shell profile for a fast ssh, if desired.

```
alias fssh=$(fetch-host)
```
