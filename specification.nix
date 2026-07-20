{
  pkgs,
  config,
  ...
}: let
  cfg = config.projects.hack-ai-cli;
  file = cfg.configuration.file.content;
in {
  projects.hack-ai-cli = {
    format = "cli";
    language = "go";
    configuration.file = {
      format = pkgs.formats.toml;
      path = "${config.xdg.configHome}/hack-ai/config.toml";
      # Example values
      content = {
        "api-key" = "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX";
        # Don't use this model
        "model" = "openai/gpt-3.5-turbo-0613";
        "base-url" = "https://openrouter.ai/api/v1";
      };
    };
    cli = {
      command = "hack";
      io = {
        input = {
          description = "Prompt is stdin, unless overwritten by -p, in which stdin is included as context";
        };
        output = {
          description = "Model output is stdout, everything else (just errors) stderr";
        };
      };
      modes = {
        command = {
          description = ''
              Outputs command, to be piped in shell
              eg. `hack -x -p "Give me a command doing the same as ls -la, but for eza" | bash -c`
              This example is pointless, but I think it ilustrates the intent well enough
            '';
          prompt = "TBD";
        };
        normal = {
          description = ''
            The model response isn't manipulated
          '';
          prompt = "TBD";
        };
        code = {
          description = ''
            Outputs code, as specified by user.
            Intended for prototyping or writing small scripts, unusable for agentic coding.
          '';
          prompt = "TBD";
        };
      };
      flags = {
        "-p" = {
          alt = "--prompt";
          description = "Prompt, reads prompt from stdin if not provided";
        };
        "-c" = {
          alt = "--config";
          description = "${cfg.configuration.file.path} overwrite";
        };
        "-k" = {
          alt = "--key";
          description = "${file."api-key"} overwrite";
        };
        "-m" = {
          alt = "--model";
          description = "${file."model"} overwrite";
        };
        "-b" = {
          alt = "--base";
          description = "${file."base-url"} overwrite";
        };
        "-s" = {
          alt = "--shell";
          description = "Enable ${cfg.cli.modes.command}";
        };
        "-w" = {
          alt = "--write";
          description = "Enable ${cfg.cli.modes.code}";
        };
        "-v" = {
          alt = "--version";
          description = "Print current version";
        };
        "-h" = {
          alt = "--help";
          description = "Print help informantion with all the flags";
        };
      };
    };
  };
}
