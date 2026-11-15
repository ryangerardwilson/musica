## REMOVED Requirements

### Requirement: The player surface uses color

The colored surface is replaced by white text.

## ADDED Requirements

### Requirement: The player text is white

The visualizer glyphs, status line, help, download view, and YouTube prompt MUST use white text. They MUST NOT use a separate hue for titles, status, errors, selection, playback, or download state. Visualizer rows MUST share one base color. Empty visualizer cells MAY stay dark.

#### Scenario: Text is white

- **GIVEN** the player is drawing
- **WHEN** a title, a status line, a rain glyph, a running download, and a failed download are styled
- **THEN** each foreground MUST be white
- **AND** the running download and the failed download MUST share that white

#### Scenario: Visualizer rows share one color

- **GIVEN** the visualizer is drawing
- **WHEN** the top, middle, and bottom rows are styled
- **THEN** those rows MUST use the same base color
