# MicroServicesLogic Json Catfact API

Bu,Basit bir MicroServices JSON mimarisidir.

## Kurulum

1. Depoyu klonlayın:

   ```bash
   git clone https://github.com/karalakrepp/MicroServicesLogic.git
2. Proje dizinine gidin: 
    ```bash 
   cd MicroServicesLogic

3. Bağımlılıkları yükleyin: 
   ```bash 
   go mod download

## Kullanım

1. Sunucuyu başlatın:
   ```bash  
   go run main.go


Sunucu varsayılan olarak 8000 numaralı portta çalışacaktır. PORT ortam değişkenini ayarlayarak özel bir port belirleyebilirsiniz.

2. API uç noktalarıyla etkileşime geçin:
- GET /catfact: kedi bilgilerini getirir.

Example Response:
       ```bash  
    {"fact": "Cats can make over 100 different sounds!"}
  


- Bu uç noktalara HTTP istekleri yapmak için cURL veya Postman gibi araçları kullanabilirsiniz.
