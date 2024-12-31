package jobs

import (
	"log"
	"time"
)

// Fungsi untuk mengirimkan email pengingat
func SendReminderEmail() {
	log.Println("Sending reminder email...")
	// Simulasi logika pengiriman email
	time.Sleep(2 * time.Second)
	log.Println("Reminder email sent successfully!")
}

// Fungsi untuk membersihkan file sementara
func CleanupTempFiles() {
	log.Println("Cleaning up temporary files...")
	// Simulasi logika pembersihan file
	time.Sleep(1 * time.Second)
	log.Println("Temporary files cleaned up successfully!")
}
