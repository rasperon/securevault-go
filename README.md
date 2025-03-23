# SecureVault

SecureVault is a zero-knowledge encrypted storage solution that provides secure data encryption with multi-factor authentication.

## 🇬🇧 English

### Features
- AES encryption for secure data storage
- PKCS7 padding implementation
- Multi-factor authentication using OTP
- Command-line interface with interactive prompts
- 16-character encryption key requirement
- Zero-knowledge architecture

### Requirements
- Go 1.23.1 or higher
- Required dependencies (automatically installed via go.mod):
  - github.com/AlecAivazis/survey/v2
  - Other supporting packages

### Installation
```bash
git clone https://github.com/rasperon/securevault-go.git
cd SecureVault
go mod download
```

### Usage
```bash
go run main.go
```

1. Enter your 16-character encryption key when prompted
2. Input the data you want to encrypt
3. The system will show you the encrypted data
4. Data will be decrypted to verify the process
5. Complete the MFA process by entering the generated OTP

### Security Features
- AES encryption with CBC mode
- Secure random OTP generation
- Input validation for encryption keys
- Padding security measures

## 🇹🇷 Türkçe

### Özellikler
- Güvenli veri depolama için AES şifreleme
- PKCS7 dolgu uygulaması
- OTP kullanan çok faktörlü kimlik doğrulama
- Etkileşimli komut satırı arayüzü
- 16 karakterli şifreleme anahtarı gereksinimi
- Sıfır-bilgi mimarisi

### Gereksinimler
- Go 1.23.1 veya üstü
- Gerekli bağımlılıklar (go.mod aracılığıyla otomatik kurulur):
  - github.com/AlecAivazis/survey/v2
  - Diğer destek paketleri

### Kurulum
```bash
git clonehttps://github.com/rasperon/securevault-go
cd SecureVault
go mod download
```

### Kullanım
```bash
go run main.go
```

1. İstendiğinde 16 karakterlik şifreleme anahtarınızı girin
2. Şifrelemek istediğiniz veriyi girin
3. Sistem size şifrelenmiş veriyi gösterecek
4. İşlemin doğruluğunu kontrol etmek için veri şifresi çözülecek
5. Oluşturulan OTP'yi girerek MFA sürecini tamamlayın

### Güvenlik Özellikleri
- CBC modunda AES şifreleme
- Güvenli rastgele OTP oluşturma
- Şifreleme anahtarları için girdi doğrulama
- Dolgu güvenlik önlemleri

## License
MIT License - See [LICENSE](LICENSE) file for details