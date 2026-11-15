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

### Requirement: Pause opens the cassette menu

Space MUST pause or resume the current track. Pausing MUST open the cassette menu and MUST leave the download view. Resuming MUST leave the menu as it is. Space MUST NOT cancel a download. When nothing is playing, space MUST NOT open the menu.

#### Scenario: Pause while the library is showing

- **GIVEN** a track is playing and the cassette menu is closed
- **WHEN** the operator presses space
- **THEN** playback MUST pause
- **AND** the cassette menu MUST open

#### Scenario: Pause from the download view

- **GIVEN** a track is playing and the download view is open
- **WHEN** the operator presses space
- **THEN** playback MUST pause
- **AND** the cassette menu MUST open
- **AND** the download MUST keep running

#### Scenario: Resume leaves the menu

- **GIVEN** playback is paused and the cassette menu is open
- **WHEN** the operator presses space
- **THEN** playback MUST resume
- **AND** the cassette menu MUST stay open

### Requirement: The player text is white

The visualizer glyphs, status line, help, download view, and YouTube prompt MUST use white text. They MUST NOT use a separate hue for titles, status, errors, selection, playback, or download state. Visualizer rows MUST share one base color. Empty visualizer cells MAY stay dark.

#### Scenario: Text is white

- **GIVEN** the player is drawing
- **WHEN** a title, a status line, a rain glyph, a running download, and a failed download are styled
- **THEN** each foreground MUST be white
- **AND** the running download and the failed download MUST share that white

#### Scenario: Visualizer rows share one color

- **GIVEN** the visualizer is drawing
- **WHEN** the top, middle, and bottom rows are styled
- **THEN** those rows MUST use the same base color

### Requirement: q ends the session

`q` MUST stop playback and quit the TUI when no text prompt is open, including from the library, the cassette menu, help, and the download view. `q` MUST NOT cancel a background download. While a text prompt is open, `q` MUST be typed and MUST NOT quit.

#### Scenario: q closes a playing session

- **GIVEN** a track is playing and no text prompt is open
- **WHEN** the operator presses `q`
- **THEN** playback MUST stop
- **AND** the TUI MUST quit

#### Scenario: q from the download view leaves the download running

- **GIVEN** a track is playing and the download view shows a running download
- **WHEN** the operator presses `q`
- **THEN** the TUI MUST quit
- **AND** that download MUST keep running

#### Scenario: q types into a prompt

- **GIVEN** a text prompt is open
- **WHEN** the operator presses `q`
- **THEN** the prompt MUST contain `q`
- **AND** the TUI MUST stay open

