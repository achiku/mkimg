# mkimg

Create image with text from CLI

### Installation

```
go get -u github.com/achiku/mkimg
```

### How to use

```
Usage of mkimg:
  -fontfile string
        ttf font file path
  -fontsize float
        font size (default 100)
  -height int
        image height (default 630)
  -outfile string
        output image path (default "out.png")
  -txt string
        text to image (default "hello, world!")
  -width int
        image width (default 1200)
```

```
mkimg -txt="それは本当にそう" -fontfile=./Koruri-Bold.ttf
```

![img](./img/shs.png)
