## Context

`download.List` returns every sqlite row, including status `done`. The download view renders that slice. Finished rows show "in the cassette" at 100% and crowd out the running download.

## Goals / Non-Goals

**Goals:**

- The download view omits status `done`.
- Finishing a download still reloads the open cassette.
- Failed and cancelled rows stay so the operator can see the outcome.

**Non-Goals:**

- Deleting finished rows from sqlite or from the cassette.
- Hiding a download that is still running at 100%.

## Decisions

1. Filter after `download.List` and before the view stores the slice. Render, selection, and cancel then share one visible list.
2. Compare the previous visible rows with the full list so a transition to `done` still reloads the cassette.
3. Keep the selected row by id when it is still visible. Otherwise select the running row.

## Risks / Trade-offs

- A finished download cannot be reviewed from this view. The cassette has the audio.
