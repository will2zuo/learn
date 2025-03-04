#### url 访问网站的过程

1. dns 解析，将域名解析为 ip 地址

2. tcp 连接，tcp 的三次握手

   - 由浏览器发起，告诉服务器要开始请求数据了 发起 syn
     
   - 由服务器发起，告诉了浏览器准备好了接收数据，可以发起请求 回复 syn 并ACK 客户端的 syn

   - 由浏览器发起，告诉服务器马上发送数据，准备接收，收到服务端的 syn，发送对服务端的 syn 的 ack

3. 发送请求

4. 接收响应

5. 渲染页面

6. 断开连接，tcp 的四次挥手

   - 由浏览器发起，告诉服务器请求报文发送完毕，准备关闭

   - 由服务器发起，告诉浏览器，请求报文接收完毕，准备关闭，你也准备关闭

   - 由服务器发起，告诉浏览器，响应报文发送完毕，你准备关闭

   - 由浏览器发起，告诉服务器，响应报文接收完毕，准备关闭

## TCP 三次握手、四次挥手
### 三次握手
![img.png](img.png)

### 四次挥手
![img_1.png](img_1.png)

#### 常见的 http 状态码

> 1xx：请求被接收
>
> 2xx：成功
>
> 3xx：重定向
>
> 4xx：客户端错误
>
> 5xx：服务器错误

```
301:永久重定向
302:临时重定向
401:权限不足，请求认证用户身份信息
403:拒绝访问
404:无法找到请求的资源
405:客户端请求中的方法被禁止
500:服务器内部错误
501:此请求方法不被服务器支持且无法被处理
502:网关错误
503:服务器繁忙
504:不能及时响应
```

#### osi 的七层模型 和 tcp/ip 四层关系

<table>
   <tr>
        <td>OSI 七层网络模型</td>
        <td>TCP/IP四层概念模型</td>
        <td>对应网络协议</td>
    </tr>
    <tr>
        <td>应用层</td>
        <td rowspan="3">应用层</td>
        <td>HTTP、TFTP, FTP, NFS, WAIS、SMTP</td>
    </tr>
    <tr>
        <td>表示层</td>
        <td>Telnet, Rlogin, SNMP, Gopher</td>
    </tr>
    <tr>
        <td>会话层</td>
        <td>SMTP, DNS</td>
    </tr>
    <tr>
        <td>传输层</td>
        <td>传输层</td>
        <td>TCP, UDP</td>
    </tr>
    <tr>
        <td>网络层</td>
        <td>网络层</td>
        <td>IP, ICMP, ARP, RARP, AKP, UUCP</td>
    </tr>
    <tr>
        <td>数据链路层</td>
        <td rowspan="2">数据链路层</td>
        <td>FDDI, Ethernet, Arpanet, PDN, SLIP, PPP</td>
    </tr>
    <tr>
        <td>物理层</td>
        <td>IEEE 802.1A, IEEE 802.2到IEEE 802.11</td>
    </tr>
</table>

![img_2.png](img_2.png)

#### linux 发送接受网络包流程
![img_3.png](img_3.png)

#### 跨域怎么出现的，怎么解决跨域

**出现**：浏览器的同源策略，限制了一个源的文件或者脚本如何和另一个源的资源进行交互，如果没有同源策略，容易收到 XSS/CSRF 等攻击

**解决**：

1. jsonp，利用 <script> 标签没有跨域限制的漏洞，页面可以动态的得到其他源的 json 数据

   - 优点：兼容性好，可以用于主流浏览器的跨域访问问题
   - 缺点：仅支持 get 方法；不安全，容易遭受 xss 攻击

2. cors 跨域资源共享，分为简单请求和复杂请求

   - 简单请求
   - 复杂请求，在正式请求之前，增加一次 http 查询请求

3. nginx 的反向代理

   配置一个代理服务器做跳板机

4. node 中间件代理（两次跨域）

   实现原理：就是服务器向服务器请求


### get 和 post 的区别
- 数据传输大小： get 传输数据的大小是 2kb，而 post 一般是没有限制的，但是会受内存大小影响，一般通过修改 php.ini 配置文件来修改
- 数据传输方式： get 是通过 url 传递参数的，在 url 中可以看到参数；post 是在表单中使用 post 方法提交
- 数据安全性：get 参数可见，容易被攻击
- 缓存： get 可以被缓存， post 不能被缓存

### HTTP/1.1、HTTP/2 和 HTTP/3 对比表

| **特性**               | HTTP/1.1                  | HTTP/2                    | HTTP/3                    |
|-------------------------|---------------------------|---------------------------|---------------------------|
| **传输层协议**          | TCP                       | TCP                       | QUIC (基于 UDP)           |
| **数据传输方式**        | 文本格式                  | 二进制帧                  | 二进制帧                  |
| **多路复用**            | ❌ 不支持（管道化有缺陷） | ✅ 支持（同一 TCP 连接）  | ✅ 支持（独立 QUIC 流）    |
| **队头阻塞**            | ⚠️ 应用层和传输层均存在   | ⚠️ 仅传输层（TCP 丢包）   | ✅ 彻底解决                |
| **头部压缩**            | ❌ 无                     | ✅ HPACK 压缩             | ✅ QPACK 压缩              |
| **服务器推送**          | ❌ 无                     | ✅ 支持                   | ✅ 支持                   |
| **握手延迟**            | ⏳ 高（TCP + TLS 1-2 RTT）| ⏳ 高（同 HTTP/1.1）      | ⚡ 低（0-1 RTT，支持 0-RTT）|
| **移动端优化**          | ❌ 无                     | ❌ 无                     | ✅ 连接迁移、抗丢包       |
| **兼容性**              | 🌍 全平台兼容             | 🌐 主流现代浏览器/服务器   | 🚧 逐步普及（需支持 QUIC） |

### 关键说明
- **多路复用**：HTTP/2 在单 TCP 连接上并行传输，HTTP/3 通过 QUIC 流彻底消除队头阻塞。
- **队头阻塞**：HTTP/1.1 因顺序处理请求而阻塞；HTTP/2 仅因 TCP 丢包阻塞；HTTP/3 无阻塞。
- **握手延迟**：HTTP/3 的 0-RTT 需已建立过连接，首次连接仍需 1-RTT。
- **移动端优化**：HTTP/3 支持 IP 切换不断连（如 Wi-Fi 切 5G），且抗弱网能力更强。

### tcp 如何解决粘包的问题

#### 1. 消息定长法
- ​**原理**​  
  固定每个消息的长度（如每个包 100 字节），不足部分填充空值。

- ​**代码示例**
  ```python
  # 发送端：填充定长
  message = "data".ljust(100, '\0')
  socket.send(message)
  
  # 接收端：按定长读取
  while True:
      chunk = socket.recv(100)
      if not chunk: break
      # 处理 chunk（需去除填充的空值）
  ```

- ​**优点**​  
  ✅ 实现简单
- ​**缺点**​  
  ❌ 浪费带宽  
  ❌ 灵活性差

---

## 2. 分隔符法
- ​**原理**​  
  在消息结尾添加特殊分隔符（如 `\n` 或自定义符号）。

- ​**代码示例**
  ```python
  # 发送端：添加分隔符
  message = "data|"
  socket.send(message.encode())
  
  # 接收端：按分隔符拆分
  buffer = ""
  while True:
      data = socket.recv(1024).decode()
      if not data: break
      buffer += data
      while "|" in buffer:
          msg, buffer = buffer.split("|", 1)
          # 处理 msg
  ```

- ​**优点**​  
  ✅ 灵活，适合文本协议
- ​**缺点**​  
  ❌ 需转义分隔符  
  ❌ 性能较低（需遍历字符）

---

## 3. 长度前缀法（推荐）
- ​**原理**​  
  在消息头部添加长度字段（如 4 字节表示数据长度）。

- ​**代码示例**
  ```python
  # 发送端：添加长度前缀
  data = "payload"
  length = len(data).to_bytes(4, byteorder='big')  # 4字节大端序
  socket.send(length + data.encode())
  
  # 接收端：先读长度，再读数据
  buffer = bytearray()
  while True:
      chunk = socket.recv(4096)
      if not chunk: break
      buffer.extend(chunk)
      while len(buffer) >= 4:
          length = int.from_bytes(buffer[:4], byteorder='big')
          if len(buffer) < 4 + length: break
          msg = buffer[4:4+length].decode()
          buffer = buffer[4+length:]
          # 处理 msg
  ```

- ​**优点**​  
  ✅ 高效，适合二进制协议  
  ✅ 无数据冗余或转义问题
- ​**缺点**​  
  ❌ 需处理字节序和长度溢出

---

## 4. 协议封装法
- ​**原理**​  
  使用现成的应用层协议（如 HTTP、Protobuf）封装数据。

- ​**示例协议**​
   - ​**HTTP**：通过 `Content-Length` 或分块传输标识长度
   - ​**Protobuf**：序列化消息自带长度前缀

---

## 5. 高级框架支持
- ​**Netty (Java)**​  
  通过 `LengthFieldBasedFrameDecoder` 自动处理粘包。

- ​**Go 的 `bufio.Scanner`**​
  ```go
  scanner := bufio.NewScanner(conn)
  scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
      // 自定义拆分逻辑（如按长度前缀）
  })
  ```

---

## 方法对比表
| 方法         | 适用场景                          | 性能  | 复杂度 |
|--------------|----------------------------------|-------|--------|
| 消息定长法   | 固定长度协议（如传感器数据）      | 高    | 低     |
| 分隔符法     | 文本协议（如日志、命令行交互）    | 中    | 中     |
| 长度前缀法   | 二进制协议（如游戏、金融数据）    | 高    | 中     |
| 协议封装法   | 标准化或跨平台通信                | 中    | 低     |

---

## 总结建议
1. ​**优先选择长度前缀法**：适合高性能二进制协议。
2. ​**文本协议可用分隔符法**：如日志流处理。
3. ​**避免重复造轮子**：直接使用成熟协议（如 Protobuf/HTTP）。
