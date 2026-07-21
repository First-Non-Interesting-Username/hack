{ pkgs, ... }:

{
  languages.go.enable = true;

  packages = [ pkgs.cobra-cli ];
}
