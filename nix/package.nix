{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:
buildGoModule (finalAttrs: {
  pname = "hack";
  version = "1-unstable-2026-07-23";
  structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "First-Non-Interesting-Username";
    repo = "hack";
    rev = "a61328b58b06ebb3a32427adec65c5829477f376";
    hash = "sha256-aYgukxqLaULmbNRpq2Men5oFQXP4gAv1yrtKbpSeJoM=";
  };

  vendorHash = "sha256-qomzX5XGDe0XeD45styCwjS/2tGVeNpS7sqpKtxHnHk=";

  ldflags = ["-s"];

  passthru.updateScript = nix-update-script {};

  meta = {
    description = "CLI tool for interacting with LLMs";
    homepage = "https://github.com/First-Non-Interesting-Username/hack";
    license = lib.licenses.gpl3Only;
    # maintainers = with lib.maintainers; [ ];
    mainProgram = "hack";
  };
})
