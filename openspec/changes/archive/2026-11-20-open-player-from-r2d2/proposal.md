## Why

The operator opens terminal screens from r2d2. musica is a separate binary with its own popup, so the player sits outside that menu. Folding the player into r2d2 as a screen breaks downloads if the worker keeps re-running the current executable, because that executable would then be r2d2.

## What Changes

- A public package `github.com/ryangerardwilson/musica/player` exposes `Run(args, stdin, stdout, stderr) int`, which runs the existing player command. r2d2 opens it as `r2d2 musica`.
- The download worker execs the `musica` binary found on PATH with `download-worker`. If `musica` is not on PATH, spawning fails with that error. It does not fall back to the current executable.
- Bare `musica` with no arguments execs `r2d2 musica`. Every other musica command stays in the musica binary.
- Help says the player opens with `r2d2 musica`.

## Capabilities

### New Capabilities

### Modified Capabilities

- `musica`: the player opens from r2d2, and downloads stay on the musica binary.

## Impact

- Spec: `openspec/specs/musica/spec.md`.
- Code: `player/`, `cmd/musica/`, `internal/cli/`, `internal/download/`.
- Non-goals: the player look, colors, and keys do not change. The player implementation stays in this module.
