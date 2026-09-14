# 援疆语文教研资源协同服务

Echo 后端使用 SQLite 记录课例版本、学校订阅和反馈。数据库路径由 `DATABASE_PATH` 环境变量指定，迁移脚本在 `migrations`，默认启动提供健康检查。

运行：`go run .`；测试：`go test ./...`。
