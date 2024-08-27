{inputs, ...}: {
  imports = [
    inputs.devenv.flakeModule
  ];
  perSystem = {pkgs, ...}: {
    devenv.shells.default = {
      enterShell = ''
        printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n\n"

        help-me
      '';
      languages = {
        go.enable = true;
        javascript.enable = true;
        nix.enable = true;
        typescript.enable = true;
      };
      packages = [
        pkgs.commitizen
      ];
      pre-commit = {
        default_stages = ["pre-push"];
        hooks = {
          actionlint.enable = true;
          alejandra.enable = true;
          check-added-large-files = {
            enable = true;
            stages = ["pre-commit"];
          };
          check-yaml.enable = true;
          commitizen.enable = true;
          deadnix.enable = true;
          detect-private-keys = {
            enable = true;
            stages = ["pre-commit"];
          };
          end-of-file-fixer.enable = true;
          flake-checker.enable = true;
          markdownlint.enable = true;
          mixed-line-endings.enable = true;
          nil.enable = true;
          no-commit-to-branch = {
            enable = true;
            stages = ["pre-commit"];
          };
          ripsecrets = {
            enable = true;
            stages = ["pre-commit"];
          };
          shellcheck.enable = true;
          shfmt.enable = true;
          statix.enable = true;
        };
      };

      scripts = {
        "help-me".exec = ''
          printf "Welcome to Generate onboarding!\n\n" | ${pkgs.lolcat}/bin/lolcat

          printf "The commands below are available to help you during onboarding:\n"

          printf "\nRun 'help-me' to see this message again\n"
        '';
      };
    };
  };
}
