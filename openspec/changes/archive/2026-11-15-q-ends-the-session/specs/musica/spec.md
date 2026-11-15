## ADDED Requirements

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
