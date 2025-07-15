# 🛰️ Kafka ZooKeeper Cluster Example with Producer & Consumer

A minimal example of a Kafka setup **with ZooKeeper** with:

- ✅ Basic Go **Producer & Consumer**
- ✅ Kafka UI dashboard
- ✅ Docker Compose orchestration

---

## 🚀 Getting Started

### 1. Start Kafka Cluster
```bash
docker-compose up -d
```
### 2. Access Kafka UI

Visit: http://localhost:8080

You should see Kafka broker registered.
🧪 Test Producer & Consumer
Run Producer and Consumer
```bash
go run main.go
```
#### Default topic
example-topic is used by default. It will be auto-created if not exists.
