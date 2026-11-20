## Context

The download view adds the delta and then clamps. The library lists already wrap with a remainder.

## Goals / Non-Goals

**Goals:**

- `j` and `k` wrap the download rows.
- An empty view does not move.

**Non-Goals:**

- Wrapping track reorder. That motion swaps two tracks and still stops at the ends.
- Changing which download `x` cancels.

## Decisions

- Use the same remainder as the library lists: `(index + delta%n + n) % n`.
- Leave `moveTrack` clamped.

## Risks / Trade-offs

- An already-open player keeps the current binary until it quits.
