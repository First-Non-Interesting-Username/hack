{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:

buildGoModule (finalAttrs: {
  pname = "hack";
  version = "0-unstable-2026-07-22";
  __structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "First-Non-Interesting-Username";
    repo = "hack";
    rev = "611ce15e057fa248c05524d79196d5d059c7fa17";
    hash = "sha256-vsa/iu/Qg079KtQ9N9yX5jKJrArpf0DlqlxUqtmHogg=";
  };

  vendorHash = "sha256-lo1Dz5isHCfX+mQ0zoF0jCXwnT7WXBfzwgC9UJBvtQg=";

  ldflags = [ "-s" ];

  passthru.updateScript = nix-update-script { };

  meta = {
    description = "";
    homepage = "https://github.com/First-Non-Interesting-Username/hack";
    license = lib.licenses.gpl3Only;
    # maintainers = with lib.maintainers; [ ];
    mainProgram = "hack";
  };
})
