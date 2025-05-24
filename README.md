# 🌤️ GoExpert Weather Services

Este projeto consiste em dois microserviços escritos em Go para a Pós de GOEXPERT da FullCycle

- **Serviço A**: Recebe um `CEP` via POST e envia ao Serviço B.
- **Serviço B**: Orquestração — consulta o ViaCEP, busca a temperatura via WeatherAPI e retorna:
  - Cidade
  - Temperaturas em Celsius, Fahrenheit e Kelvin

Todos os serviços são **instrumentados com OpenTelemetry** e expõem spans via **Zipkin**.

---

## 📦 Requisitos

- [Podman](https://podman.io/) ou Docker
- [Podman Compose](https://github.com/containers/podman-compose) ou Docker Compose
- Chave da API [https://www.weatherapi.com/](https://www.weatherapi.com/)

---

## 🚀 Clonar e executar

```bash
git clone https://github.com/rafael1abrao/goexpert-weather-services.git
cd goexpert-weather-services
```

---

## ⚙️ Configuração

Copie os arquivos `.env.example` para `.env` em cada serviço:

```bash
cp service-a/service-a.env.example service-a/.env
cp service-b/service-b.env.example service-b/.env
```

Edite o arquivo `services/service-b/.env` e informe sua chave da WeatherAPI:

```env e docker-compose.yml
WEATHER_API_KEY=coloque_sua_api_key_aqui
```

---

## 🐳 Executar com Docker Compose

```bash
docker compose up --build
```

ou com Podman Compose:

```bash
podman-compose up --build
```

Os seguintes serviços serão iniciados:

| Serviço     | Porta | Endpoint                        |
|-------------|-------|----------------------------------|
| Serviço A   | 8080  | `POST /input`                   |
| Serviço B   | 8081  | `POST /weather`                 |
| Zipkin      | 9411  | [http://localhost:9411](http://localhost:9411) |

---

## 📮 Exemplo de uso

```bash
curl -X POST http://localhost:8080/input \
  -H "Content-Type: application/json" \
  -d '{"cep":"01001000"}'
```

✅ Resposta esperada:

```json
{
  "city": "São Paulo",
  "temp_C": 26.4,
  "temp_F": 79.52,
  "temp_K": 299.4
}
```

---

## 🔍 Observabilidade

Abra o Zipkin para visualizar os traces:

```
http://localhost:9411
```

Clique em “Find Traces” para explorar os spans entre:

- Serviço A → Serviço B
- Serviço B → ViaCEP
- Serviço B → WeatherAPI

---

## 📂 Estrutura

```bash
.
├── docker-compose.yml
│   ├── service-a/
│   │   ├── .env.example
│   │   ├── main.go, handler/, service/
│   └── service-b/
│       ├── .env.example
│       ├── main.go, handler/, viacep/, weather/
```

---

## 👨‍💻 Autor

Rafael Abrao  
[github.com/rafael1abrao](https://github.com/rafael1abrao)