{
  description = "restart-and-tail";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };
  outputs =
    { nixpkgs, ... }:
    let

      forAllSystems =
        function:
        nixpkgs.lib.genAttrs [
          "x86_64-linux"
          "aarch64-linux"
          "aarch64-darwin"
        ] (system: function nixpkgs.legacyPackages.${system});
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = [
            pkgs.go
            pkgs.gopls
            pkgs.devenv
            pkgs.nixfmt-rfc-style
            pkgs.golangci-lint
          ];
        };
      });
      packages = forAllSystems (pkgs: {
        default = pkgs.buildGoModule {
          pname = "restart-and-tail";
          version = "1.0.0";
          src = ./.;
          vendorHash = "sha256-0YQ1WlPnvjRwg7nWJkOh3IdHhowBGbutbQOTPgBfxO4=";
        };
      });
    };
}
