## 1. Spec

- [x] 1.1 `c3po ast exec -- openspec validate play-in-the-tesseract-center --strict --no-interactive`
- [x] 1.2 Living `openspec/specs/musica/spec.md` matches the modified and added requirements

## 2. Control socket

- [x] 2.1 `watch` streams published lines and reads key lines; `end` unchanged; the TUI answers `unknown command`
- [x] 2.2 A watcher whose queue fills is dropped; nothing is built while no one watches

## 3. Served player

- [x] 3.1 `musica serve` runs the model with no renderer and no input and ends on `q`, Ctrl+C, `musica end`, and signals
- [x] 3.2 Frames carry state, rows, prompt, help, downloads, levels, peaks, bass, and energy
- [x] 3.3 `m` and the cassette-list Escape send `leave`; key lines map to the TUI's keys (Escape from anywhere since 8.1)

## 4. Command

- [x] 4.1 Bare `musica` and `player.Run` with no arguments run the launcher's `open.sh musica`
- [x] 4.2 Help names the Tesseract Musica center and `musica serve`

## 5. Verification

- [x] 5.1 `GOWORK=off go test ./...`
- [x] 5.2 Install from `/tmp` with `bash /home/ryan/OpenSource/musica/install.sh from /home/ryan/OpenSource/musica`

## 6. Takeover

- [x] 6.1 `musica end` notes the track the session's ffplay has loaded, sends `end`, and returns once the lock is free, or fails with `musica did not end`
- [x] 6.2 `musica serve` takes the note and holds that track paused in its cassette without starting ffplay; Space does not start it, Enter plays it from the start, `x` puts it down
- [x] 6.3 `musica serve` resumes queued downloads beside the player, so the first frame does not wait on the store
- [x] 6.4 The living spec carries the modified `musica is the player command`; `c3po ast exec -- openspec validate play-in-the-tesseract-center --strict --no-interactive`
- [x] 6.5 `GOWORK=off go test ./...` with `SQLX_ENGINE` unset
- [x] 6.6 Install from `/tmp` with `bash /home/ryan/OpenSource/musica/install.sh from /home/ryan/OpenSource/musica`

## 7. Kick

- [x] 7.1 Frames carry `thump`, the hardest low-end hit since the frame before. The analyzer hears it in windows that overlap by half, over the louder of its slow mean and what is left of the windows before, and every served update takes it, watched or not
- [x] 7.2 `GOWORK=off go test ./...` with `SQLX_ENGINE` unset: a kick in one window is near 1 on its frame and near 0 on the frames after; a held tone hits only as it starts; a kick no one watched is not sent late
- [x] 7.3 `c3po ast exec -- openspec validate play-in-the-tesseract-center --strict --no-interactive`; the living spec carries the same `thump` sentence and scenarios
- [x] 7.4 Install from `/tmp` with `bash /home/ryan/OpenSource/musica/install.sh from /home/ryan/OpenSource/musica`

## 8. Escape and h

- [x] 8.1 In a served player Escape sends `leave` from anywhere and backs out of nothing; `h` goes back to the cassette list from a cassette's tracks, help, or the download view with that cassette selected, and moves on the cassette list; the TUI keeps Escape as back and `h` as up
- [x] 8.2 `GOWORK=off go test ./...` with `SQLX_ENGINE` unset; the new tests fail on the model before the change; a real `musica serve` on a temp runtime goes back on `h` and stays on the tracks on `esc`
- [x] 8.3 `c3po ast exec -- openspec validate play-in-the-tesseract-center --strict --no-interactive`; the living spec carries the same sentence and scenarios
- [x] 8.4 Install from `/tmp` with `bash /home/ryan/OpenSource/musica/install.sh from /home/ryan/OpenSource/musica`; the running `musica serve` was left alone and keeps the old keys until it is replaced
