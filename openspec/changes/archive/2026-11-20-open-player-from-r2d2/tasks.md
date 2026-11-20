## 1. Player entry

- [x] 1.1 `player.Run` runs the existing player command
- [x] 1.2 Bare `musica` execs `r2d2 musica`, and help names `r2d2 musica`

## 2. Download worker

- [x] 2.1 The worker execs `musica download-worker` from PATH and fails without a fallback when musica is missing

## 3. Verification

- [x] 3.1 `openspec validate open-player-from-r2d2 --strict --no-interactive`
- [x] 3.2 `GOWORK=off go test ./...`
