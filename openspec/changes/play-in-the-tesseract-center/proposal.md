## Why

Musica opens with Super+M as its own Alacritty popup: a second full-screen surface with its own ASCII rain on top of the widget-chrome rain, outside the Super+T menu where the tmux sessions and the space navigator live. The operator wants Musica to be the third center of that menu, drawn in the launcher with the launcher's one rain, and Super+M is going away. The launcher can only show the player if the player runs somewhere it can reach and tells it what is playing, levels included.

## What Changes

- `musica serve` runs the existing player model with no terminal. It holds the musica session and the control socket as the TUI does, and keeps playing after the launcher closes.
- The control socket learns `watch`: a client that writes it gets one JSON object per line, a `frame` per player update with the playing state, the list around the selection, the prompt, help, the download rows, and the visualizer's 32 band levels and peaks with bass and energy; then `leave` and `ended` events. Lines the client writes after `watch` are keys, taken with the TUI's rules.
- In the served player, `m`, and Escape from anywhere, send `leave` instead of toggling the menu or hiding a terminal. Escape backs out of nothing there; `h` goes back to the cassette list from a cassette's tracks, help, or the download view. `q` still stops playback and ends the session.
- Bare `musica`, and `player.Run` with no arguments, run the launcher's `open.sh musica`, which shows the Tesseract Musica center. Bare `musica` no longer execs `r2d2 musica`.
- `musica end` returns once the session has let go of its lock, so `musica serve` can start at once. Before it sends `end` it notes the track the session's ffplay has loaded, playing or paused. A terminal session from an older musica answers only `end`, so the track is read from its ffplay child in `/proc`.
- `musica serve` takes that note and opens the track's cassette with the track selected and shown paused. It starts nothing, because where the ended session was in the track is not known. This is how the Tesseract center takes over a terminal that holds the session.
- `musica play`, `musica open`, `musica config`, `musica end`, `musica download-worker`, `musica help`, `musica version`, and `musica upgrade` keep their behavior; `play` and `open` still open the terminal TUI.

## Capabilities

### Modified Capabilities

- `musica`: the saved library opens in the Tesseract Musica center; bare `musica` and `player.Run` with no arguments open that center; `q` ends a served session too; `musica serve` streams the player over the control socket; `musica end` waits for the session to let go and hands its track to the next `musica serve`, which holds it paused.

## Non-goals

- No visualizer drawing for the center in this module. The launcher draws it from the levels; the terminal TUI keeps its own rain.
- No second socket, no sqlite or hyprctl reads for the client, and no change to the download worker, the store, or nebula keep.
- No r2d2 change here. r2d2 runs `open.sh musica` itself for bare `r2d2 musica` (r2d2 change `drop-musica-from-the-menu`) and keeps its released musica pin.

## Impact

- New: `internal/app/center.go` and its frame and key code, `musica serve`, the `watch` command in `internal/session`, and `internal/session/held.go`, the held track and the wait in `musica end`.
- Changed: `internal/cli` dispatch and help, `player` doc.
- The launcher (`~/.config/quickshell/omarchy-launcher`) must accept `open.sh musica` and run `musica serve`; that lives in the home repository.
