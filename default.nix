with import <nixpkgs> {};
stdenv.mkDerivation rec {
  env = buildEnv { name = name; paths = buildInputs; };
  name = "env";
  buildInputs = [
    go
    gnumake
  ];
}
