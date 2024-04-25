package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// WeatherData MongoDB'de saklanacak hava durumu verisini temsil eden yapı
type WeatherData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // MongoDB için özgün kimlik
	Temperature float64            `bson:"temperature"`   // Sıcaklık değeri
	City        string             `bson:"city"`          // Şehir adı
	Timestamp   primitive.DateTime `bson:"timestamp"`     // Veri kaydedildiği zaman
}
