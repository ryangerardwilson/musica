## Why

The colored player is louder than the operator wants. Musica should read as simple white text on the black terminal.

## What Changes

- Draw titles, status, help, the YouTube prompt, the download view, and visualizer glyphs in white.
- Keep the dark field dark. Do not use a separate hue for errors, selection, playback, or download state.

## Capabilities

### Modified Capabilities

- `musica`: replace the colored surface with white text. Pause still opens the cassette menu.

## Impact

- Foreground styles in `internal/app` and `internal/download` become white.
- Playback, downloads, and the pause menu stay as they are.
