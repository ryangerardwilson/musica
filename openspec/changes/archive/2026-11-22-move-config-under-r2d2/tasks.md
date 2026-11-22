## 1. Spec

- [x] 1.1 Validate this change strict.

## 2. Config

- [x] 2.1 Move the default path to `<config>/r2d2/musica/config.json`.
- [x] 2.2 Copy `<config>/musica/config.json`, or else `<config>/rlc/config.json`, once when the new file is absent.
- [x] 2.3 Adopt the old file before `musica config` and before every player run.
- [x] 2.4 Name the new path in README and AGENTS.md.
- [x] 2.5 Run `GOWORK=off go test ./...`.
