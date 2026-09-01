# SPDX-FileCopyrightText: 2026 First-Non-Interesting-Username
#
# SPDX-License-Identifier: GPL-3.0-only
{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:
buildGoModule (finalAttrs: {
  pname = "hack";
  version = "4";
  structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "First-Non-Interesting-Username";
    repo = "hack";
    tag = finalAttrs.version;
    hash = "sha256-NEoYmLkL7EN1gmI5G1Rf87avyz5o2D7AOcP5N+Uhpto=";
  };

  vendorHash = "sha256-gTj1xJwj/Qf+v6wY5FEheWxFbHQ4NEV68FflgGbqnDc=";

  ldflags = ["-s"];

  env = {
     CGO_ENABLED = 0;
  };

  passthru.updateScript = nix-update-script {
    extraArgs = [ "--flake" ];
  };

  meta = {
    description = "CLI tool for interacting with LLMs";
    homepage = "https://github.com/First-Non-Interesting-Username/hack";
    license = lib.licenses.gpl3Only;
    # maintainers = with lib.maintainers; [ ];
    mainProgram = "hack";
  };
})
