# Exifixer

Sometimes, I put pictures received via Whatsapp into my photo collection directory structure.
This tool add date/time EXIF metadata to those images, implied from their file names,
so that [Photoprism](https://github.com/photoprism/photoprism) can sort them correctly in the timeline.

Only files which do not already have date/time EXIF metadata are modified.

Needs exiftool installed (`sudo apt-get install exiftool`).

```
$ go install jo-m.ch/go/exifixer@latest
$ exifixer --help
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
