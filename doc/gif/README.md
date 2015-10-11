# Demo Gif

This directory includes `.gif` animation to show demo of `gcli`.

This demo is created by QuickTime Player, `ffmpeg` and `gifsicle`. After taking new screen recording (`.mov`), run the following command,

```bash
$ ffmpeg -i gcli.mov -s 1000x600 -pix_fmt rgb24 -r 15 -f gif - | gifsicle --optimize=3 --delay=3 > gcli.gif
```

## References

- [OS X Screencast to animated GIF](https://gist.github.com/dergachev/4627207)
