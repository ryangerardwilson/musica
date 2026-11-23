## Why

The player's 50 ms frame tick listed downloads with two sqlite3 processes each time, about 40 a second. Queueing and cancelling a download ran sqlite writes inside the update, and the play keys waited up to about 1.2 s inside the update for ffplay and ffmpeg to stop. musica also carried its own backup code: an `s3 put` of the live file in a goroutine and an `s3 get` restore.

## What Changes

- The frame tick only animates. The download list is read every 500 ms in the background and replaces the screen's copy whole.
- Queue and cancel paint `queued` and `cancelling` first, run in the background, and revert with the error when the store refuses.
- Play, seek, restart, stop, and auto-advance paint `starting…`, `seeking…`, or `stopping…` first and run in the background; `q` paints `stopping…` and quits once playback has stopped.
- The sqlite file is backed up by nebula keep: each write runs `nebula keep mark musica <path>`, and a missing file is restored with `nebula keep restore musica <path>`. musica's own upload and restore code is deleted.

## Capabilities

### Modified Capabilities

- `musica`: the backup requirement moves to nebula keep; the screen never waits on the store or the player.

## Impact

Install nebula with `keep` before this musica. `s3://backups/musica.sqlite` keeps its key and format. r2d2 shows the new player once it bumps its musica pin.
