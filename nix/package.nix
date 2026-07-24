{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:
buildGoModule (finalAttrs: {
  pname = "hack";
  version = "2-unstable-2026-07-24";
  structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "First-Non-Interesting-Username";
    repo = "hack";
    rev = "c82a4cb7c52a7f86b2617deaf33c077164f2868f";
    hash = "sha256-OPF4BLcJ6gW6lTsBtpwk7pgboADp5Tv+KX4URJO0M58=";
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
