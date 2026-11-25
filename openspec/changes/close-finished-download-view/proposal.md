## Why

After a YouTube download finishes, its row leaves the list and the download view stays open. In the Tesseract center that view is the box over the picture, so the operator is left on DOWNLOADS NONE instead of the picture.

## What Changes

- When a download finishes and no queued, running, failed, or cancelled download remains, the download view closes.
- A download that is still queued, running, failed, or cancelled keeps the view open.
- The finished audio stays in the cassette.

## Capabilities

### Modified Capabilities

- `musica`: a finished download closes the download view when nothing remains.

## Impact

- `internal/app` download list update.
- The terminal player and `musica serve` share that update. The center draws whatever the frame says is shown.
