let
  unstableTarball =
    fetchTarball
      https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz;
  pkgs = import <nixpkgs> {};
  unstable = import unstableTarball {};
  shell = pkgs.mkShell {
    buildInputs = [
      unstable.delve
      pkgs.git
      unstable.go
      unstable.gopls
      unstable.go-tools
      pkgs.pre-commit
    ];
  };
in shell
