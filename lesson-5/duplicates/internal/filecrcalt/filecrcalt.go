// Package filecrcalt рассчитывает контрольную сумму файла для виртуальной файловой системы
package filecrcalt

import (
	"crypto/md5" //nolint:gosec
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
)

func FileMD5(path string, fs afero.Fs) string {
	log.Debug().Stack().Msg("star func fileMD5")

	h := md5.New() //nolint:gosec
	f, err := fs.Open(path)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by open file")
		return ""
	}
	defer func(f afero.File) {
		err = f.Close()
		if err != nil {
			log.Error().Stack().Err(err).Msg("error by close file")
		}
	}(f)

	_, err = io.Copy(h, f)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by count MD5")
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
