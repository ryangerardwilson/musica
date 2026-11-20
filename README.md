# musica

`musica` is a terminal-native local music player written in Go with a Bubble Tea
TUI, analog-shaped `ffplay` playback, and an `ffmpeg`-powered ASCII visualizer.
It keeps download jobs and the saved library path in its own sqlite file.
[nebula](https://github.com/ryangerardwilson/nebula) `s3` backs that file up.

It keeps the narrow local-file loop: launch the saved library, open a directory
for one run, play a track directly, or edit the user config.

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
r2d2 musica                    # launch the saved library as an r2d2 screen
musica                         # exec r2d2 musica
musica play /path/to/track.mp3 # launch and play this track
musica open /path/to/music_dir # use this directory for this run
musica config                  # open the user config in your editor
musica end                     # stop playback and end the running session
musica help                    # show help
musica version                 # print version
musica upgrade                 # upgrade to latest release
```

Config defaults to:

- `$XDG_CONFIG_HOME/musica/config.json` if `XDG_CONFIG_HOME` is set
- `~/.config/musica/config.json` otherwise

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
- `Esc` - return to cassettes; from the cassette menu, hide the overlay while playback continues
- `m` - toggle the cassette and track menu, which opens by default
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

## Storage

Config lives at `$XDG_CONFIG_HOME/musica/config.json`, or
`~/.config/musica/config.json`. An existing rlc config is copied once.

Downloads and the saved music directory live in `musica.sqlite`. While the
machine is up the file is `$XDG_RUNTIME_DIR/musica/musica.sqlite`. A set
`SQLX_ROOT` keeps it at `$SQLX_ROOT/data/musica.sqlite` and skips s3. nebula
`s3` stores the backup at `s3://backups/musica.sqlite`. Audio files stay in
the music directory.

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
- `internal/app/` - Bubble Tea TUI
- `internal/config/` - user config path, load/save, first-run defaults
- `internal/library/` - cassette directory and local music scanning
- `internal/download/` - background YouTube downloads and the progress view
- `internal/store/` - sqlite state and nebula s3 backup
- `internal/player/` - `ffplay` process backend
- `internal/visualizer/` - `ffmpeg` spectrum analyzer and ASCII renderer
