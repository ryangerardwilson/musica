## RENAMED Requirements

- FROM: `### Requirement: nebula s3 backs up application state`
- TO: `### Requirement: nebula keep backs up application state`

## MODIFIED Requirements

### Requirement: nebula keep backs up application state

musica MUST keep downloads and the saved music directory in its own sqlite file. While the machine is up and `SQLX_ROOT` is unset, that file MUST be `$XDG_RUNTIME_DIR/musica/musica.sqlite`. nebula keep MUST store the backup at `s3://backups/musica.sqlite`: after each write musica MUST run `nebula keep mark musica <path>`, and MUST NOT upload the file itself. The application MUST NOT call nebula sql. A failed mark MUST NOT roll back the local sqlite write. When the runtime file is missing, musica MUST run `nebula keep restore musica <path>` before it creates a file, and when that restore fails it MUST fail to open and MUST NOT create the file. Audio files MUST remain in the configured music directory.

#### Scenario: Restore a missing database

- **GIVEN** the runtime sqlite file is missing and `s3://backups/musica.sqlite` exists
- **WHEN** musica opens its database
- **THEN** it MUST restore that object through `nebula keep restore musica <path>` before creating a new file

#### Scenario: Publish after a download changes

- **GIVEN** a download is queued, finished, failed, or cancelled
- **WHEN** musica records that state
- **THEN** it MUST run `nebula keep mark musica <path>` after the write commits

## ADDED Requirements

### Requirement: The player never waits on the store or the audio

The frame tick MUST NOT read the database. The download view MUST be read in the background about every 500 ms, and each answer MUST replace the screen's copy whole. Queueing and cancelling a download MUST show `queued` or `cancelling` at once and MUST run in the background; a refused write MUST bring the row back and show the error. Starting, seeking, restarting, and stopping playback MUST paint `starting…`, `seeking…`, or `stopping…` before ffplay or ffmpeg is started or stopped, and MUST run off the update. `q` MUST paint `stopping…` and quit once playback has stopped.

#### Scenario: A slow stop does not freeze the keys

- **GIVEN** a track is playing
- **WHEN** the operator plays another track
- **THEN** the screen MUST show `starting…` before the old player has stopped

#### Scenario: A refused cancel brings the row back

- **WHEN** the store refuses a cancel
- **THEN** the download row MUST return to its state before the cancel
- **AND** the status MUST show the error
