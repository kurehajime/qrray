# qrray
Create HTML with a QR codes from multiple lines of text.
![qrray](https://user-images.githubusercontent.com/4569916/183288217-4bd5590e-c4d5-48ad-aacd-ac99988d5760.png) 

## install 

[Download here](https://github.com/kurehajime/qrray/releases)

## Usage

1. Prepare multiple lines of text.

```example.txt
https://www.google.com
https://www.apple.com
https://www.facebook.com
https://www.amazon.com
```

2. Execute

```
$ qrray example.txt
```

3. Done.

By default, `./qr/index.html` will be created.

![result](https://user-images.githubusercontent.com/4569916/183292804-bff6fc2e-f63d-4bb9-ad76-a4e9eb546486.png)

## Option

Option of qrray

```
  -out string
        out put path (default "./qr")
  -prefix string
        file name prefix (default "qr_")
  -size int
        size of qr code (default 128)
```
