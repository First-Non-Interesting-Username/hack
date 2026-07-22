self:

{ config, lib, pkgs, ... }:

let
  cfg = config.programs.hack;
  inherit (pkgs.stdenv.hostPlatform) system;
in
{
  options.programs.hack = {
    enable = lib.mkEnableOption "hack, CLI tool for interacting with LLMs";

    package = lib.mkOption {
      type = lib.types.package;
      default = self.packages.${system}.hack
        or (throw "hack: no package available for system '${system}'");
      defaultText = lib.literalExpression "self.packages.\${pkgs.stdenv.hostPlatform.system}.hack";
      description = "The hack package to install.";
    };
    settings = {
      model = lib.mkOption {
        type = lib.types.string or null;
        default = null;
        example = "anthropic/claude-fable-5";
        description = "Model to be used if no overwrite was provided";
      };
      base_url = {
        type = lib.types.string or null;
        default = null;
        example = "https://ai.hackclub.com/proxy/v1";
        description = "API base to be used if no overwrite was provided";
      };
      api_key_path = {
        type = lib.types.path or null;
        default = null;
        example = "\${config.sops.secrets.HACK_CLUB_AI_API_KEY.path}";
        description = "Path to file containing the API key";
      };
    };
  };
}
