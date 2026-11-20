## MODIFIED Requirements

### Requirement: YouTube downloads continue outside the TUI

From an open cassette, `y` MUST ask for a YouTube URL and then a song name. The download MUST be queued and shown in a download view whose running progress animation changes over time. The download process MUST keep running after the TUI exits and while another track is playing. Cancel from the download view MUST stop that download and MUST NOT stop other downloads or playback. `u` MUST open the download view. A download that has finished MUST leave the download view. A queued, running, failed, or cancelled download MUST stay visible. The finished audio MUST remain in the cassette. `j` and `k` MUST wrap the rows in that view. On the last row, `j` MUST move to the first row. On the first row, `k` MUST move to the last row. An empty view MUST leave the cursor in place.

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

#### Scenario: A finished download leaves the view

- **GIVEN** one download has finished and another is running
- **WHEN** the operator opens the download view
- **THEN** the finished download MUST be absent
- **AND** the running download MUST remain

#### Scenario: Completion removes the row

- **GIVEN** a download is running in the download view
- **WHEN** that download finishes
- **THEN** its row MUST leave the download view
- **AND** the audio MUST remain in the cassette

#### Scenario: j and k wrap the download rows

- **GIVEN** the download view lists two downloads and the cursor is on the last
- **WHEN** the operator presses `j`
- **THEN** the cursor MUST be on the first download
- **AND** `k` MUST move it back to the last download
