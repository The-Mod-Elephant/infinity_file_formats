{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.git
    pkgs.go
    pkgs.pre-commit
  ];
}
