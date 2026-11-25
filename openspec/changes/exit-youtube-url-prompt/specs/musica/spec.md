## ADDED Requirements

### Requirement: /exit closes a YouTube prompt

When the YouTube URL prompt or the song-name prompt is open, submitting `/exit` MUST close that prompt, clear any pending URL, and MUST NOT queue a download. `/exit` MUST NOT be kept as a URL or a song name. Other text MUST keep the prompt's current rules.

#### Scenario: /exit closes the URL prompt

- **GIVEN** the YouTube URL prompt is open
- **WHEN** the operator submits `/exit`
- **THEN** the prompt MUST close
- **AND** no download MUST be queued

#### Scenario: /exit closes the song-name prompt

- **GIVEN** the song-name prompt is open and a URL is waiting
- **WHEN** the operator submits `/exit`
- **THEN** the prompt MUST close
- **AND** the waiting URL MUST be cleared
- **AND** no download MUST be queued

#### Scenario: other text stays on the URL prompt

- **GIVEN** the YouTube URL prompt is open
- **WHEN** the operator submits text that is not a YouTube URL and is not `/exit`
- **THEN** the URL prompt MUST stay open
