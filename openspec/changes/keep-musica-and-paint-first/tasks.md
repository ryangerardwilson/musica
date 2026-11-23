## 1. Player

- [x] 1.1 Background download list poll; queue and cancel as pending intents
- [x] 1.2 Playback changes run off the update and paint first
- [x] 1.3 Store marks and restores through `nebula keep`; own backup code deleted

## 2. Verification

- [x] 2.1 `c3po ast exec -- openspec validate keep-musica-and-paint-first --strict --no-interactive`
- [x] 2.2 `GOWORK=off go test ./...`
