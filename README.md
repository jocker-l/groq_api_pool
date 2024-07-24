# Groq api key pool

---

使用Groq API Key Pool，可以方便地管理和使用多个API Key，避免因API Key过期或使用量限制导致限速的问题。

## env

1. SERVER_HOST 服务域名
2. OpenAuthSecret 是否开启auth校验
3. IsVercel 是否部署在vercel上
4. CHINESE_PROMPT 中文提示词
5. SUPPORT_APIKEY 支持apikey调用
6. Authorization 访问密钥
7. SERVER_PORT 服务端口
8. BASE_URL 基础请求地址
9. AuthSecret 鉴权密钥
10. http_proxy http代理

## 使用方法

1. 在.env文件中配置相关环境变量
2. session_tokens.txt 文件中添加API Key，每行一个
3. 启动服务，即可使用多个API Key进行调用

## docker打包

```bash
docker build -t groq-api-key-pool:latest .
```

## docker启动

```bash
docker run -d -p 8080:8080 --name groq-api-key-pool -v /path/to/session_tokens.txt:/app/session_tokens.txt groq-api-key-pool
```

## 注意事项

1. 请确保session_tokens.txt文件中的API Key是有效的，并且有足够的调用次数
2. 请勿将API Key泄露给他人，以免造成损失