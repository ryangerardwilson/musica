# Product Engineer Role

## Purpose

Own musica-specific facts that should not live in root generalists.

## Load Guidance

Load this file for `musica` implementation, CLI/TUI, installer, release, storage, configuration, or project-specific product work.

Root generalists own role behavior. This file owns only project facts and
repo-local operating constraints.

## Owns

- repo-local product and implementation facts
- CLI/TUI contract, command grammar, config, storage, and installer constraints
- release, upgrade, and verification expectations specific to this app

## Project Context

## Scope
- `musica` is a terminal-native local music player for a saved music library.
- Preserve the narrow loop: launch the saved library, open a directory for one run, play one track directly, or edit the user config.
- Do not grow this into a streaming service, library daemon, or web UI unless the user explicitly asks.
- Download jobs and the saved library path live in sqlite. nebula `keep` backs up `s3://backups/musica.sqlite`. Audio files stay in the music directory.
- `y` queues a YouTube audio download. The worker keeps running after the TUI closes or other tracks play. Cancel only from the download view.

## Command Surface
- `musica` launches the saved library and may bootstrap the first-run config.
- `musica play <track>` launches with one track selected for immediate playback.
- `musica open <music_dir>` launches against a directory for that run without changing saved config.
- `musica config` opens the real user config in `$VISUAL`, then `$EDITOR`, then `vim`.
- Do not restore removed aliases, bare target paths, or shorthand config/render flags.

## Implementation Rules
- Keep help, version, upgrade, config opening, and dispatch app-local in `internal/cli/`.
- Do not restore the deleted shared CLI runtime dependency or contract file.
- Keep `help`, `version`, and `upgrade` fast and free of TUI/player initialization.
- The terminal surface is Go/Bubble Tea only.
- Keep the TUI keyboard-first, low-chrome, and local-file oriented.
