{
  description = "jason-cli development environment";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/master";
    go22.url = "github:NixOS/nixpkgs/9a9dae8f6319600fa9aebde37f340975cab4b8c0";
  };
  outputs = { self, nixpkgs, ... }@inputs:
    let 
      system = "x86_64-linux";
      gopkg = import inputs.go22 {inherit system; };
      jasonCli = gopkg.buildGo122Module {
        name = "jason-cli";
        src = ./.;
        vendorHash = "sha256-NGXcrO4BhLm2bSkiGc9ueOzieqvtoLGsksBcX0oEZ8g=";
      };
  in 
  { 
    devShells.${system}.default = 
      gopkg.mkShell {
        name = "jason-cli dev";
        nativeBuildInputs = with gopkg; [
          go
        ];
        packages = [ jasonCli ];
      };
    packages.${system}.default = jasonCli;
  };
}
