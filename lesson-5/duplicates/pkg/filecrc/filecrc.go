// Рассчитывает контрольную сумму файла
package filecrc

import (
	"crypto/md5"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func FileMD5(path string) string {
	log.Debug().Stack().Msg("star func fileMD5")

	h := md5.New()
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by open file")
		return ""
	}

	_, err = io.Copy(h, f)
	if err != nil {
		log.Error().Stack().Err(err).Msg("error by count MD5")
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
