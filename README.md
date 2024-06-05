# Terminal-Hook

> As the name suggests, hooks for your terminal.

I know there are some `cd` tools that remember your path(s), but I don't like them.
Because no solution fits my style, I created this terminal-hook that lets you easily create hooks to jump around with.

## How it works

Terminal-Hook (**th**) has multiple functionalities, here is a small overview:

- cd: You can use '**th**' as a replacement for cd
  - It will remember the _last_ visited path, so it lets you easily jump to your starting point
- hook: You can create a named **hook**, this will remember the path so you can easily navigate your file system

### cd

You can use th as an replacement for cd but with smarter features.
I already have some little bash script in use for this that remembers my _starting point_ when changing directory, this way I can easily jump back and forth without the nasty `cd ../../../..` stuff.

The command looks like the following: `th the/path/to/visit` and to come back simply type `th`. This will put you right where you started.

### hook

The script lets you create aliases, meaning, you can say `th --hook {hook-name}` and it will remember the current path.

To activate a hook, simply type `th {hook-name}` or to get a list use `th -l | th --list` .
This will spawn an interactive list where you can select your destination.

If you want to delete a hook, simply run `th -d | th --delete`.
