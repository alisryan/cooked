{inputs, ...}: {
  imports = [
    inputs.devenv.flakeModule
  ];
  perSystem = {pkgs, ...}: {
    devenv = {
      modules = [
        inputs.env-help.devenvModule
      ];
      shells.default = {
        enterShell = ''
          printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n\n"

          printf "Welcome to Generate onboarding!\n\n" | ${pkgs.lolcat}/bin/lolcat
          printf "The commands below are available to help you during onboarding:\n"

          env-help

          printf "\nRun 'env-help' to see this message again\n"
        '';
        env-help.enable = true;
        languages = {
          go.enable = true;
          javascript = {
            enable = true;
            pnpm.enable = true;
          };
          nix.enable = true;
          typescript.enable = true;
        };
        packages = with pkgs; [
          commitizen
          golangci-lint
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
          "backend-dev" = {
            description = "Runs the backend server in development mode.";
            exec = ''
              cd $DEVENV_ROOT/backend
              go run cmd/server/main.go
            '';
          };
          "backend-lint" = {
            description = "Lints backend code.";
            exec = ''
              cd $DEVENV_ROOT/backend
              golangci-lint run ./...
            '';
          };
          "frontend-dev" = {
            description = "Runs the frontend server in development mode.";
            exec = ''
              cd $DEVENV_ROOT/frontend
              pnpm install
              pnpm exec next dev
            '';
          };
          "frontend-lint" = {
            description = "Lints frontend code.";
            exec = ''
              cd $DEVENV_ROOT/frontend
              pnpm exec next lint
            '';
          };
        };
      };
    };
  };
}
