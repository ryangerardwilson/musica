## MODIFIED Requirements

### Requirement: musica is the player command

The module SHALL build one command, `musica`. It MUST play a track, open a directory for one run, edit config, serve the player to the Tesseract Musica center, and end the running session. The saved library MUST open in the Tesseract Musica center, the third center of the Super+T launcher. Bare `musica` with no arguments MUST run `<config>/quickshell/omarchy-launcher/open.sh musica`, `<config>` being `$XDG_CONFIG_HOME` when it is absolute and `~/.config` otherwise, and MUST NOT exec `r2d2`. `musica serve`, `musica end`, `musica download-worker`, `musica play`, `musica open`, `musica config`, `musica help`, `musica version`, and `musica upgrade` MUST stay in the musica binary. `musica play` and `musica open` MUST still open the terminal TUI. Help MUST say the player opens in the Tesseract Musica center with `musica` or `r2d2 musica`. Config MUST live at `$XDG_CONFIG_HOME/r2d2/musica/config.json`, or `~/.config/r2d2/musica/config.json` when `XDG_CONFIG_HOME` is unset, and musica MUST write only that path. When that file is absent, musica MUST copy `<config>/musica/config.json` once, or else `<config>/rlc/config.json` once, before it reads or writes config, and MUST leave the old file in place. `musica end` MUST return only once the session has let go of its lock, so that a `musica serve` started next holds the session at once, and MUST fail with `musica did not end` when the session still holds it after five seconds. Before it sends `end`, `musica end` MUST note the track the session's ffplay has loaded, playing or paused, and MUST read it from the session's ffplay process, because a terminal session from an older musica answers only `end`; with no ffplay loaded it MUST forget an older note. The next `musica serve` MUST take that note and open the track's cassette with the track selected and the cassette menu open, its frames naming the track with the state `paused` and no elapsed time. It MUST NOT start ffplay for that track and MUST NOT restart it from the beginning on its own: Space MUST say there is nothing to resume, Enter on the track MUST play it from the start, and a stop MUST put the held track down.

#### Scenario: Launch the saved library

- **GIVEN** a saved `music_dir`
- **WHEN** the operator runs `musica` or `r2d2 musica` with no further arguments
- **THEN** the Tesseract Musica center MUST show that library
- **AND** playback MUST continue when the operator leaves the center

#### Scenario: Bare musica opens the Tesseract center

- **WHEN** the operator runs `musica` with no arguments
- **THEN** musica MUST run `~/.config/quickshell/omarchy-launcher/open.sh musica`
- **AND** it MUST NOT exec `r2d2`
- **AND** the environment MUST pass through unchanged

#### Scenario: Adopt the musica config

- **GIVEN** `~/.config/musica/config.json` exists and `~/.config/r2d2/musica/config.json` does not
- **WHEN** musica loads config
- **THEN** it MUST copy the musica config to `~/.config/r2d2/musica/config.json`
- **AND** `~/.config/musica/config.json` MUST stay in place
- **AND** later saves MUST write only `~/.config/r2d2/musica/config.json`

#### Scenario: Adopt the previous config

- **GIVEN** `~/.config/rlc/config.json` exists and neither `~/.config/musica/config.json` nor `~/.config/r2d2/musica/config.json` does
- **WHEN** musica loads config
- **THEN** it MUST copy the rlc config to `~/.config/r2d2/musica/config.json`

#### Scenario: The new config wins

- **GIVEN** `~/.config/r2d2/musica/config.json` exists
- **WHEN** musica loads config
- **THEN** it MUST NOT copy an old file over it

#### Scenario: musica end waits for the session to let go

- **GIVEN** a session holds the musica control socket
- **WHEN** the operator runs `musica end`
- **THEN** it MUST return only after the session has let go of its lock
- **AND** a `musica serve` started next MUST hold the session

#### Scenario: A paused terminal track opens paused in musica serve

- **GIVEN** a terminal TUI holds the session with `focus/Flow.mp3` loaded and paused
- **WHEN** `musica end` ends it and `musica serve` starts
- **THEN** the first frame MUST show the `focus` cassette with `Flow` selected, the title `Flow`, and the state `paused`
- **AND** no ffplay MUST start
- **AND** Space MUST NOT start `Flow.mp3`

### Requirement: q ends the session

`q` MUST stop playback and end the session when no text prompt is open, including from the library, the cassette menu, help, and the download view: the terminal TUI quits, and a `musica serve` session sends `ended` to its watchers and exits. `q` MUST NOT cancel a background download. While a text prompt is open, `q` MUST be typed and MUST NOT quit.

#### Scenario: q closes a playing session

- **GIVEN** a track is playing in the TUI or in `musica serve`, and no text prompt is open
- **WHEN** the operator presses `q`
- **THEN** playback MUST stop
- **AND** the TUI MUST quit, or the served session MUST send `ended` and exit

#### Scenario: q from the download view leaves the download running

- **GIVEN** a track is playing and the download view shows a running download
- **WHEN** the operator presses `q`
- **THEN** the session MUST end
- **AND** that download MUST keep running

#### Scenario: q types into a prompt

- **GIVEN** a text prompt is open
- **WHEN** the operator presses `q`
- **THEN** the prompt MUST contain `q`
- **AND** the session MUST stay open

### Requirement: r2d2 opens the player

The module MUST provide the public package `github.com/ryangerardwilson/musica/player` with `Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int`. `Run` MUST run the same player command as the musica binary: with no arguments it MUST open the Tesseract Musica center as bare `musica` does, and every other argument list MUST behave as it does on the binary. `Run` MUST NOT exec `r2d2`. The player implementation MUST stay in this module.

#### Scenario: r2d2 opens the library

- **WHEN** r2d2 calls `player.Run` with no arguments
- **THEN** the launcher's `open.sh musica` MUST run
- **AND** the Tesseract Musica center MUST show the saved library

#### Scenario: r2d2 ends the session

- **WHEN** r2d2 calls `player.Run` with `end`
- **THEN** the running musica session MUST end

## ADDED Requirements

### Requirement: musica serve streams the player to the Tesseract center

`musica serve` SHALL run the player with no terminal. It SHALL hold the musica session and its control socket as the TUI does, play through ffplay, keep playing after every client disconnects, and end only on `q`, Ctrl+C, `musica end`, SIGTERM, SIGINT, or SIGHUP, stopping playback first. A client that writes `watch` and a newline SHALL receive one JSON object per line: a `frame` after each player update, holding the playing state, elapsed time, and title, the status line, the view, the open cassette, up to nine rows around the selection, the open prompt and its text, help, the download view with its rows and selection, and the visualizer's 32 band `levels` and `peaks` with `bass`, `energy`, and `thump`, each from 0 to 1, `thump` being the hardest low-end hit the analyzer heard since the frame before, so a held tone hits only as it starts; a `leave` event; and an `ended` event when the session ends. Each line the client writes after `watch` SHALL be a JSON object `{"key":"<name>"}` that the player takes as that key with the TUI's rules, except that `m` with no prompt open, and Escape from anywhere, SHALL send `leave` and leave playback, the menu, and the selection as they were. Escape SHALL back out of nothing in `musica serve`: it SHALL NOT cancel a prompt, close help or the download view, or go back from a cassette, which stay for the next showing. `h` with no prompt or leader open SHALL go back to the cassette list from an open cassette's tracks, help, or the download view, showing the menu with that cassette selected, and on the cassette list SHALL move the selection as in the TUI. The terminal TUI SHALL keep Escape as back. A client that cannot keep up SHALL be dropped and MUST NOT stall the player. No frame SHALL be built while no client watches. `end` SHALL still end the session. The TUI session SHALL answer `watch` with `unknown command`.

#### Scenario: A watcher gets frames

- **GIVEN** `musica serve` is playing a track
- **WHEN** a client writes `watch`
- **THEN** it MUST receive `frame` lines whose `levels` and `peaks` hold 32 numbers from 0 to 1
- **AND** each frame MUST name the playing track

#### Scenario: A kick reaches one frame

- **GIVEN** a watcher and a track playing
- **WHEN** a kick lands in one analysis window and quiet windows follow
- **THEN** the next frame's `thump` MUST be near 1
- **AND** the frames after it MUST carry a `thump` near 0

#### Scenario: A held tone does not thump

- **GIVEN** a watcher and a track that holds one low tone
- **WHEN** frames follow the tone's start
- **THEN** their `thump` MUST stay near 0

#### Scenario: A key line drives the player

- **GIVEN** a watcher and the cassette list showing its last row selected
- **WHEN** the client writes `{"key":"j"}`
- **THEN** a later frame MUST show the first row selected

#### Scenario: Space pauses and opens the cassette actions

- **GIVEN** a watcher and a track playing with the cassette menu closed
- **WHEN** the client writes `{"key":" "}`, and then `{"key":" "}` again
- **THEN** playback MUST pause and a frame MUST show the menu open
- **AND** the second space MUST resume playback with the menu closed and the frame state `playing`

#### Scenario: m asks the center to leave

- **GIVEN** a watcher, a track playing, and no prompt open
- **WHEN** the client writes `{"key":"m"}`
- **THEN** the client MUST receive `leave`
- **AND** playback MUST continue

#### Scenario: Escape leaves from anywhere and backs out of nothing

- **GIVEN** a watcher and an open cassette, help, the download view, or a prompt
- **WHEN** the client writes `{"key":"esc"}`
- **THEN** the client MUST receive `leave`
- **AND** a later frame MUST still show the open cassette, help, the download view, or the prompt and its text

#### Scenario: h goes back to the cassette list

- **GIVEN** a watcher and an open cassette's tracks, help, or the download view, with no prompt open
- **WHEN** the client writes `{"key":"h"}`
- **THEN** a later frame MUST show the cassette list with the menu open and that cassette selected
- **AND** the client MUST NOT receive `leave`

#### Scenario: Playback outlives the watcher

- **GIVEN** `musica serve` is playing a track for a watcher
- **WHEN** the watcher disconnects
- **THEN** playback MUST continue and the session MUST stay up

#### Scenario: The TUI refuses watch

- **GIVEN** the terminal TUI holds the session
- **WHEN** a client writes `watch`
- **THEN** it MUST read `unknown command`
