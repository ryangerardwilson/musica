## 1. View

- [x] 1.1 Close the download view when a completion leaves no queued, running, failed, or cancelled download

## 2. Verification

- [x] 2.1 `c3po ast exec -- openspec validate close-finished-download-view --strict --no-interactive`
- [x] 2.2 `GOWORK=off go test ./internal/app/`
