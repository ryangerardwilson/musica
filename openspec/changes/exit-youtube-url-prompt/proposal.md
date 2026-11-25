## Why

In the Tesseract Musica center, Escape closes the whole menu and is never sent to the player, so a YouTube URL prompt cannot be cancelled. The operator is stuck in that prompt.

## What Changes

- Submitting `/exit` in the YouTube URL prompt or the song-name prompt closes that prompt, clears any pending URL, and does not queue a download.
- `/exit` is not a URL and not a song name.

## Capabilities

### Modified Capabilities

- `musica`: `/exit` closes a YouTube prompt.

## Impact

- `internal/app` input commit for the YouTube prompts.
- The terminal player gains the same exit, because both fronts share the model. Escape still cancels a prompt in the terminal player.
