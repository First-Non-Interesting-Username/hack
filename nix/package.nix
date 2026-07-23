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
    rev = "2b88c96c9052495cec023da58c9d2f6349c2961b";
    hash = "sha256-sB6gwYUPMujG4KdrqWvhiphUD+ztKxeMz/J2nN/eOdg=";
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
