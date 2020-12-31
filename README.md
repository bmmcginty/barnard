# Barnard

## Documentation

Please feel free to give suggestions and corrections for this file (as wellas Barnard propper).
Find a sample notification script in examples/.

## Audio boost

If a user is too soft to hear, you can boost their audio.
The audio should drastically increase once you have hit the VolumeUp key over 10 times (from the silent/0 position).
The boost setting is saved per user, just like per user volume.

## FIFO Control

If you pass the --fifo option to Barnard, a FIFO pipe will be created.
You can control Barnard by sending commands to this FIFO.
Each command must end with a  \n (0x0a) character.
Commands may be added at any time.
Per the robustness principle, be liberal in what you receive.
Current Commands:
* error: An error has occured to prevent transmitting audio, or taking another action.
* micup: Start transmitting, just as when you hold down the talk key. Does nothing if you are already transmiting.
* micdown: Stop transmitting, just like when you release your talk key. Does nothing if you are not already transmitting.
* toggle: Toggle your transmission state.
* talk: Synonym for toggle.
* exit: Exit Barnard, just like when you press your quit key.

## Event Notification

You can use the notifycommand parameter in your config file to run a program on certain events.
Each event has the following parameters:
* event: the name of the event
    - join: user has joined the channel you are in
    - leave: user has left the channel you are in
    - micup: you have begun transmitting
    - micdown: you have stopped transmitting
    - connect: you have connected to a server
    - disconnect: you have disconnected from a server
    - msg: the channel you are currently connected to has received a message
    - pm: you have received a private message
* who: the person causing initiation of the event ("me" for self-generated events)
* what: the body of the event as applicable (message, channel name, etc)

Warning:
Keep in mind that Barnard opens an Alsa sound device when starting.
For this reason, any notification command used here will need to be able to work while other sound is playing.
It is recommended that you test  your notification command by hand, while Barnard is running, before including it here.

You can create a command that will take any of these parameters as desired, by prepending the name of the parameter in your command with a % (percent) sign.
As an example, to attempt to play wave files for each event, you could set notifycommand to:
aplay /home/username/sounds/mumble/%event.wav
When you begin transmitting, aplay will attempt to play /home/username/sounds/mumble/micup.wav.
The same will be attempted for the other events, such as leave, join, micdown, etc.

In order to process messages and the like, Barnard will parse your command as a properly quoted shell command.
For this reason, you should put quotes around arguments that have spaces.
If you want to do more complex things, write a shell script (or c application, python script, etc) to process the arguments passed into it.

## Connecting Via Text Interface

You can now manage your server lists in a text GUI.
An Ncurses interface has been created by members of the [F123 Group](https://gitlab.com/f123).
Make sure the folder in which you store the barnard binary is in your path. This should be the default for any f123 user.
Then just run ./barnard-ui from this folder, and follow the instructions.
You can add barnard-ui to your path as well, and access it from anywhere.

## Modifications

This copy of Barnard and it's associated Gumble library have been modified to support usage by blind users.
Our thanks go out to Tim Cooper for the massive amount of work put into this client, originally found at [github.com/layeh/barnard](https://github.com/layeh/barnard).

## Config

By default, the file ${XDG_CONFIG_HOME:-${HOME}/.config}/barnard/config.yaml will hold the configuration for Barnard.
You can have barnard read another file by using the -c option, like `./barnard -c ~/.anotherbarnard.yaml`.
It will be created automatically if it doesn't exist.
If you modify the config file while Barnard is running, your changes may be overwritten.

## Defaults

You can set username and defaultserver in your config file, and they will be used if none is specified when launching barnard.
(Note that the default username (an empty string) and the default server name (localhost:64738) have been the defaults for barnard up to this  point, and have been left that way for compatibility.)

## Audio Devices

You can set the default input and output devices in the config file as well.
Pass the -list_devices parameter to barnard to be given a list of audio input and output devices.
Copy lines from the above list into inputdevice and outputdevice as desired.
To clear your inputdevice or outputdevice options and set them to defaults, set them to "" or delete them entirely.

## Keystrokes

You can see the below keystrokes in your config file.

Pressing tab inside the main window switches between the user/channel tree view and the message input box.
When in the message input box:
* left and right arrow keys move by character
* home/end moves to the beginning/end of the text respectively
* enter submits the entered message

When in the treeview, pressing:
* f5 or f6 on a channel changes the volume for all users in that channel
* f5 or f6 on a user changes the volume for that user.
* enter on de-selected user selects that user for PM mode.
* enter on selected user de-selects the user
* enter on a channel de-selects any selected users (if any) and moves you to the specified channel.

## Volume

If you set the volume for a user (using the F5/F6 keys by default), Barnard will remember that volume, and will keep that user at that volume.
The volume is set for a single user on a single server.
This means you may have to set a person to a custom volume multiple times, if you are both on multiple servers together.

If you set the volume of a channel, you are basically relatively adjusting each user's volume.
If Jim's volume is set to 0.1, and larry's volume is set to 0.9, lowering the channel by one increment will mute Jim,and set Larry to 0.8.

You can change the volume for a user once that user has spoken at least once during a session.
Attempts to change the volume of a user who has not spoken will be ignored.
If you are unable to hear a user speaking, you can edit the barnard.yaml file, after closing Barnard, and set the volume parameter to 1.0 for a particular user.

### Technical

The volume for each user is set via the audio session created for that user; no talking means no session means no settable volume.

## PM Mode

This mode sets the text to the left of your message entry box to [@username], where username is the name of the person you are PMing.
While a user is selected, both private messages and channel messages are displayed.
However, any messages you send will be delivered only to the selected user.
Private messages between you and another party are shown as `pm/source/dest`, where source and dest are the sender and receiver of the message respectively.
When you are finished sending private messages to a user, press tab to reactivate the tree view, and de-select the user as above.

## Info

barnard is a terminal-based client for the [Mumble](https://mumble.info) voice
chat software.

![Screenshot](https://i.imgur.com/B8ldT5k.png)

## Installation

Requirements:

1. [Go](https://golang.org/)
2. [Git](https://git-scm.com/)
3. [Opus](https://opus-codec.org/) development headers
4. [OpenAL](http://kcat.strangesoft.net/openal.html) development headers

To fetch and build:

    go get -u github.com/bmmcginty/barnard

After running the command above, `barnard` will be compiled as `$(go env GOPATH)/bin/barnard`.

To build locally:

    go get -v .
    go build -v ./*.go

## Manual

### Key bindings

- <kbd>F1</kbd>: toggle voice transmission
- <kbd>Ctrl+L</kbd>: clear chat log
- <kbd>Tab</kbd>: toggle focus between chat and user tree
- <kbd>Page Up</kbd>: scroll chat up
- <kbd>Page Down</kbd>: scroll chat down
- <kbd>Home</kbd>: scroll chat to the top
- <kbd>End</kbd>: scroll chat to the bottom
- <kbd>F10</kbd>: quit

## License

GPLv2

## Contributors

Tim Cooper (<tim.cooper@layeh.com>)
Brandon McGinty-Carroll (<bmmcginty.barnard@bmcginty.us>)
