"# microservices" 
go mod init "go project name"

# Ghi chú về MongoDB trong Golang

## 🔹 1. Các kiểu `bson` trong Golang

| Kiểu      | Đặc điểm | Khi nào nên dùng? |
|-----------|---------|------------------|
| `bson.M` | Map không có thứ tự, dễ viết | Khi không quan trọng thứ tự trường |
| `bson.D` | Slice có thứ tự | Khi thứ tự trường quan trọng (hiếm gặp) |
| `bson.A` | Slice array | Khi làm việc với danh sách giá trị |

Ví dụ:
```go
filterM := bson.M{"age": bson.M{"$gte": 30}}
filterD := bson.D{{"age", bson.D{{"$gte", 30}}}}
filter := bson.M{"tags": bson.M{"$all": bson.A{"golang", "mongodb"}}}