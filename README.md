thymer
======

Console Pomodoro Timer written in golang. It prints a countdown with progress bar.

Best fit using conky or your terminal.

Installation and Startup
-----

    go get github.com/blang/thymer
    $GOPATH/bin/thymer -duration="25m" -interval="1s" -bar=30  

Usage
-----

    Usage of ./thymer:
    -bar=20: Length of progressbar
    -duration=25m0s: Pomodoro duration
    -interval=1s: Update interval



My Conky
-----

    {"full_text": " Thymer ${exec tail -n 1 ~/.thymerlog} ", "color":"\#268BD2"},
    {"full_text": " Doros: ${exec cat ~/.thymerlog | grep "Stopped" | wc -l} / ${exec cat ~/.thymerlog | grep "Interrupted" | wc -l}", "color":"\#268BD2"},

Start - Stop 
-----

    Primitive Stop: killall thymer
    Start: $HOME/bin/thymer -duration="25m" -bar=30 >> $HOME/.thymerlog 2>&1

Workflow
-----

- Start thymer in background (in my case shortcut on i3)
- tail the last log line to print current pomodoro status
- Use grep to determine completed and interrupted pomodoros
- Stop thymer via killall
- Clear log every day

Contribution
-----

Feel free to make a pull request. For bigger changes create a issue first to discuss about it.


License
-----

See [LICENSE](LICENSE) file.
