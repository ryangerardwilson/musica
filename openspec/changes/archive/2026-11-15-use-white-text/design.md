## Context

The accepted spec requires amber, gold, mint, coral, and rose, plus a teal-to-gold visualizer. The operator wants that color gone and the font left white.

## Goals / Non-Goals

**Goals:**

- Every drawn glyph that is text or rain uses the terminal white, ANSI 15.
- Empty visualizer cells stay dark so the field does not turn into a white sheet.
- Visualizer rows still share one base color.

**Non-Goals:**

- A theme setting.
- Changing layout, keys, or pause behavior.

## Decisions

1. One white style for text. Bold stays on titles and the selected download row so weight, not hue, marks them.
2. Rain intensities share that white. The dark core stays ANSI 240.
3. Download rows no longer change hue by status.

## Risks / Trade-offs

- A failed download is no longer red. The word "failed" remains in the row.
- Rollback is restoring the colored styles and the color requirement.
