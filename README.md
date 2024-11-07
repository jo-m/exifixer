# Exifixer

Find JPEG files which are named like Whatsapp images,
and add an EXIF timestamp tag to them if they do not already have one.

Needs exiftool installed (`sudo apt-get install exiftool`).

```
$ go install github.com/jo-m/exifixer@latest
$ ./exifixer --help
Exifixer finds media files with missing EXIF tags,
        tries to imply the timestamp from the file name,
        and stores it in the file as EXIF tag.
Usage: exifixer [--log-pretty] [--log-level LEVEL] [--dry-run] DIR

Positional arguments:
  DIR                    directory to process

Options:
  --log-pretty           log pretty [default: true, env: LOG_PRETTY]
  --log-level LEVEL      log level [default: info, env: LOG_LEVEL]
  --dry-run, -d          do not change any files [default: false]
  --help, -h             display this help and exit
```
