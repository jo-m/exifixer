package main

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/barasher/go-exiftool"
	"github.com/jo-m/exifixer/internal/pkg/logging"
	"github.com/rs/zerolog/log"
)

type flags struct {
	logging.LogConfig

	DryRun bool   `arg:"-d,--dry-run" default:"false" help:"do not change any files"`
	Dir    string `arg:"positional,required"  help:"directory to process" placeholder:"DIR"`
}

func (flags) Description() string {
	return `Exifixer finds media files with missing EXIF tags,
	tries to imply the timestamp from the file name,
	and stores it in the file as EXIF tag.`
}

func getExt(fname string) (string, error) {
	ext := filepath.Ext(fname)
	if len(ext) == 0 {
		return "", errors.New("file name has no extension")
	}
	ext = strings.ToLower(ext[1:])
	if ext == "jpeg" {
		ext = "jpg"
	}
	return ext, nil
}

type file struct {
	path    string
	relPath string // For user display.
	ts      time.Time
	ext     string
}

func handleFile(f file, et *exiftool.Exiftool, dryRun bool) {
	log := log.With().Str("path", f.relPath).Str("ext", f.ext).Logger()
	log.Debug().Msg("handling file")

	// TODO: Support more.
	if f.ext != "jpg" {
		log.Debug().Msg("not supported, skipping")
		return
	}

	meta := et.ExtractMetadata(f.path)[0]
	hasDateTime := false
	for k := range meta.Fields {
		if strings.Contains(k, "DateTime") {
			hasDateTime = true
			val, err := meta.GetString(k)
			if err != nil {
				panic(err)
			}
			meta.GetString(k)
			log.Trace().Str("k", k).Str("v", val).Msg("date field")
		}
	}

	if hasDateTime {
		log.Debug().Msg("file already has datetime metadata, skipping")
		return
	}

	ts := f.ts.Format("2006:01:02 15:04:05")
	ofs := f.ts.Format("-07:00")
	if dryRun {
		log.Warn().Str("DateTimeOriginal", ts).Str("OffsetTimeOriginal", ofs).Msg("would write to file (dry run)")
		return
	}

	log.Warn().Str("DateTimeOriginal", ts).Str("OffsetTimeOriginal", ofs).Msg("write to file")
	// TODO: actually write.

}

func main() {
	f := flags{}
	arg.MustParse(&f)
	logging.MustInit(f.LogConfig)
	log.Debug().Interface("flags", f).Msg("flags")

	// Collect files.
	files := make(chan file)
	dir, err := filepath.Abs(f.Dir)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	go func() {
		filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			fname := filepath.Base(path)
			ts := tsFromFileName(fname)
			if ts == nil {
				return nil
			}
			ext, err := getExt(fname)
			if err != nil {
				log.Panic().Str("fname", fname).Msg("file has no extension")
			}

			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				log.Panic().Err(err).Send()
			}

			files <- file{
				path:    path,
				relPath: relPath,
				ts:      *ts,
				ext:     ext,
			}

			return nil
		})

		close(files)
	}()

	// And handle them.
	et, err := exiftool.NewExiftool()
	if err != nil {
		log.Panic().Err(err).Msg("Is exiftool installed? sudo apt-get install exiftool")
	}
	defer et.Close()
	for file := range files {
		handleFile(file, et, f.DryRun)
	}
}
