## MODIFIED Requirements

### Requirement: musica is the player command

The module SHALL build one command, `musica`. It MUST play a track, open a directory for one run, edit config, and end the running session. The saved library MUST open through `r2d2 musica`. Bare `musica` with no arguments MUST exec `r2d2` with the argument `musica`, found on PATH. `musica end`, `musica download-worker`, `musica play`, `musica open`, `musica config`, `musica help`, `musica version`, and `musica upgrade` MUST stay in the musica binary. Help MUST say the player opens with `r2d2 musica`. An existing `~/.config/rlc/config.json` MUST be copied once when the musica config is absent.

#### Scenario: Launch the saved library

- **GIVEN** a saved `music_dir`
- **WHEN** the operator runs `r2d2 musica`
- **THEN** the TUI MUST open that library
- **AND** playback MUST continue when the overlay is hidden

#### Scenario: Bare musica opens r2d2

- **WHEN** the operator runs `musica` with no arguments
- **THEN** musica MUST exec `r2d2 musica`
- **AND** the environment MUST pass through unchanged

#### Scenario: Adopt the previous config

- **GIVEN** `~/.config/rlc/config.json` exists and `~/.config/musica/config.json` does not
- **WHEN** musica loads config
- **THEN** it MUST copy the rlc config to the musica path

## ADDED Requirements

### Requirement: r2d2 opens the player

The module MUST provide the public package `github.com/ryangerardwilson/musica/player` with `Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int`. `Run` MUST run the same player command as the musica binary, with no arguments opening the saved library, and MUST NOT exec `r2d2`. The player implementation MUST stay in this module.

#### Scenario: r2d2 opens the library

- **WHEN** r2d2 calls `player.Run` with no arguments
- **THEN** the player MUST open the saved library in that process

#### Scenario: r2d2 ends the session

- **WHEN** r2d2 calls `player.Run` with `end`
- **THEN** the running musica session MUST end

### Requirement: The download worker runs on the musica binary

The player MUST start the download worker by execing `musica` found on PATH with the argument `download-worker`, in a new session. It MUST NOT exec the current executable. When `musica` is not on PATH, starting the worker MUST fail with that error. `musica download-worker` MUST drain the queue.

#### Scenario: The player runs inside r2d2

- **GIVEN** the player runs inside r2d2
- **WHEN** a download is queued
- **THEN** the worker MUST be `musica download-worker`
- **AND** it MUST NOT be `r2d2 download-worker`

#### Scenario: musica is missing

- **GIVEN** `musica` is not on PATH
- **WHEN** the player starts the worker
- **THEN** starting MUST fail with the lookup error
