# musica Specification

## Purpose

Define the musica player: local playback, nebula s3 backup of its sqlite state, and YouTube downloads that continue outside the TUI.

## Requirements

### Requirement: musica is the player command

The module SHALL build one command, `musica`. It MUST open the saved music library, play a track, open a directory for one run, edit config, and end the running session. An existing `~/.config/rlc/config.json` MUST be copied once when the musica config is absent.

#### Scenario: Launch the saved library

- **GIVEN** a saved `music_dir`
- **WHEN** the operator runs `musica`
- **THEN** the TUI MUST open that library
- **AND** playback MUST continue when the overlay is hidden

#### Scenario: Adopt the previous config

- **GIVEN** `~/.config/rlc/config.json` exists and `~/.config/musica/config.json` does not
- **WHEN** musica loads config
- **THEN** it MUST copy the rlc config to the musica path

### Requirement: nebula s3 backs up application state

musica MUST keep downloads and the saved music directory in its own sqlite file. While the machine is up and `SQLX_ROOT` is unset, that file MUST be `$XDG_RUNTIME_DIR/musica/musica.sqlite`. nebula `s3` MUST store the backup at `s3://backups/musica.sqlite`. The application MUST NOT call nebula sql. A failed s3 call MUST NOT roll back the local sqlite write. Audio files MUST remain in the configured music directory.

#### Scenario: Restore a missing database

- **GIVEN** the runtime sqlite file is missing and `s3://backups/musica.sqlite` exists
- **WHEN** musica opens its database
- **THEN** it MUST restore that object before creating a new file

#### Scenario: Publish after a download changes

- **GIVEN** a download is queued, finished, failed, or cancelled
- **WHEN** musica records that state
- **THEN** it MUST put the sqlite file at `s3://backups/musica.sqlite`

### Requirement: YouTube downloads continue outside the TUI

From an open cassette, `y` MUST ask for a YouTube URL and then a song name. The download MUST be queued and shown in a download view whose running progress animation changes over time. The download process MUST keep running after the TUI exits and while another track is playing. Cancel from the download view MUST stop that download and MUST NOT stop other downloads or playback. `u` MUST open the download view.

#### Scenario: Queue a download

- **GIVEN** a cassette is open
- **WHEN** the operator submits a YouTube URL and a song name
- **THEN** musica MUST queue the download and show the download view
- **AND** the operator MUST be able to return to the library and play another track while it runs

#### Scenario: Close the app during a download

- **GIVEN** a download is running
- **WHEN** the operator closes the TUI
- **THEN** the download worker MUST keep running
- **AND** opening musica again MUST show that download's progress

#### Scenario: Cancel from the download view

- **GIVEN** the download view has a selected running download and music is playing
- **WHEN** the operator cancels that download
- **THEN** that download MUST stop
- **AND** playback MUST continue
