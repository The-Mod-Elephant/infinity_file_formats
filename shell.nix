let
  unstableTarball =
    fetchTarball {
      url = "https://github.com/NixOS/nixpkgs/archive/2cd3cac16691a933e94276f0a810453f17775c28.tar.gz";
      sha256 = "06ad2257srqw8q4504hqm5hyb50vlckhgdfcjx81hxq3l1wy9i5j";
    };
  pkgs = import unstableTarball {};
  shell = pkgs.mkShell {
    # Libs
    buildInputs = [
      pkgs.go
    ];
    # Tools
    nativeBuildInputs = [
      pkgs.delve
      pkgs.git
      pkgs.gopls
      pkgs.go-tools
      pkgs.pre-commit
    ];
    shellHook = ''
    # Make sure Go always has a valid temp dir
    export TMPDIR=$(mktemp -d)
    pre-commit install -f
    '';
  };
in shell
