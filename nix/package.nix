{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:
buildGoModule (finalAttrs: {
  pname = "hack";
  version = "0-unstable-2026-07-23";
  structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "First-Non-Interesting-Username";
    repo = "hack";
    rev = "8710c4e75b05a312e75995ea550e84b7fb7625d4";
    hash = "sha256-v9ZM2qtcfFsB5fQkYkaOcBmzb4aqKzV8N/yvgtYs5hY=";
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
