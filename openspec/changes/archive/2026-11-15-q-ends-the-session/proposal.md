## Why

`q` leaves the Musica session open. The operator wants `q` to close that session.

## What Changes

- `q` stops playback and quits the TUI from the library, the menu, help, and the download view.
- `q` still types while a text prompt is open.
- A background download keeps running. Escape still hides the overlay without ending the session.

## Capabilities

### Modified Capabilities

- `musica`: `q` ends the running session the way Ctrl+C does, except while text is being entered.

## Impact

- Key handling in `internal/app`, the shortcut list, and the README.
- Playback stops. The download worker is unchanged.
