# NixOS module: programs.hack.enable installs hack system-wide.
self:

{ config, lib, ... }:

let
  cfg = config.programs.hack;
in
{
  imports = [ (import ./module.nix self) ];

  config = lib.mkIf cfg.enable {
    environment.systemPackages = [ cfg.package ];
  };
}
