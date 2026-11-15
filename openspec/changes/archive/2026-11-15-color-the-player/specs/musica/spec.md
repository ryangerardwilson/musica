## ADDED Requirements

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

### Requirement: The player surface uses color

The visualizer, status line, help title, download view, and YouTube prompt MUST use color. Visualizer rows MUST share one base color so the field does not form horizontal bands. Rain intensity MUST use more than one color. A failed download MUST use a different color from a running download.

#### Scenario: Rain intensity is colored

- **GIVEN** the visualizer is drawing
- **WHEN** a bright rain cell and a trace rain cell are styled
- **THEN** their colors MUST differ
- **AND** every visualizer row MUST use the same base color

#### Scenario: A failed download is colored

- **GIVEN** one download has failed and another is running
- **WHEN** the download view chooses their colors
- **THEN** those colors MUST differ
