## Why

musica is an r2d2 screen now, and r2d2 keeps each feature's config under `~/.config/r2d2/<feature>/`. musica still writes `~/.config/musica/config.json`, the directory it had as a separate program.

## What Changes

- Config moves to `$XDG_CONFIG_HOME/r2d2/musica/config.json`, or `~/.config/r2d2/musica/config.json`.
- musica writes only the new path.
- When the new file is absent, musica copies `~/.config/musica/config.json` once, or else `~/.config/rlc/config.json` once. The old file stays in place.
- `musica config` and every player run adopt the old file before they read or write.

## Capabilities

### Modified Capabilities

- `musica`: the config path moves under `r2d2/musica`.

## Impact

- `internal/config` path and legacy copy.
- `internal/cli` calls the copy before every config read or write.
- README and AGENTS.md name the new path.
