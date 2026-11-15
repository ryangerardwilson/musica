package cli

const Usage = "Use: musica | musica play <track> | musica open <music_dir> | musica config | musica end"

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
  launch the Bubble Tea music player with your saved library
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
