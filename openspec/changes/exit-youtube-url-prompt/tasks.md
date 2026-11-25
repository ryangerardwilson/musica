## 1. Prompt

- [x] 1.1 Submitting `/exit` closes the YouTube URL prompt and the song-name prompt without queueing a download

## 2. Verification

- [x] 2.1 `c3po ast exec -- openspec validate exit-youtube-url-prompt --strict --no-interactive`
- [x] 2.2 `GOWORK=off go test ./internal/app/`
