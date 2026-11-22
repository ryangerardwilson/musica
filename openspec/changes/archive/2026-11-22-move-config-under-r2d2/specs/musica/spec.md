## MODIFIED Requirements

### Requirement: musica is the player command

The module SHALL build one command, `musica`. It MUST play a track, open a directory for one run, edit config, and end the running session. The saved library MUST open through `r2d2 musica`. Bare `musica` with no arguments MUST exec `r2d2` with the argument `musica`, found on PATH. `musica end`, `musica download-worker`, `musica play`, `musica open`, `musica config`, `musica help`, `musica version`, and `musica upgrade` MUST stay in the musica binary. Help MUST say the player opens with `r2d2 musica`. Config MUST live at `$XDG_CONFIG_HOME/r2d2/musica/config.json`, or `~/.config/r2d2/musica/config.json` when `XDG_CONFIG_HOME` is unset, and musica MUST write only that path. When that file is absent, musica MUST copy `<config>/musica/config.json` once, or else `<config>/rlc/config.json` once, before it reads or writes config, and MUST leave the old file in place.

#### Scenario: Launch the saved library

- **GIVEN** a saved `music_dir`
- **WHEN** the operator runs `r2d2 musica`
- **THEN** the TUI MUST open that library
- **AND** playback MUST continue when the overlay is hidden

#### Scenario: Bare musica opens r2d2

- **WHEN** the operator runs `musica` with no arguments
- **THEN** musica MUST exec `r2d2 musica`
- **AND** the environment MUST pass through unchanged

#### Scenario: Adopt the musica config

- **GIVEN** `~/.config/musica/config.json` exists and `~/.config/r2d2/musica/config.json` does not
- **WHEN** musica loads config
- **THEN** it MUST copy the musica config to `~/.config/r2d2/musica/config.json`
- **AND** `~/.config/musica/config.json` MUST stay in place
- **AND** later saves MUST write only `~/.config/r2d2/musica/config.json`

#### Scenario: Adopt the previous config

- **GIVEN** `~/.config/rlc/config.json` exists and neither `~/.config/musica/config.json` nor `~/.config/r2d2/musica/config.json` does
- **WHEN** musica loads config
- **THEN** it MUST copy the rlc config to `~/.config/r2d2/musica/config.json`

#### Scenario: The new config wins

- **GIVEN** `~/.config/r2d2/musica/config.json` exists
- **WHEN** musica loads config
- **THEN** it MUST NOT copy an old file over it
