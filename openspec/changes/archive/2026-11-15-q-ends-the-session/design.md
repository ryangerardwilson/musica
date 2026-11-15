## Context

Ctrl+C and `musica end` call `close`, which stops the player and visualizer, then quit the TUI. `q` currently does neither. Escape hides the window and leaves the process running.

## Goals / Non-Goals

**Goals:**

- `q` outside a text prompt takes the same path as Ctrl+C.
- A running download is not cancelled.
- Typing `q` into search, rename, or a YouTube prompt still inserts `q`.

**Non-Goals:**

- Changing Escape, `x`, or `musica end`.

## Decisions

1. Handle `q` beside the Ctrl+C check in `Update`, and only when `inputKind` is empty. That covers help and the download view without growing `updateKey`.
2. `close` already stops playback and leaves the download worker alone.

## Risks / Trade-offs

- `q` no longer does nothing. Rollback is restoring the old key check and the "q does not quit" test.
