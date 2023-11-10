package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	doHardJob()
	if _, err := w.Write([]byte("Hard job done!")); err != nil {
		log.Printf("can't response for TestEndpoint request: %s", err)
		w.WriteHeader(520)
		return
	}
	w.WriteHeader(200)
}

func doHardJob() {
	timeRnd := rand.Intn(30) + 1
	log.Print("job time:", timeRnd)
	timeForHardWork := time.Duration(timeRnd) * time.Second
	time.Sleep(timeForHardWork)
}

func main() {
	// максимальное время для "грациозного" завершения
	maxGracefulShutdownTime := 15 * time.Second

	// создаём новый сервер при помощи gorilla mux
	// и публикуем эндпоит (см. глоссарий)
	router := mux.NewRouter()
	router.HandleFunc("/test", TestEndpoint).Methods("GET")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// запускаем сервер в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen error: %s\n", err)
		}
	}()

	// создаём канал для перехвата сигналов ОС
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Print("server started")
	<-done

	// код ниже срабатывает после полученного сигнала завершения
	// делаем контекст завершения, в котором указано максимально время которое
	// мы готовы потратить на то чтобы запрос пользователя был обработан
	ctx, cancel := context.WithTimeout(context.Background(), maxGracefulShutdownTime)
	defer cancel()
	log.Print("server stopped")

	if err := srv.Shutdown(ctx); err != nil {
		// если выключение сервера не смогло отработать в заданном контексте (в данном случае по времени)
		// то запросы пользователей будет прерваны
		log.Printf("graceful shutdown failed: %s", err)
	} else {
		// если все запросы пользователей успели отработать то мы оказываемся в этом блоке кода
		log.Print("server exited properly")
	}
}
