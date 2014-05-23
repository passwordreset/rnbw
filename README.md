# rnbw - A rainbow for your terminal

![screenshot](/screenshot.png?raw=true)

rnbw is nothing more than a pipe, where text goes in, and rainbows come out.
Chances are, it may break something in the process of doing this.

## So why not lolcat?
rnbw's faster. Don't believe me? Here's a completely impartial and scientific
benchmark:

    $ time ls -al | ./rnbw

    real    0m0.007s
    user    0m0.003s
    sys 0m0.004s

    $ time ls -al | lolcat

    real    0m0.070s
    user    0m0.062s
    sys 0m0.010s

rnbw won't stop dumping you colours if your tty doesn't appear to support them
by default, has no dependencies and if printed out with a 14 point typeface
would fit on an A4 page.
