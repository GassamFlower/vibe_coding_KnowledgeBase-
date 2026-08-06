# Vibe_Coding companion — 开发工具里的 process companion

把 Vibe_Coding 方法论固化成一个「能在开发工具里用的 agent 功能」：开发者一边动手，
Go 后端生成的 `AGENTS.md` + `PROJECT_STATUS.md` 让 IDE 里的 LLM 随时知道
「做到哪了、下一步干啥、门禁还差什么」。

## 设计

```
Go 后端（大脑）                 IDE（开发者的手）            开发者
─────────────                 ───────────────            ───────
状态机引擎                     读 AGENTS.md               动手写代码
信息模块库(modules/)  ─拼装→  PROJECT_STATUS.md  ←──     边干边看提示
门禁/校验逻辑        ─生成→   "下一步该做 X"             走到哪、下一步啥
```

- **Go 后端 = 大脑**：跑状态机、管信息模块、算门禁。不直接干活，产出两份文件给 IDE 消费。
- **IDE 里的 LLM = companion**：读这两份文件，一边开发者动手、一边提示。
- **信息模块（modules/）**：每个阶段 / 门禁一个独立 `.md`，带 frontmatter。Go 按当前阶段拼装进 AGENTS.md。
  你要加「信息模块」= 在 modules/ 加一个 `.md` 文件，无需改生成器逻辑。
- **MCP = 可选升级**：companion 跑顺后，Go 后端可再暴露 MCP Server，让 agent 主动调工具（自动跑门禁、改文件）。

## 目录

```
companion/
├── modules/            信息模块（独立 .md，Go 拼装）
│   ├── 00-state-judgment.md   cross：5 种项目状态判断
│   ├── 00-cost-context.md     cross：成本与心力管理（每次会话注入）
│   ├── 01-idea.md … 15-release.md   phase：14 阶段（stage 1-15）
│   └── gate-立项门.md … gate-上线门.md   gate：4 道门禁硬指标
├── templates/          渲染模板
│   ├── AGENTS.md.tmpl
│   └── PROJECT_STATUS.md.tmpl
├── generator/          生成器
│   ├── main.go         Go 版（接你 DC Ops 后端）
│   ├── main.py         Python 原型（本环境验证用）
│   └── go.mod
└── README.md
```

## 用法

```bash
# Python 原型（即时验证）
python generator/main.py --project /path/to/your/project --stage 1

# Go 版（接入 DC Ops 后端后）
go run ./generator --project /path/to/your/project --stage 1
```

生成后，把目标项目在 Cursor / VSCode+Cline / Claude Code 中打开，IDE 的 LLM 会自动读
`AGENTS.md` 与 `PROJECT_STATUS.md`，在右侧 / 对话里提示你下一步。

## 加一个信息模块

在 `modules/` 新建 `16-xxx.md`（示例，stage 编号延续 1-15 之后或插入对应阶段）：

```markdown
<!--
id: 16-xxx
stage: 16
type: phase
title: 阶段16 新阶段
next: 下一个阶段
-->
# 阶段16：新阶段
...
```

生成器按 `stage` 自动纳入对应阶段的 AGENTS.md，零改码。

## 状态 / 阶段 / 门禁（stage 1-15 全流程）

| stage | 阶段 | 类型 | 临近门禁 |
|-------|------|------|----------|
| 1 | 纯想法 → 立项 | phase | — |
| 2 | 立项 | phase | — |
| 3 | 选型 | phase | — |
| 4 | 架构设计 | phase | — |
| 5 | 宪法 | phase | 立项门 |
| 6 | 实施真源文档拆解 | phase | — |
| 7 | 前端骨架 | phase | — |
| 8 | 数据库设计 | phase | — |
| 9 | 后端骨架 | phase | — |
| 10 | 后端联调 | phase | — |
| 11 | 架构验收 | phase | 架构门 |
| 12 | 业务模块开发 | phase | — |
| 13 | 后端安全审查 | phase | 业务门 |
| 14 | 上线前总验收 | phase | — |
| 15 | 上线 + Skill 反哺 | phase | 上线门 |

> 阶段编号说明：生成器用自身编号（stage 1-15），与知识库 14 阶段模型偏移 +1（gen N = 14模型 N-1，N≥6），为兼容既有 01-05 模块沿用。cross 模块每个 stage 都注入；gate 模块在抵达对应 stage（5/11/13/15）时注入硬指标正文。
