# gozip
handy tool/lib for compressing strings/files

I use this to compress some static content into strings and use those content dynamically at runtime.

## Usage

compress:

```bash
cat xxx.txt | ./gozip
```

decompress:

```bash
echo <compressed text> | ./gozip -d
```
