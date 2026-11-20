package cli

const Usage = "Use: r2d2 musica | musica play <track> | musica open <music_dir> | musica config | musica end"

const HelpText = `musica terminal music player

global actions:
  musica help
    show this help
  musica version
    print the installed version
  musica upgrade
    upgrade to the latest release
  musica end
    stop playback and end the running musica session

features:
  open the Bubble Tea music player with your saved library in r2d2
  # r2d2 musica
  r2d2 musica

  bare musica with no arguments execs r2d2 musica
  # musica
  musica

  play a specific track immediately
  # play <track>
  musica play ~/Music/song.mp3

  use a directory for this run without changing saved config
  # open <music_dir>
  musica open ~/Music

  open the user config in your editor
  # config
  musica config
`
