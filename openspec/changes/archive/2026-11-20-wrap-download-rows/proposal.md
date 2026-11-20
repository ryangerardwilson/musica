## Why

`j` and `k` stop at the ends of the download view. Moving past the last row should return to the first, and moving above the first should return to the last.

## What Changes

- On the download view, `j` and `k` wrap the rows.
- An empty view stays put.
- Reordering tracks still stops at the ends of the cassette.

## Capabilities

### Modified Capabilities

- `musica`: `j` and `k` wrap the download rows.

## Impact

- `internal/app` download motion.
- The library, playlist, and track lists already wrap and stay as they are.
