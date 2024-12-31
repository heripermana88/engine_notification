package main

import (
	"log"
	"time"

	"decron/jobs"

	"github.com/go-co-op/gocron"
)

func main() {
	// Inisialisasi scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	// Menambahkan tugas ke scheduler
	scheduler.Every(2).Second().Do(jobs.SendReminderEmail)
	// scheduler.Every(10).Seconds().Do(jobs.CleanupTempFiles)

	// Mulai scheduler
	scheduler.StartAsync()

	// Log untuk memastikan aplikasi tetap berjalan
	log.Println("Scheduler is running... Press Ctrl+C to exit.")
	select {} // Agar aplikasi tetap berjalan
}
