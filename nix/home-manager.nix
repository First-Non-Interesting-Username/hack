self:

{ config, lib, ... }:

let
  cfg = config.programs.hack;
in
{
  imports = [ (import ./module.nix self) ];

  config = lib.mkIf cfg.enable {
    home.packages = [ cfg.package ];
  };
}
