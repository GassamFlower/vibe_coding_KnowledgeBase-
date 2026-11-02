# VibeVibe 进阶篇索引

> 用途：把 vibevibe.cn「进阶篇」16 章整理成逐章摘要，覆盖每章的核心方法、关键做法、子页要点与适用阶段。
> 来源：https://www.vibevibe.cn/Advanced/
> 许可：VibeVibe 使用 CC BY-NC-SA 4.0；本文件只写摘要、要点与原文链接，不复制正文。
> 全文缓存：`.audit/crawl/text/`（本机目录，已被 `.gitignore` 忽略，不入库）。
> 整理日期：2026-11-02。
> 边界：本文件写内容细节；章节到阶段、门禁、Skill 的总对照见 [VibeVibe工程手册索引](VibeVibe工程手册索引.md)。

## 怎么用

1. 先按当前阶段找章节，读「核心」判断是否需要细读。
2. 「关键做法」是整章最值得照做的条目，含原文给出的数字、命令与判断标准。
3. 「子页」给出每节的具体做法；需要完整步骤时打开章节页，从页尾的小节导航进入对应子页。
4. 同一主题在 [VibeVibe优质文章索引](VibeVibe优质文章索引.md) 或 [VibeVibe实践案例索引](VibeVibe实践案例索引.md) 中有更细的专题或复盘时，先读专题，本索引保留章节骨架。

---

## 第 1 章 环境搭建与代码运行基础

- 章节页：<https://www.vibevibe.cn/Advanced/01-environment-setup/>
- 核心：把空白电脑配到能跑起自己的网页项目：装好 Git、Node、包管理器、命令行 AI 工具与国内模型，用脚手架建项目，再用一条命令启动并在浏览器查看。
- 关键做法：
  - 顺序为 Git → Node 24.13.0 → 配 npmmirror 镜像装 pnpm → 装 Claude Code；国内模型用 `npx @z_ai/coding-helper` 填 API Key。
  - 完整路径每一级都不能含中文与空格，否则出现 ENOENT、MODULE_NOT_FOUND；项目名用小写字母与连字符。
  - 用 `pnpm create next-app@latest my-app` 建项目，TypeScript、ESLint、src 目录、App Router 都选 Yes。
  - pnpm 用硬链接共享依赖，10 个项目约 1.5GB，npm 约 5GB；`pnpm-lock.yaml` 要提交 Git。
  - 启动用 `pnpm install` 与 `pnpm dev`，地址是 localhost:3000；端口被占时用 `pnpm dev -- -p 3001`，Ctrl+C 停止。
- 子页：

| 子页 | 要点 |
|---|---|
| 1.0 快速开始 | Windows 装 Git 2.52.0 与 Node 24.13.0，配镜像后装 pnpm 与 Claude Code，再接入 GLM；GLM 价格约为官方 API 的 0.1 折。 |
| 1.1 代码格式演变 | 单文件 HTML 超过 200 行难维护，依次走向文件分离、模块化与 TypeScript 工程化；.ts 需要编译，靠 pnpm dev 运行。 |
| 1.2 认识技术构成 | 分前端、后端、数据与基础设施四层；本教程使用 Next.js + TypeScript + Tailwind + Drizzle + PostgreSQL，看 package.json 可识别项目组成。 |
| 1.3 浏览器与服务器基础 | 浏览器只懂 HTML、CSS、JavaScript，Node.js 负责编译与构建；密钥只能放服务端。 |
| 1.4 终端入门 | 掌握 pwd、ls、cd、mkdir 与 Tab、Ctrl+C，理解 PATH；命令找不到时按拼写、是否安装、当前目录、PATH 四项排查。 |
| 1.5 包管理与项目配置 | nvm 管 Node 版本，pnpm 比 npm 快 2-3 倍；`pnpm-lock.yaml` 自动生成，不可手改。 |
| 1.6 模型与工具 | 模型决定能力上限，工具负责读写文件与执行命令；最低配置是 Claude Code + GLM + 终端 + 浏览器。 |
| 1.7 创建项目 | 脚手架生成 src/app、components、lib、public；页面放 app，组件放 components。 |
| 1.8 Localhost 与端口 | localhost 即 127.0.0.1；0-1023 为系统端口，需要管理员权限；EADDRINUSE 表示端口被占用。 |

- 适用：阶段 2 技术选型 / 立项门；writing-plans

---

## 第 2 章 AI 使用说明书

- 章节页：<https://www.vibevibe.cn/Advanced/02-ai-tuning-guide/>
- 核心：让 AI 稳定产出可用代码：用精准提示词控制上下文与花费，用五步工作流约束过程，用版本记录兜底，用扩展配置与完整日志提升效率。
- 关键做法：
  - 五步工作流为探索项目结构 → 规划步骤 → 编码 → 测试验证 → 提交；复杂功能严格走完，微小改动可跳过探索。
  - 提示词指定文件路径与功能范围、删掉客套话、不设专家角色，并给 AI「不确定就说出来」的出口。
  - 三种权限模式用 Shift+Tab 循环切换：Default（读自动、写询问，推荐）、Plan（只读）、Accept Edits（写需确认，其余自动）。
  - 在 CLAUDE.md 写明「完成独立功能或修完 Bug 并验证通过后自动 git commit 并生成中文说明」；检查点 Esc+Esc 只回滚 AI 写入的文件。
  - 调试公式为完整错误日志 + 操作步骤 + 预期结果，多数问题 2-3 轮解决；先 git commit，再让 AI 跑 `pnpm install && pnpm build` 直到成功。
- 子页：

| 子页 | 要点 |
|---|---|
| 2.0 推荐配置 | 用 `/plugin` 装 code-review、context7 等；用 CC-SWITCH 管供应商、MCP 与 Skills，技能装到 `~/.claude/skills/`。 |
| 2.1 AI 编程的经济学 | 1 个汉字约 1 Token，1 行代码 5-15 Token，GLM 上下文 200K；节省成本的关键是缩小上下文范围。 |
| 2.2 VibeCoding 工作流 | Agent Native 三原则是意图驱动、异步协作、信任但验证；另含斜杠命令、Hooks、自定义 Agent、多 Agent 并行。 |
| 2.3 MCP、插件与 Skills | 插件可打包 Skills、命令、Agent、MCP；MCP 用 `claude mcp add --transport http|stdio` 接入，Windows 需 `cmd /c` 包装。 |
| 2.4 项目规则配置 | CLAUDE.md、.cursorrules、.iderules 作用相同，写清技术组合、编码规范、禁止行为；密钥放用户级配置。 |
| 2.5 高效调试心法 | 三个案例示范日志写法；错误模式速查覆盖类型、空值、导入、网络、端口、权限、语法错误。 |

- 适用：全阶段 / 全门禁；systematic-debugging、verification-before-completion、writing-plans

---

## 第 3 章 产品思维与文档驱动

- 章节页：<https://www.vibevibe.cn/Advanced/03-prd-doc-driven/>
- 核心：先验证想法值不值得做、再与 AI 确认理解，然后用五段式 PRD 作为唯一事实来源驱动写代码，减少 AI 猜测与范围膨胀。
- 关键做法：
  - 产品验证三步法：灵魂三问（用户是谁、痛点在哪、为何用你）、MVP 思维、快速验证；创业失败首要原因是「没有市场需求」，占 42%。
  - 访谈三条原则：谈对方生活不谈自己想法、问过去行为不问未来意愿、多听少说；目标是 5-10 次有效访谈，连续 3 次无新信息可停。
  - 让 AI 复述目标用户、3-5 个核心功能、不做的事情、可能的问题；8 项讨论结束清单全部打钩后再生成 PRD。
  - PRD 五部分对应初稿、中稿、定稿；没讨论过的字段标 TBD；Out-of-Scope 必须写，否则 AI 会自行加登录、云同步。
  - AI 严格按字面理解，不会自行补全；没写边缘情况就不做防抖，例如 0.5 秒内只响应一次。
- 子页：

| 子页 | 要点 |
|---|---|
| 3.0 PRD 模板 | 含文档信息与更新记录、用户故事、SMART 目标、需求列表 R001 起、交互四状态、边缘 Case、埋点表、灰度 1%→10%→50%。 |
| 3.1 想法验证实战 | 灵魂三问配需求三层次（痛点、痒点、爽点）；MVP 可用手工服务、介绍页收邮箱、原型或简化版本。 |
| 3.2 与 AI 确认需求 | 先说「列出多个你不确定的问题向我确认」，再按用户场景、核心功能、不做的事情、数据状态、技术约束逐项确认。 |
| 3.3 PRD 编写实战 | 用 Markdown 与 Mermaid；含语义化版本号、好与坏 PRD 对照；让 AI 基于讨论生成并检查有无猜测。 |
| 3.4 从 PRD 到代码 | AI 按接收、提取信息、建数据模型、设计接口、生成前端与后端执行；盲区是默认值、状态变化、优先级。 |

- 适用：阶段 1 立项 / 立项门；product-validation、brainstorming

---

## 第 4 章 你必须知道的开发基础

- 章节页：<https://www.vibevibe.cn/Advanced/04-dev-fundamentals/>
- 核心：补齐代码如何在互联网上运行的常识，让 PRD 能写成技术文档、让接口与配置能被看懂，并在出错时能定位。
- 关键做法：
  - 三种运行状态对应 `pnpm dev`、`pnpm build`、`pnpm start`；构建产物不能双击打开，Next.js 用 pnpm start，Vite 用 pnpm preview。
  - 技术选型在 PRD 之后按「明确需求 → 评估复杂度 → 最小可行方案」决策，默认 Next.js + PostgreSQL，纯前端才考虑 Vite。
  - 技术文档五部分不可省略：数据模型、接口设计、架构与流程图、第三方集成、技术决策；小项目可与 PRD 合并。
  - HTTP 请求含方法、URL、Headers、Body；401 未认证、403 无权限、404 不存在、429 限流、5xx 服务器错误。
  - 外部接口六步：获取凭证、选 SDK、配环境变量、最小测试、归档文档、业务开发；Key 只放 .env，敏感调用走后端代理。
- 子页：

| 子页 | 要点 |
|---|---|
| 4.0 代码运行的三种状态 | Dev 支持热重载，Build 生成 .next 或 dist，Production 模仿正式环境；改 .env 必须重启服务。 |
| 4.1 技术选型决策框架 | 回答项目要做什么、规模与并发、团队背景三问；0-50 人的公司推荐 Next.js + PostgreSQL + Vercel。 |
| 4.2 从 PRD 到技术文档 | 数据模型写表、字段、关系，接口写路径与方法；变更即更新文档，避免文档与代码脱节。 |
| 4.3 如何读懂 AI 生成的代码 | 变量、函数、条件、循环四个构件；用伪代码描述需求；数据量大时提示 AI 用高效查找。 |
| 4.4 API 与 HTTP 基础 | API 相当于远程函数；JSON 是接口标准格式；排查表覆盖 401、404、字段名与类型不符。 |
| 4.5 前后端分离概念 | 前端负责展示、后端负责处理，前后端一体的框架同项目但职责不变；按 auth、posts、users 划分模块。 |
| 4.6 配置文件格式 | JSON 用双引号、不支持注释；YAML 用空格缩进、支持 # 注释；传数据用 JSON，写配置用 YAML。 |
| 4.7 API 集成实战 | AI 应用用 `@ai-sdk/openai-compatible`，只改 baseURL、API Key 与模型名；限流、超时、认证失败各有处理办法。 |
| 4.8 项目说明书结构 | README 含简介、快速开始、环境变量、核心功能、项目结构、开发指南；可加给 AI 的项目上下文。 |
| 4.9 别再重复造轮子 | 先用 awesome-nodejs、npm trends、bundlephobia 找库，按 stars 超 1000、近 6 个月有更新判断；数据获取选 SWR（约 5KB）或 TanStack Query。 |

- 适用：阶段 2-3 / 立项门；api-design、database-design

---

## 第 5 章 界面与交互

- 章节页：<https://www.vibevibe.cn/Advanced/05-ui-ux/>
- 核心：先把界面交给 Tailwind 与可被 AI 修改的组件库生成，再按用户使用场景描述交互状态，让 AI 处理样式细节、人负责规则与场景。
- 关键做法：
  - 固定组件来源：Tailwind 内联 class 加 shadcn/ui（源码复制进项目），图标统一用 Lucide React，避免 AI 编造图标名。
  - 交互状态写进提示词：提交按钮加 Loading 并禁止重复点击，删除用 AlertDialog 二次确认，空列表给引导按钮。
  - 提示词说清响应式：移动端单列、桌面端双列；一次只做一块，分步迭代。
  - 一致性靠文档与 Skill：并行扫描 tailwind.config.js 与 src/components/ui/ 生成 DESIGN_SYSTEM.md；SKILL.md 只保留触发条件，详细规范放 references/，目录提交 Git。
  - 动画时长 200-500ms，一页最多 2-3 个高级效果。
- 子页：

| 子页 | 要点 |
|---|---|
| 5.1 AI 设计工具 | v0.app 直出 React + Tailwind + shadcn/ui 代码，Stitch 出设计稿，Figma 经 MCP 供 AI 读取颜色、间距、字体。 |
| 5.2 组件库 | 后台管理选 Ant Design 一类企业级库，面向用户的产品选 shadcn/ui，提示词里写明库名与组件名。 |
| 5.3 动画与交互库 | 纯 CSS 能做的就不引库，列表交错与页面切换用 Motion，滚动视差与时间线用 GSAP ScrollTrigger，3D 用 Three.js。 |
| 5.4 UI 风格与灵感 | Awwwards、Dribbble、Mobbin、Land-book 找参考，描述公式是风格关键词 + 配色方案 + 参考截图 + 具体场景。 |
| 5.5 让页面更高级的效果 | 视差滚动、3D 翻转、骨架屏、打字机等按类别提供现成提示词，可直接交给 AI 实现。 |
| 5.6 组件库 Skills 化 | 建 .claude/skills/my-project-design/，靠渐进式加载让规范按需进入上下文；B 端补表单分组、表格操作确认、Zod 校验。 |

- 适用：阶段 6 前端骨架 / 架构门；ui-design-system

---

## 第 6 章 数据持久化与数据库

- 章节页：<https://www.vibevibe.cn/Advanced/06-data-persistence-database/>
- 核心：按数据复杂度从 JSON 文件升级到 PostgreSQL，用 Drizzle 定义 Schema，让人能看懂并审查 AI 给出的表结构与查询。
- 关键做法：
  - 领 Neon 免费库（512 MB 存储、100 个项目），连接串写入 .env 的 DATABASE_URL，用 `pnpm drizzle-kit studio` 验证连通。
  - 表结构由 AI 1 号设计、AI 2 号以资深架构师身份审查，来回 1-2 轮，即 AI 交叉论证法。
  - PostgreSQL 不自动给外键建索引，必须手动加；索引常带来 100-1000 倍提升；小于 1000 行的表与布尔列不加。
  - 校验三层：数据库约束（notNull、unique、references）、后端校验、前端表单校验。
  - 生产要点：事务只放数据库操作，Supabase 直连 5432、连接池 6543，备份含自动备份、多地备份、恢复演练。
- 子页：

| 子页 | 要点 |
|---|---|
| 6.0 领取你的数据库 | Neon 区域选 Singapore，Supabase 的 Method 选 Transaction pooler，连接数超限时改用连接池地址。 |
| 6.1 数据存储演进 | CSV 二十条顺利、两百条吃力、两千条不堪用；JSON 解决多值与嵌套，但查询、约束、并发不足；有用户概念就该用数据库。 |
| 6.2 数据库基础概念 | 表行列、主键外键；一对多最常见、多对多靠中间表、一对一用 UNIQUE 保证；金额用 numeric，时间用 timestamptz，文本用 text。 |
| 6.3 如何操作数据库 | CRUD 审查重点看 UPDATE/DELETE 有无 WHERE；N+1、深页变慢、先查再改分别改用 JOIN、游标分页、UPSERT；Schema 变更走迁移。 |
| 6.4 数据库设计与优化 | 外键列逐个加索引，jsonb 用 GIN，deleted_at 标记用部分索引；RLS 需 ENABLE ROW LEVEL SECURITY 且四类策略齐全；诊断顺序为是否数据库、缺索引（EXPLAIN ANALYZE）、N+1、连接与事务。 |

- 适用：阶段 7 数据库设计 / 架构门；database-design

---

## 第 7 章 后端 API 开发

- 章节页：<https://www.vibevibe.cn/Advanced/07-backend-api/>
- 核心：用 Next.js 的 route.ts 把前端与数据库连起来，先跑通 CRUD，再按真实需求补分页、校验、错误处理与文档。
- 关键做法：
  - app/api/ 下一文件一接口，导出 GET、POST 等函数；统一响应为 `{ success: true, data }` 与 `{ success: false, error: { message } }`。
  - 页面展示直接在 Server Component 查库，简单写操作用 Server Action，只有外部消费者与动态交互才需要 API Route。
  - 列表只留一个入口靠参数组合：page/limit、sort/order、tag、q；小于 1 万条用 OFFSET 分页，超过 10 万条改游标分页。
  - 所有 POST 与 PATCH 加 Zod 校验；400、404、409、500 分流，500 只回通用提示，细节进日志。
  - 防重复靠联合唯一约束或幂等键；批量操作走一个事务并限制 100 条；上传限制 5MB、只收 jpg/png/webp。
- 子页：

| 子页 | 要点 |
|---|---|
| 7.0 跑通你的第一个前后端应用 | Todo 应用分 schema.ts、route.ts、page.tsx 三件；`pnpm drizzle-kit push` 同步表，pnpm dev 启动，Drizzle Studio 核对写入。 |
| 7.1 一个接口不够用了 | 详情页用嵌套结构一次返回，列表页用扁平结构；分页响应带 meta 的 total、page、limit、totalPages。 |
| 7.2 当接口出了问题 | 前端校验挡不住 Postman 与脚本，后端必须再校验；写操作后调 revalidatePath，发通知等用 after() 推到响应之后。 |
| 7.3 让接口更好用 | OpenAPI 3.0 加 Swagger UI 由代码生成文档；改字段名先双写再废弃，大改用 /api/v1 与 /api/v2；实时更新按频率选轮询、SSE、WebSocket。 |

- 适用：阶段 8 后端骨架 / 架构门；api-design、backend-architecture-acceptance

---

## 第 8 章 安全与用户认证

- 章节页：<https://www.vibevibe.cn/Advanced/08-auth-security/>
- 核心：从密钥泄露的代价出发，把密钥管理、身份认证、路由权限三层守住，再补上注入类攻击与 AI 接口的特有风险。
- 关键做法：
  - 密钥全放 .env，代码用 process.env 读取，.gitignore 屏蔽 .env；只有 NEXT_PUBLIC_ 前缀的值会进浏览器；已追踪的文件先 `git rm --cached`。
  - 认证用 Better Auth，Session 存自己的库，含 user、session、account、verification 四张表；密码用 bcrypt 哈希，不自写认证。
  - 受保护页面在服务端取 session，未登录 redirect('/login')；middleware.ts 用 matcher 匹配 /dashboard/:path*、/admin/:path*，API 内部再查角色。
  - 角色从 admin 与 user 两档起步；Cookie 设 HttpOnly、Secure、SameSite；生产必须 HTTPS。
  - 每月跑 `pnpm audit`，密钥每 3-6 个月轮换，AI 接口限制每用户每分钟 10 次并返回 429。
- 子页：

| 子页 | 要点 |
|---|---|
| 8.0 用户系统快速示例 | 用 Better Auth 跑注册、登录、登出与受保护页面，`pnpm drizzle-kit push` 建表，Drizzle Studio 可见 password 存的是哈希值。 |
| 8.1 密钥管理与环境变量 | .env、.gitignore、.env.example 三件配齐；改了 .env 必须重启；部署时在平台环境变量页逐条填。 |
| 8.2 认证方式与方案选择 | Session 可随时撤销、Token 免查库但难撤销；OAuth、Passkeys、Magic Link 按场景选，新项目用 Better Auth。 |
| 8.3 路由保护与权限控制 | 前端隐藏入口无效；Middleware 拦页面、API 内部再校验构成纵深防御；同项目内通常不配 CORS，跨域只放行具体域名。 |
| 8.4 安全检查与问题排查 | 开发前配 .gitignore，开发中查边界与校验，上线前跑 `pnpm audit` 并确认 HTTPS，上线后定期轮换；手册按症状给原因与解法。 |
| 8.5 进阶安全防护 | ORM 参数化防 SQL 注入，React 自动转义且不用 dangerouslySetInnerHTML，Server Action 与 SameSite 防 CSRF；AI 侧防提示注入、过滤输出、限制速率。 |

- 适用：阶段 12 安全审查 / 业务门；backend-security-review、frontend-security-review

---

## 第 9 章 功能测试与自动化

- 章节页：<https://www.vibevibe.cn/Advanced/09-testing-automation/>
- 核心：用回归测试防止修改一处却影响另一处，按测试金字塔分配单元、API、E2E 三层检查，再用 Git Hooks 与持续集成让检查自动执行。
- 关键做法：
  - 一个接口至少覆盖四类场景：正常、校验、权限、边界；例如合法评分返回 200，评分 6 返回 400，未登录返回 401，电影不存在返回 404。
  - 需要自动化测试的判断依据：页面超过 5 个、核心流程超过 3 步、遇到过修改一处却影响另一处、开始害怕改代码；命中 0/4 项时手动测试足够。
  - 两层检查机制：Husky 的 pre-commit 运行 `pnpm test`，失败则阻止提交；GitHub Actions 在 ubuntu-latest、node-version 20 的干净环境重新运行。
  - Flaky 测试三大成因是异步时序、数据竞争、环境差异；用 `--repeat-each=5` 连跑 5 次区分偶发失败与真实缺陷，禁用 waitForTimeout。
  - 失败定位按错误信息、截图（test-results/）、trace 中的网络请求三步走；TDD 分为 Red、Green、Refactor 三段，适合核心逻辑与稳定接口。
- 子页：

| 子页 | 要点 |
|---|---|
| 9.1 为什么需要测试 | 解释回归测试与测试金字塔的分层成本，判定标准是「不依赖 UI 就能验证就不要用 E2E」，并给出人把控策略与 AI 生成测试的分工。 |
| 9.2 API 测试与 E2E 测试 | 给出评分接口四类场景的 Vitest 示例、Playwright 录制命令 `npx playwright codegen`、Flaky 排查清单与日志三步定位。 |
| 9.3 自动化工作流 | 给出 Husky 安装与初始化命令、test.yml 的 on/jobs/steps 结构、TDD 短评接口示例，并主张每次修完 bug 补一个测试。 |

- 适用：阶段 11 业务模块开发 / 业务门；testing-strategy、systematic-debugging、verification-before-completion

---

## 第 10 章 Localhost 与公网访问

- 章节页：<https://www.vibevibe.cn/Advanced/10-localhost-public-access/>
- 核心：说明 localhost 只指向请求设备自身，按本机、局域网、广域网三层梳理可达范围，给出局域网真机调试与临时内网穿透两种分享方法。
- 关键做法：
  - localhost 对应 127.0.0.1 回环地址，数据包不经过网线或 WiFi；别人的 localhost 指向他自己的设备。
  - 局域网 IP 通常是 192.168.x.x 或 10.x.x.x，用 ipconfig（Windows）、ifconfig（macOS）、ip addr（Linux）查找，手机与电脑须连同一 WiFi。
  - 手机打不开按顺序排查：同一 WiFi、服务监听 0.0.0.0（`pnpm dev --hostname 0.0.0.0` 或 `pnpm dev --host`）、防火墙放行 3000 端口，原文称九成情况是防火墙拦截。
  - 家庭宽带一般无独立公网 IP，运营商 NAT 让外部请求无法直达；要让所有人访问必须部署到有公网 IP 的服务器。
  - 隧道命令 `cloudflared tunnel --url http://localhost:3000`，免费模式每次生成新的 trycloudflare.com 链接，Ctrl+C 后立即失效。
- 子页：

| 子页 | 要点 |
|---|---|
| 10.1 从 Localhost 到互联网 | 讲回环地址语义、局域网 IP 查找、框架监听参数、Windows 防火墙放行命令，以及局域网直连、内网穿透、正式部署三条路径的场景与局限。 |
| 10.2 内网穿透 | 讲隧道由本机主动出站连接公网服务器，Cloudflare Tunnel 安装与启动、ngrok 与 cpolar 对比，以及只暴露 Web 端口、用完立即关闭的安全清单。 |

- 适用：阶段 9 后端联调 与 阶段 14 上线 / 架构门 与 上线门；browser-verification、deployment-and-ops

---

## 第 11 章 Git 版本控制与协作开发

- 章节页：<https://www.vibevibe.cn/Advanced/11-git-collaboration/>
- 核心：用 commit 记录项目快照取代文件名备份，把代码推到远程仓库实现备份与协作，用分支与 PR 隔离实验并在合并前审查。
- 关键做法：
  - Git 保存整个项目在某一时间点的完整快照，提交采用增量存储，几百次提交的 .git 通常只有几十 MB；完成一件能用一句话说清的事就提交一次。
  - 三区模型：工作区经 `git add` 进入暂存区，再经 `git commit` 进入本地仓库；`git status` 判断修改处在哪个环节。
  - 回滚分三种：尚未提交时丢弃工作区修改；已提交未推送用 git reset 改写历史；已推送用 git revert 新增反向提交，保留历史。
  - 提交信息用类型前缀：feat、fix、docs、refactor、style、chore。
  - 安全底线：.env、含密码的 DATABASE_URL、.pem/.key/id_rsa、各类 Token 不得进入仓库；已推送的敏感文件要清除历史（git filter-branch 或 BFG）并立即轮换密钥。
- 子页：

| 子页 | 要点 |
|---|---|
| 11.1 为什么需要 Git | 从「最终版」文件命名困境讲起，介绍 commit 与存档的对应、三区模型、三种回滚场景、看 diff 的意义与提交信息前缀约定。 |
| 11.2 推上云端，开始协作 | 讲单点故障、平台选择、SSH 密钥认证及 HTTPS 加 Token 备选、首次 push 命令、push 与 pull 节奏、clone 与压缩包的区别、冲突标记读法、.gitignore 与已被跟踪文件的移除、仓库页面导览。 |
| 11.3 分支、PR 与团队工作流 | 讲 feature/xxx 命名、GitHub Flow 十步流程、PR 与 Code Review 的作用、PR 页面标签、个人项目是否开分支，以及 .gitattributes 写入 `* text=auto eol=lf` 统一换行符、文件名大小写与路径分隔符。 |

- 适用：全阶段 / 全门禁；requesting-code-review

---

## 第 12 章 无服务器部署与 CI/CD 自动化

- 章节页：<https://www.vibevibe.cn/Advanced/12-serverless-deploy-cicd/>
- 核心：把 GitHub 代码接到托管平台，用 Serverless 环境获得固定公网地址，由平台自动构建与发布，并建立监控与成本意识。
- 关键做法：
  - 首次部署三个前提：代码已推到 GitHub、本地 `pnpm build` 构建完成、记录 .env 中全部变量，因为 .env 被 .gitignore 排除，平台拉不到。
  - 平台通用五步：连接 GitHub、选择仓库、配置构建、添加环境变量、部署；换平台无需改代码。
  - 环境变量修改后需重新部署才生效，且只影响新部署；国内平台支持把 .env 内容整段粘贴批量导入。
  - 加速区域含中国大陆需 ICP 备案，审核一到两周；选「全球（不含中国大陆）」无需备案，预览链接默认 3 小时有效。
  - 成本分阶段：低于 1000 用户用免费套餐月成本 0；1000 到 10000 用户约 50 到 200 元；超过 10000 用户再评估自建；需设置账单告警。
- 子页：

| 子页 | 要点 |
|---|---|
| 12.1 部署到 EdgeOne Pages | 给出导入仓库、框架预设 Next.js、根目录 ./、构建命令 pnpm build、输出目录 .next、Node 22 默认值、最小权限授权与按日志关键词排错的完整步骤。 |
| 12.2 部署到类 Vercel 平台 | 对比 Vercel、Cloudflare Pages、阿里云 ESA、Netlify、Railway 的访问速度、免费额度与适用场景，并给出 `vercel login`、`vercel deploy`、`vercel --prod` 命令。 |
| 12.3 CI/CD 与自动化 | 讲 webhook 触发推送即部署、main 与 feature 分支对应生产与预览环境、GitHub Actions 的 on、runs-on、steps 结构，以及 .cnb.yml 的 master 与 develop 双环境配置。 |
| 12.4 运维基础与成本优化 | 讲各平台分析面板与日志入口、数据库超时这类日志的排查套路、平台默认提供的扩缩容、CDN、图片优化与 HTTPS，以及免费额度的具体限制。 |

- 适用：阶段 14 上线 / 上线门；deployment-and-ops

---

## 第 13 章 域名、DNS 与网络接入

- 章节页：<https://www.vibevibe.cn/Advanced/13-domain-dns/>
- 核心：解决平台默认链接别人打不开的问题，用自己的域名替换平台临时地址，把域名购买、DNS 解析与 HTTPS 证书一次配齐。
- 关键做法：
  - 默认链接报 403 来自国内平台对默认域名的鉴权限制，只有项目所有者凭带 token 的链接可访问，绑定自定义域名后对所有人开放。
  - 云服务器用 A 记录指向 IP，平台部署用 CNAME 指向平台域名；CNAME 跟随平台，换 IP 不必改记录。
  - 绑定流程是平台侧给出 CNAME 值，注册商侧填写类型 CNAME、名称 @ 或 www、TTL 自动，再用 `nslookup`、`ping` 验证，传播从几分钟到几小时。
  - 选择域名优先 .com，注意续费价格（存在首年 ¥1、续费 ¥300 的情况）；购买得到二级域名加顶级域名，子域名可无限创建。
  - 证书由 Let's Encrypt 免费签发、有效期 90 天并自动续期；EdgeOne Pages 需在域名管理页手动点击申请免费证书，未见小锁等待 5-15 分钟。
- 子页：

| 子页 | 要点 |
|---|---|
| 13.1 域名购买与 DNS 配置 | 介绍 Cloudflare、阿里云、腾讯云、Namecheap 四个注册商与选名原则，按 A/CNAME 两类记录完成解析，再配 HTTPS。 |
| 13.2 备案与访问问题排查 | 按服务器位置判断是否需 ICP 备案，给出三个免备案方案（部署平台选「全球（不含中国大陆）」、Vercel / Cloudflare Pages、香港节点 VPS），并用状态码首位数字定位 403、超时、访问慢与证书错误。 |

- 适用：阶段 14 上线 / 上线门；deployment-and-ops

---

## 第 14 章 云服务器运维与项目部署

- 章节页：<https://www.vibevibe.cn/Advanced/14-vps-ops-deploy/>
- 核心：在自己购买的 VPS 上用 1Panel 与 Docker 完成部署，把加固、容器、端口映射、证书与日志这套运维动作走通。
- 关键做法：
  - 内存比 CPU 重要，2 核 2GB 是 Node.js 加数据库加面板的下限，1GB 机器装完面板几乎没余量；1GB 内存建 2GB Swap，2GB 内存建 1-2GB Swap。
  - 安全组默认只放行 22，必须另开 80、443 与面板端口；改 SSH 端口（如 24324）后先用新端口连通再关 22；Fail2Ban 默认连续失败 5 次封 IP。
  - 容器找邻居用容器名（形如 1Panel-postgresql-xxxx），主机找容器用 127.0.0.1，外部用公网 IP 加映射端口；填 localhost 会报 ECONNREFUSED。
  - 运行环境主机端口 3001 映射容器 3000，启动命令 `git pull && pnpm build && pnpm start`，更新只需在面板重启；新增依赖才需手动执行 `pnpm install`。
  - 证书在「网站 > 证书」用 HTTP 验证申请并开启自动续签，HTTP 选项选自动跳转到 HTTPS；排查问题先看日志里的 Error、Warning 与自定义日志。
- 子页：

| 子页 | 要点 |
|---|---|
| 14.1 VPS 选购指南 | 对比国内与海外厂商，按 CPU、内存、硬盘、带宽四个参数选型，机房优先香港（免备案，延迟 30-60ms），系统选 Ubuntu 22.04/24.04 LTS。 |
| 14.2 VPS 初始化与安全配置 | 给出三种登录方式、`apt update && apt upgrade -y`、装 1Panel、开安全组、改 SSH 端口、Fail2Ban、密钥登录、WAF 的清单式加固顺序。 |
| 14.3.1 应用商店与 Docker 基础 | 用应用商店装 PostgreSQL、OpenResty、Node.js，理解镜像、容器、仓库三层概念，Volume 数据存放在 /opt/1panel/，掌握容器寻址三条规则。 |
| 14.3.2 部署 Next.js 应用 | 上传代码、创建运行环境、配端口映射与环境变量，用 http://服务器IP:3001 访问，更新靠面板重启；内存不足时按前文创建 Swap。 |
| 14.3.3 部署静态网站 | 本地构建出 dist 或 out 产物，创建静态网站后上传解压；SPA 需加 try_files 伪静态规则，静态资源可设 30 天缓存。 |
| 14.3.4 部署前后端分离应用 | 前端由 OpenResty 托管，后端单独起运行环境，把 /api/ 反向代理到容器端口，最后将前端 API 地址改成相对路径 /api。 |
| 14.4 配置域名与证书 | 添加 @ 与 www 两条 A 记录（TTL 600），在 1Panel 创建网站认领域名，申请 Let's Encrypt 证书、启用强制 HTTPS 跳转与自动续签。 |
| 14.5 其他好玩的应用 | 用应用商店装 Umami、n8n、Bitwarden（Vaultwarden 的 Rust 轻量实现）、AList（默认端口 5244），扩展服务器用途。 |

- 适用：阶段 14 上线 / 上线门；deployment-and-ops

---

## 第 15 章 SEO、分享与数据统计

- 章节页：<https://www.vibevibe.cn/Advanced/15-seo-analytics/>
- 核心：让上线后的站点能被搜到、能被分享成卡片、能被数据观测，并补齐隐私政策、用户协议与备案。
- 关键做法：
  - Metadata、Sitemap、Robots.txt 是能否被找到的入场券，排名取决于内容本身，效果通常 3-6 个月才显现。
  - 标题 50-60 字符并把关键词前置；Core Web Vitals 良好值为 LCP < 2.5s、INP < 200ms、CLS < 0.1。
  - OG 五个核心标签为 og:title、og:description、og:image、og:url、og:type，万能图片尺寸 1200 x 630，重要内容避开边缘 20%，带图链接点击率通常是纯文本的 2-5 倍；微信不读 OG，需微信 JS-SDK。
  - Umami 脚本小于 1KB 且不用 Cookie，默认账号 admin / umami 登录后必须改密码；跳出率高不必然是问题，PV 高 UV 低说明有忠实用户。
  - 合规三件：隐私政策放 /privacy、用户协议放 /terms 并在注册时勾选同意、大陆服务器需 ICP 备案并在底部展示备案号，备案约 2-4 周；GDPR 罚款上限为全球年营业额 4% 或 2000 万欧元。
- 子页：

| 子页 | 要点 |
|---|---|
| 15.1 Open Graph 与社交分享 | 解释 OG 协议与平台差异（X 回退读 OG、LinkedIn 严格按 OG 渲染、微信独立生态），给图片尺寸与设计原则，强调用调试工具验证。 |
| 15.2 SEO 全攻略 | 讲爬取、索引、排名三步，对应 Next.js 的 metadata 对象、app/sitemap.ts、app/robots.ts，再讲标题、E-E-A-T、URL 结构、页面速度、结构化数据与百度对未备案站的收录门槛。 |
| 15.3 Umami 数据统计 | 加追踪代码后解读 PV、UV、Visits、跳出率、流量来源与设备，区分主要指标与虚荣指标，用 data-umami-event 属性或 window.umami.track 做事件追踪。 |
| 15.4 法律合规 | 列出隐私政策七项内容、用户协议七项内容、备案类型与时间，并给合规清单核对页面位置与用户同意环节。 |

- 适用：阶段 13 上线前总验收 / 上线门；seo-analytics-compliance

---

## 第 16 章 用户反馈与产品迭代

- 章节页：<https://www.vibevibe.cn/Advanced/16-user-feedback-iteration/>
- 核心：上线后靠反馈渠道、优先级排序、用户访谈与更新节奏，把「数据说发生了什么」和「用户说为什么」合并成下一步动作。
- 关键做法：
  - 反馈分 Bug、功能请求、体验、性能、内容及其他几类；按影响范围与严重程度分 P0 立即处理、P1 二十四小时内、P2 本周内、P3 有时间再说。
  - RICE = (Reach × Impact × Confidence) / Effort，Impact 取 1-3 分；示例中修复登录 Bug 得 3000 分，是深色模式 167 分的 18 倍；简化版 ICE 三项各 1-10 分取平均。
  - 访谈遵循 The Mom Test 三原则：谈过去不谈未来、谈具体不谈笼统、少问多听；5-10 次即够，从第 4 次开始反复听到同样问题。
  - 从反馈到行动分四层提炼：原始反馈、问题、洞察、行动，跳过中间两层容易改错地方。
  - 更新分三层：热修复随时发、小更新每周一次、大版本每月一次；精力按 80% 稳定性与 Bug、20% 新功能分配。
- 子页：

| 子页 | 要点 |
|---|---|
| 16.1 面对真实用户 | 讲上线初期的期望差距与心理调适，早期只看留存、核心功能使用率、问题反馈、使用场景，并把应用内反馈按钮、邮箱、社交媒体、用户群组（早期 20-30 人）建起来。 |
| 16.2 反馈分类与优先级 | 先按类型归纳，再用影响范围与严重程度矩阵分 P0-P3，必要时用 RICE 或 ICE 排序，建立收集、分类、排序、处理、验证、通知的闭环，并学会礼貌拒绝。 |
| 16.3 理解用户 | 用 The Mom Test 设计问题并观察操作，访谈后分四层提炼，把定性反馈与 Umami 定量数据对照，确定北极星指标，并用单一变量的 A/B 测试验证假设。 |
| 16.4 迭代与成长 | 用三层更新、灰度发布（白名单、10% 百分比、条件触发、功能开关）与更新日志管理变更，早期按一周一版推进，识别并加强受欢迎功能、移除无人使用的功能。 |

- 适用：阶段 14 上线后观察 / 上线门；user-feedback-iteration

---

## 延伸页

| 页面 | 内容 |
|---|---|
| [Next Level](https://www.vibevibe.cn/Advanced/99-next-level/) | 全流程收尾：从依赖到掌控、从黑盒到透明、从恐慌到从容，强调持续学习与把想法变成可访问产品。 |
| [5 分钟搭建云端开发环境](https://www.vibevibe.cn/Advanced/web-ide.html) | 在 cnb.cool 上用预装镜像开云端开发环境：Claude Code、Codex、Gemini Code Assist、Node 24、Python 3.11、Docker、Git、VS Code 53 个扩展、1Panel（端口 34246）。 |
| [随时随地 AI 编程](https://www.vibevibe.cn/Advanced/happy-coder.html) | Happy Coder 远程控制工具：用手机遥控电脑端的 Claude Code、Codex、Gemini，端到端加密，支持多会话与语音输入。 |
