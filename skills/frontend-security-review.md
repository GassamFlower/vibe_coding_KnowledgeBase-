---
name: frontend-security-review
version: 1.0
category: 前端安全
---

# Skill: frontend-security-review（前端安全审查）

> 来源：Vibe Coding 脚手架化 + 后端安全审查（backend-security-review）的前端对偶
> 适用：纯前端 / 小程序 / 前后端分离项目的前端安全审查（后端项目走 backend-security-review）

## 何时使用
- 纯前端 / 小程序项目走到业务门的安全审查阶段
- 前后端分离项目的前端部分安全收尾
- 用户说"检查一下前端安全""有没有 XSS"

## 解决什么问题
安全审查不能只查后端。前端有自己的攻击面：XSS、CSRF、敏感信息泄露、依赖漏洞、明文传输。纯前端项目没有后端可查时，这几条就是全部安全边界，跳过等于没做安全审查。

---

## 一、XSS（跨站脚本）

### 检查内容
用户输入会不会被当成 HTML / 脚本执行。

### 规则
- 渲染用户内容用框架默认转义机制（Vue `{{ }}` / React JSX），不用 `v-html` / `dangerouslySetInnerHTML` / `innerHTML` 除非确认内容可信
- 必须用 `v-html` / `innerHTML` 时，先做白名单过滤（如 DOMPurify）
- URL 跳转参数（`redirect=`）做白名单校验，防 `javascript:` 协议注入

### Prompt
> 请检查前端所有渲染用户输入的位置。逐项说明：哪些用了 v-html / dangerouslySetInnerHTML / innerHTML，内容是否来自用户输入，有没有做白名单过滤；URL 跳转参数有没有防 javascript: 协议注入。

---

## 二、CSRF（跨站请求伪造）

### 检查内容
状态变更请求有没有被跨站伪造。

### 规则
- 状态变更请求（增删改）用 POST，且带 CSRF Token 或同源校验
- Cookie 设置 `SameSite=Lax/Strict`、`HttpOnly`
- 用 token 认证（`Authorization` header）的方案天然免疫 CSRF，确认 token 不放 Cookie

### Prompt
> 请说明前端状态变更请求（增删改）如何防 CSRF：用的是 CSRF Token、SameSite Cookie，还是 token 认证（Authorization header）？token 是否误放在 Cookie 里？

---

## 三、敏感信息与明文传输

### 检查内容
前端有没有泄露敏感信息、有没有明文传输。

### 规则
- 前端代码不硬编码密钥 / token / 接口签名
- 敏感数据不进 localStorage / sessionStorage（token 尽量内存态或 HttpOnly Cookie）
- 生产环境全站 HTTPS，接口不混用 http
- 日志 / 错误上报不打印 token、手机号、身份证号等敏感字段

### Prompt
> 请检查前端是否硬编码了密钥 / token，是否把敏感数据存进 localStorage，接口是否全部走 HTTPS，console.log / 错误上报是否打印了 token、手机号等敏感信息。

---

## 四、依赖与构建安全

### 检查内容
第三方依赖漏洞 + 构建产物。

### 规则
- 依赖版本查已知漏洞（`npm audit` / `yarn audit` / `pnpm audit`）
- 构建产物不带 source map 到生产（避免泄露源码）
- 关键依赖（富文本、图表、上传、加密）标出是否仍在维护

### Prompt
> 请列出前端依赖及版本，逐项说明是否有已知漏洞；构建产物是否包含 .map 源码映射文件并部署到生产；富文本 / 图表 / 上传 / 加密等关键依赖是否仍在维护。

---

## 五、验收测试

### Prompt
> 请对前端做安全验收，输出：
> 1. XSS：往输入框 / URL 参数塞 `<script>` 或 `javascript:` 的结果
> 2. CSRF：跨站发起状态变更请求是否被拦
> 3. 敏感信息：源码 / network 面板里能不能找到 token、密钥、敏感字段
> 4. 依赖：npm audit 结果
> 没实际验证的项目标记"未验证"。

---

## 反模式
- ❌ "前端没有安全问题"（太笼统）
- ❌ 用了 v-html / innerHTML 渲染用户输入却没过滤
- ❌ token / 密钥硬编码在前端代码或存 localStorage
- ❌ 生产接口走 http 明文
- ❌ 构建产物带 source map 部署到生产
