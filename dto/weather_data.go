package dto

// WeatherData DTO, WebSocket'ten alınan hava durumu verilerini temsil eder
type WeatherData struct {
	Temperature float64 `json:"temperature"` // Sıcaklık değeri
	City        string  `json:"city"`        // Şehir adı
	Timestamp   string  `json:"timestamp"`   // Veri alındığı zaman (ISO 8601 formatında)
}
