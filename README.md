# hack

Dead-simple, UNIX style tool for interacting with LLMs on command line.

![Hackatime Badge](https://hackatime-badge.hackclub.com/@janekmusin/hack-ai-cli)

## Installation and usage

### Installation

```bash
# Clone this repo
gh repo clone First-Non-Interesting-Username/hack && cd hack
# Build the binary
go build
```

Alternatively, get the binary for your system from [github releases](https://github.com/First-Non-Interesting-Username/hack/releases)

There's also a Nix package available. Add:

```nix
hack.url = "github:First-Non-Interesting-Username/hack";
```

to your flake inputs.

Then you could either that package to your package list (either home manager or system):

```nix
inputs.hack.packages.${pkgs.stdenv.hostPlatform.system}.hack
```

or use the module:

```nix
# Add that to your nixos imports
inputs.hack.nixosModules.default
# Or this to your home-manager imports
inputs.hack.homeManagerModules.default
# And this to your nixos/home manager configuration
# The syntax is the same for both
programs.hack = {
  enable = true;
  settings = {
    # Equivalent to the config file
  };
};
```

### Configuration

Create a file named `config.toml` either in `$XDG_CONFIG_HOME/hack` or `/etc/hack`.
The user dir takes precedence over system dir.

File contents:

```toml
base_url = # Base URL (without /chat/completions)
model = # Full name of the model in your provider
api_key = # Your API key
api_key_path = # Path to a file containing your API key. It must be readable by the user you will be running hack with
```

### Usage

```bash
hack --help

hack is a command-line tool for AI-assisted development.

It sends prompts to OpenAI-compatible LLM APIs and returns output tailored
to the task at hand: shell commands, executable code, documentation, or
plain responses.

Usage:
        hack -p "your prompt here"
        echo "some content" | hack -p "summarize this"
        ls -la | hack -x "delete the largest file"

Modes:
        shell   (-x/--execute)  Generate shell commands from a prompt
        code    (-w/--write)    Output executable code (jq, python3, bash, or POSIX sh)
        normal                                  Standard prompt-and-response

Usage:
  hack [flags]

Flags:
  -b, --base string       base URL for the API (without /chat/completions)
  -c, --config string     config file path, $HOME/.config/hack-ai/config.toml if not provided
  -h, --help              help for hack
  -k, --key string        API key for selected provider
      --key-file string   Path to file containing API key
  -m, --model string      LLM used for response
  -p, --prompt string     prompt for the LLM
  -s, --shell             enable command mode
  -v, --version           version for hack
  -w, --write             enable code mode
```

## Roadmap

- Exa integration
