{ pkgs ? import (fetchTarball https://github.com/nixos/nixpkgs/archive/497961355b90890804f542a10791a8616e4dca75.tar.gz) {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.terraform
  ];
}
