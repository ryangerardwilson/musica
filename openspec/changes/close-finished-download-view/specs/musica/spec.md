## ADDED Requirements

### Requirement: A finished download closes an empty download view

When a download finishes and no queued, running, failed, or cancelled download remains, the download view MUST close. The finished audio MUST remain in the cassette. A download that is still queued, running, failed, or cancelled MUST keep the download view open, and a finished download MUST stay out of that view.

#### Scenario: The last download finishing closes the view

- **GIVEN** the download view is open on one running download
- **WHEN** that download finishes
- **THEN** the download view MUST close
- **AND** the audio MUST remain in the cassette

#### Scenario: Another download keeps the view open

- **GIVEN** the download view is open on one running download and one queued download
- **WHEN** the running download finishes
- **THEN** the download view MUST stay open
- **AND** the queued download MUST remain
- **AND** the finished download MUST be absent
