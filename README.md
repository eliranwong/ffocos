# ffocos

'ffocos' stands for 'fix fcitx on Chrome OS'

a small binary to fix fcitx on Chrome OS

# fcitx issue on Chrome OS

Gui apps behave in three different ways with fcitx on Chrome OS:

1. fully-compatible - fcitx works perfectly with some gui apps, e.g. ms vs code
2. partially-comptabile - fcitx does not work with some gui apps when they are opened, without a fcitx-fully-comptabile gui app opened at the same time.  e.g. gnome-terminal
3. incompatible - fcitx does not work at all.

The issue is unique to Chrome OS.  For more information, read:

https://github.com/eliranwong/ChromeOSLinux/blob/main/input_method/fcitx.md#troubleshooting-2-input-selection-panel-is-hidden-sometimes

Some applications does not work with fcitx if they are opened, without other fcitx-comptabile gui apps opened at the same time.

For example, fcitx does not work with 'gnome-terminal' when 'gnome-terminal' is opneded with all other linux gui apps closed.  However, fcitx works perfectly on 'gnome-terminal' when other fcitx-compatible gui apps, such as MS VS code or koncole, are opened.

# Workaround Solution

This project is to create a very light dummy fcitx-fully-compatible gui app 'ffocos'.

This is created to help fcitx working on fcitx-partially-compatible apps, described above.

Simply run 'ffocos' to keep fcitx selection panel working on other fcitx-partially-compatible apps.

