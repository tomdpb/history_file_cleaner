# history_cleaner

This simple command line program will remove lines from a file given regex expressions in a separate file.
It's main intent is to be used with zhistory or bash_history and remove lines with `cd`, and variations of `ls`, though any valid regex can be used.
Included in this repository is an example file containing regex patters which would be removed from a file. A backup of the file is made in case the user has any regrets or makes a mistake.

To pass a file to the program, call the compiled binary, pass the file name as an argument, and finally pass a file containing the regex patterns like so:
`history_cleaner --file=~/.zhistory --regexFile=regex_patterns.txt`

The program will will use `regex_patterns.txt` (or whatever was passed to it) to determine which regular expressions to delete from the file passed to it.

Finally, setting `--verbose=true` will also show which lines are being removed, but it is set to `false` by default.

# Warning

When setting regular expressions, make sure starting and ending characters are set correctly.

Setting `cd*` might accidentally trigger `chmod`, which is not usually the intended behaviour.

## Quick hints to ensure intended patterns

- `^char` indicates that the pattern _starts_ with `char`.
- `char.*` indicates that the pattern has exactly `char` somewhere, and any number of extra characters after it.
- `cd*` indicates that the pattern has `c` then any number of characters, and has a `d` afterwards. Warning: `chmod` matches this!
- `char$` indicates that the pattern _ends_ with `char`.
