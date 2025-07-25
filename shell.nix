with import <nixpkgs> {};
pkgs.mkShell {
  buildInputs = [
    delve
    git
    go
    pre-commit
  ];
}
