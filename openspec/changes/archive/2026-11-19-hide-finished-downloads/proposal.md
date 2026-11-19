## Why

The download view keeps every finished song. The operator cannot see the download that is still running.

## What Changes

- A finished download leaves the download view.
- A queued, running, failed, or cancelled download stays visible.
- The audio file stays in the cassette. The sqlite row stays stored.

## Capabilities

### Modified Capabilities

- `musica`: the download view shows work that is not finished.

## Impact

- Download list filtering in `internal/app`.
- The `u` shortcut description in the README.
