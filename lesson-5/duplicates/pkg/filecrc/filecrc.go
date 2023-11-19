// Package filecrc рассчитывает контрольную сумму файла
package filecrc

import (
	"crypto/md5" //nolint:gosec
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func FileMD5(path string) string {
	log.Debug().Stack().Msg("star func fileMD5")

	h := md5.New() //nolint:gosec
	f, err := os.Open(path)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by open file")
		return ""
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Error().Stack().Err(err).Msg("error by open file")
		}
	}(f)

	_, err = io.Copy(h, f)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by count MD5")
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
