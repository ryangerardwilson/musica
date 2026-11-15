## Why

Pausing leaves the operator on the visualizer with no cassette menu, and the whole surface is grayscale.

## What Changes

- Space that pauses playback opens the cassette menu. Resume leaves the menu as it is. Space does not cancel a download.
- The visualizer, status line, help, download view, and YouTube prompt use color. Visualizer rows keep one base color.

## Capabilities

### New Capabilities

### Modified Capabilities

- `musica`: Pause opens the cassette menu, and the player surface uses color.

## Impact

The TUI palette and the space key change. Playback, downloads, and nebula backup do not. Non-goals are theming config, a new layout, and changing which key pauses.
