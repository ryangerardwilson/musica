# musica

`musica` is a terminal-native local music player written in Go with a Bubble Tea
TUI, analog-shaped `ffplay` playback, and an `ffmpeg`-powered ASCII visualizer.
It keeps download jobs and the saved library path in its own sqlite file.
[nebula](https://github.com/ryangerardwilson/nebula) `keep` backs that file up.

It keeps the narrow local-file loop: launch the saved library, open a directory
for one run, play a track directly, or edit the user config.

The saved library opens in the Tesseract Musica center, the third center of the
Super+T launcher, next to the tmux sessions and the space navigator. The
launcher runs `musica serve`, which plays with no terminal, and draws the
center from what that session streams over its control socket. The terminal
TUI is still what `musica play` and `musica open` show.

Playback always runs through a light analog-style filter chain for EQ,
compression, soft clipping, and limiting.

## Installation

### Prebuilt binary

```bash
curl -fsSL https://raw.githubusercontent.com/ryangerardwilson/musica/main/install.sh | bash
```

If `~/.local/bin` is not already on your `PATH`, add it once to your shell:

```bash
export PATH="$HOME/.local/bin:$PATH"
```

Install a specific version:

```bash
curl -fsSL https://raw.githubusercontent.com/ryangerardwilson/musica/main/install.sh | bash -s -- version 0.1.0
```

### From source

```bash
go build -o musica ./cmd/musica
./musica help
```

## Command-Line Usage

```bash
musica                         # show the saved library in the Tesseract Musica center
r2d2 musica                    # the same, through the player package
musica serve                   # play for that center with no terminal; the launcher starts it
musica play /path/to/track.mp3 # play this track in the terminal TUI
musica open /path/to/music_dir # use this directory for one terminal run
musica config                  # open the user config in your editor
musica end                     # stop playback, end the running session, and wait until it lets go
musica help                    # show help
musica version                 # print version
musica upgrade                 # upgrade to latest release
```

Config defaults to:

- `$XDG_CONFIG_HOME/r2d2/musica/config.json` if `XDG_CONFIG_HOME` is set
- `~/.config/r2d2/musica/config.json` otherwise

On first run, if `music_dir` is not configured, `musica` prompts for a music
directory when stdin is interactive. Non-interactive first runs save `~/Music`.

```json
{
  "music_dir": "/home/you/Music",
  "fps": 20
}
```

## Keyboard Shortcuts

- `h` / `l` / `j` / `k` - move across cassettes or tracks
- `Ctrl+j` / `Ctrl+k` - move selected track down/up in the current in-memory order
- `Enter` - open selected cassette or play selected track
- `Esc` - return to cassettes; from the cassette list, hide the player while playback continues. In the Tesseract center Escape closes the menu from anywhere and backs out of nothing
- `m` - toggle the cassette and track menu, which opens by default; in the Tesseract center, leave Musica
- `h` in the Tesseract center - back to the cassette list from a cassette's tracks, help, or the download view; on the cassette list it moves up
- `Space` - pause and open the cassette menu; press again to resume
- `f` / `b` - seek forward/backward by 10 seconds
- `r` - restart the current track
- `x` - stop playback
- `dd` - delete selected cassette or track
- `s` - shuffle the current track list
- `,rn` - rename selected cassette or track
- `y` - download YouTube audio into the open cassette
- `u` - show downloads. A running download animates its progress. A finished download leaves the list
- `x` in the download view - cancel the selected download. Playback continues
- `/` - search
- `n` / `N` - next/previous search result; without search, `n` creates a cassette
- `?` - toggle shortcuts screen
- `q` - stop playback and quit, except while typing in a prompt
- `Ctrl+C` - stop playback and quit from any view or prompt

YouTube downloads prompt for the URL first, then the song name. The download
keeps running after the TUI closes and while other tracks play. Cancel it from
the download view.

## Control Socket

The running session holds `$XDG_RUNTIME_DIR/musica/control.sock`. `musica end`
writes `end` to it. A `musica serve` session also takes `watch`: the client
then reads one JSON object per line. A `frame` follows each player update with
the playing state, elapsed time, and title, the status line, up to nine rows
around the selection, the open prompt, help, the download rows, and the
visualizer's 32 `levels` and `peaks` with `bass`, `energy`, and `thump`, each
from 0 to 1. `thump` is the hardest low-end hit since the frame before, not
eased, so a kick is near 1 on its frame and gone on the next, and a held tone
hits only as it starts. `leave` comes when `m`, or Escape from anywhere, asks
the center to close; playback goes on, and Escape backs out of nothing there.
`ended` comes when the session ends. Each line the client writes after `watch`
is a key such as `{"key":"j"}` or `{"key":"h"}`, read with the TUI's rules,
except that `h` goes back to the cassette list from a cassette's tracks, help,
or the download view. The terminal TUI answers `watch` with `unknown command`.

`musica end` returns once the session has let go of its lock, so a
`musica serve` started next takes the session at once. Before it sends `end` it
notes the track the session's ffplay has loaded, playing or paused, in
`$XDG_RUNTIME_DIR/musica/held`; a terminal session from an older musica only
answers `end`, so the track is read from its ffplay child in `/proc`. The next
`musica serve` takes that note: it opens the track's cassette with the track
selected and shown paused, and does not start it, because where the ended
session was in the track is not known. Enter plays it from the start. The
Tesseract center uses this when a terminal holds the session: it runs
`musica end`, then `musica serve`.

## Storage

Config lives at `$XDG_CONFIG_HOME/r2d2/musica/config.json`, or
`~/.config/r2d2/musica/config.json`. When that file is absent, an existing
`~/.config/musica/config.json`, or else an rlc config, is copied once and left
in place.

Downloads and the saved music directory live in `musica.sqlite`. While the
machine is up the file is `$XDG_RUNTIME_DIR/musica/musica.sqlite`. A set
`SQLX_ROOT` keeps it at `$SQLX_ROOT/data/musica.sqlite` and skips the backup.
nebula keep stores the backup at `s3://backups/musica.sqlite`: musica runs
`nebula keep mark musica <file>` after each change it records and
`nebula keep restore musica <file>` when the file is missing. When the backup
cannot be reached, musica makes no empty file in its place. `NEBULA_BIN`
names the nebula binary; otherwise `nebula` is found on `PATH`. Audio files
stay in the music directory.

## Requirements

- Linux x86_64 for the current prebuilt installer
- `ffplay` available in `PATH` for playback
- `ffmpeg` available in `PATH` for the visualizer
- `yt-dlp` available in `PATH` for YouTube downloads

If `ffplay` is missing, the TUI still opens but playback commands report the
missing executable.

## Project Layout

- `cmd/musica/` - executable entrypoint
- `internal/cli/` - command dispatch, config opening, and upgrade handoff
- `internal/app/` - Bubble Tea player model: the terminal TUI and `musica serve`
- `internal/session/` - the session lock and the control socket
- `internal/config/` - user config path, load/save, first-run defaults
- `internal/library/` - cassette directory and local music scanning
- `internal/download/` - background YouTube downloads and the progress view
- `internal/store/` - sqlite state and nebula keep backup
- `internal/player/` - `ffplay` process backend
- `internal/visualizer/` - `ffmpeg` spectrum analyzer and ASCII renderer
