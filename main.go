package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/jordan-wright/email"
    "net/smtp"
)

type TimeTracker struct {
    StartTime   time.Time
    SocialTime  time.Duration
    WebTime     time.Duration
    IsTracking  bool
}

var tracker TimeTracker

func main() {
    router := gin.Default()

    // Endpoint per avviare il tracciamento
    router.POST("/start", startTracking)

    // Endpoint per fermare il tracciamento
    router.POST("/stop", stopTracking)

    // Endpoint per inviare il report via email
    router.POST("/sendReport", sendReport)

    router.Run(":8080")
}

func startTracking(c *gin.Context) {
    tracker.StartTime = time.Now()
    tracker.IsTracking = true
    c.JSON(http.StatusOK, gin.H{"message": "Tracking started!"})
}

func stopTracking(c *gin.Context) {
    if tracker.IsTracking {
        duration := time.Since(tracker.StartTime)

        // Aggiorna i tempi separati per Social e Web (semplificato)
        tracker.SocialTime += duration / 2  // Social Media
        tracker.WebTime += duration / 2     // Web

        tracker.IsTracking = false
        c.JSON(http.StatusOK, gin.H{
            "message": "Tracking stopped!",
            "SocialTime": tracker.SocialTime.String(),
            "WebTime": tracker.WebTime.String(),
        })
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No active tracking session"})
    }
}

func sendReport(c *gin.Context) {
    // Imposta il contenuto del report
    report := fmt.Sprintf("Time spent on Social: %s\nTime spent on Web: %s\n",
        tracker.SocialTime.String(), tracker.WebTime.String())

    // Invia il report via email
    err := sendEmail("emanuele.zuffranieri@gmail.com", report)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Report sent via email"})
}

func sendEmail(to string, body string) error {
    e := email.NewEmail()
    e.From = "YourApp <youremail@example.com>"
    e.To = []string{to}
    e.Subject = "Daily Report"
    e.Text = []byte(body)
    // Imposta le credenziali SMTP (usando Gmail come esempio)
    return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "emanuele.zuffranieri@gmail.com", "pesciolina8183@", "smtp.gmail.com"))
}
