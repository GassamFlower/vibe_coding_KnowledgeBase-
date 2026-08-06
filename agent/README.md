# Vibe Coding Engineer — 工程化智能体（可执行层）

> 基于 Vibe_Coding 知识库（已验证方法论）固化的工程化智能体。
> 把"人工翻文档 + 复制 prompt + 手动验收"升级为"智能体自动执行流程 + 自动收集证据 + 自动门禁判断"。

---

## 这是什么

Vibe_Coding 知识库的**可执行层**。知识库（docs/skills/templates/workflows）是静态文档，给人翻；agent/ 是可执行能力，给 AI 执行。

两者关系：知识库是单一来源，agent 引用知识库，不复制内容。知识库更新时，agent 自动继承。

---

## 目录结构

```
agent/
├── README.md                    ← 本文件（目录说明）
├── SKILL.md                     ← 智能体主入口（身份/触发/能力/调用方式）
├── state-machine.md             ← 项目状态机（5 状态 + 14 阶段 + 4 门禁）
├── gate-checks/                 ← 门禁自动验收（4 道）
│   ├── 立项门.md                ← ✅ 已实现
│   ├── 架构门.md                ← ✅ 已实现
│   ├── 业务门.md                ← ✅ 已实现（v0.3.0）
│   └── 上线门.md                ← ✅ 已实现（v0.4.0）
├── workflows/                   ← 可执行流程（按门禁组织）
│   ├── 立项门闭环.md            ← ✅ 已实现
│   ├── 架构门闭环.md            ← ✅ 已实现
│   ├── 业务门闭环.md            ← ✅ 已实现（v0.3.0）
│   └── 上线门闭环.md            ← ✅ 已实现（v0.4.0）
├── templates/                   ← 智能体专用模板
│   └── PROJECT_STATUS.agent.md  ← ✅ 智能体自动维护的状态文件
├── skills/                      ← 子能力引用（指向 ../../skills/）
│   └── README.md                ← 说明：11 个方法论在 ../../skills/
└── references/                  ← 精简版深度手册（指向 ../../docs/）
    └── README.md                ← 说明：4 篇手册在 ../../docs/
```

---

## 当前版本（v0.4.0 MVP）

已实现**全量 4 门禁完整闭环**：

- 状态判断（0 纯想法 / 1 烂尾 / 2 半成品 / 3 黑盒 / 4 健康）
- 立项门：立项 → 选型 → 架构设计 → 宪法 → 立项门自动验收
- 架构门：实施真源文档拆解 → 前端骨架 → 数据库设计 → 后端骨架 → 后端联调 → 架构验收 → 架构门自动验收
- 业务门：业务模块开发（多轮）→ 安全审查 → 业务门自动验收
- 上线门：上线前总验收 → 上线 → Skill 反哺 → 上线门自动验收
- 状态推进：立项门 → 架构门 → 业务门 → 上线门 → 项目收官

---

## 怎么用

### 方式 1：在 WorkBuddy 中使用

把本目录（agent/）复制或 symlink 到 `~/.workbuddy/skills/vibe-coding-engineer/`，WorkBuddy 会自动识别 SKILL.md，按方法论执行。

### 方式 2：在项目内使用

把 agent/ 复制到项目根目录，在项目的 AGENTS.md / CLAUDE.md 中引用 SKILL.md。

### 方式 3：移植到 DC Ops 平台

agent/ 的结构（状态机 + 门禁检查 + 工作流 + 模板）可以转化为 Go 后端的规则引擎，接入现有 Copilot。

---

## 与知识库的关系

```
Vibe_Coding/                    ← 知识库（单一来源）
├── docs/                       ← 深度手册（4 篇）
├── skills/                     ← 11 个方法论（子能力来源）
├── templates/                  ← 工作台模板（8 个）
├── workflows/                  ← 流程速查（1 篇）
├── sources/                    ← 原始笔记归档
└── agent/                      ← 智能体（可执行层，本目录）
```

**演进原则**：
- 知识库是单一来源，agent 引用不复制
- 知识库更新 → agent 自动继承
- agent 使用中发现新规则 → 反哺知识库
- 反哺时更新知识库的 skills/ 和 docs/，agent 自动继承

---

## Harness 6 支柱映射

| Harness 支柱 | 知识库来源 | agent 实现 |
|---|---|---|
| 上下文管理 | docs/操作手册 + Prompt 模板库 | references/ + 动态 prompt 生成 |
| 工具系统 | skills/ 11 个方法论 | skills/（引用 ../../skills/） |
| 执行编排 | workflows/ 流程速查 | state-machine.md + workflows/ |
| 评估观测 | templates/ 门禁卡 | gate-checks/ |
| 状态记忆 | templates/ PROJECT_STATUS | templates/PROJECT_STATUS.agent.md |
| 约束恢复 | docs/ 操作手册第 15 章 | SKILL.md 异常处理表 |
