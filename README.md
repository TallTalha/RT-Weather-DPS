---

# Gerçek Zamanlı Hava Durumu Veri İşleme Sistemi

## Türkçe Açıklama

### Genel Bakış
Bu proje, WebSocket kullanarak hava durumu verilerini gerçek zamanlı olarak yakalayan, verileri gRPC servisi üzerinden işleyip kaydeden, MongoDB'de saklayan ve aynı zamanda bir RabbitMQ kuyruğuna gönderen bir gerçek zamanlı hava durumu veri işleme sistemidir. WebSocket, gRPC, MongoDB ve RabbitMQ gibi farklı teknolojilerin tek bir uygulamada nasıl entegre edilebileceğini gösterir ve eğitim amaçlarına uygun olduğu gibi daha büyük ölçekli uygulamalar için bir prototip olarak da kullanılabilir.

### Kullanılan Teknolojiler
- **WebSocket**: Gerçek zamanlı veri alma.
- **gRPC**: Verimli sunucu-istemci iletişimi.
- **MongoDB**: Veri depolama.
- **RabbitMQ**: Mesaj kuyruklama.

### Nasıl Çalıştırılır
1. Repoyu klonlayın:
   ```
   git clone <repository-url>
   ```
2. Gerekli bağımlılıkları yükleyin:
   ```
   // Gerekli kütüphane ve araçların kurulum komutları eklenecek
   ```
3. MongoDB ve RabbitMQ servislerini başlatın.
4. Sunucuyu çalıştırın:
   ```
   go run server.go
   ```
5. İstemciyi çalıştırın:
   ```
   go run client.go
   ```

### Katkıda Bulunma
Katkılarınız bekleniyor! Lütfen repoyu forklayın ve iyileştirmelerinizi içeren bir pull request gönderin.

---

# Real-Time Weather Data Processing System

## English Description

### Overview
This project is a real-time weather data processing system that captures weather conditions using WebSocket, processes and saves the data via a gRPC service, stores it in MongoDB, and also pushes it to a RabbitMQ queue. It demonstrates the integration of different technologies such as WebSocket, gRPC, MongoDB, and RabbitMQ in a single application, suitable for educational purposes and as a prototype for larger scale implementations.

### Technologies Used
- **WebSocket**: For real-time data fetching.
- **gRPC**: For efficient server-client communication.
- **MongoDB**: For data storage.
- **RabbitMQ**: For message queuing.

### How to Run
1. Clone the repository:
   ```
   git clone <repository-url>
   ```
2. Install dependencies:
   ```
   // Commands will add to install necessary libraries and tools
   ```
3. Start the MongoDB and RabbitMQ services.
4. Run the server:
   ```
   go run server.go
   ```
5. Run the client:
   ```
   go run client.go
   ```

### Contributions
Contributions are welcome! Please fork the repository and submit a pull request with your enhancements.

---
