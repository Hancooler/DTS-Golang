package main

import (
	"assigment3/config"
	"assigment3/controller"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
)

func main() {
	// Menjalankan fungsi untuk menghubungkan ke database
	config.Connect()

	// Membuat scheduler untuk memperbarui data setiap 15 detik
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(15).Seconds().Do(controller.UpdateData)
	if err != nil {
		log.Fatalf("failed to schedule task: %v", err)
	}

	// Menangani sinyal SIGINT dan SIGTERM untuk keluar dari program secara graceful
	gracefulExit := make(chan os.Signal, 1)
	signal.Notify(gracefulExit, syscall.SIGINT, syscall.SIGTERM)

	// Memulai scheduler di dalam goroutine
	go func() {
		<-gracefulExit
		scheduler.Stop()
		os.Exit(0)
	}()

	// Menjalankan scheduler
	scheduler.StartBlocking()
}
