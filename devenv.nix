{pkgs, ...}: {
  languages.go.enable = true;

  packages = [pkgs.cobra-cli];

  git-hooks = {
    gofmt.enable = true;
    govet.enable = true;
  };
}
