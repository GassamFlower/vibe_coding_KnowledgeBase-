# 信息模块（modules/）

每个文件是一个**可插拔的信息模块**，Go 生成器按当前阶段把它拼进 `AGENTS.md`。

## 模块契约

文件顶部用 HTML 注释写 frontmatter：

```markdown
<!--
id: 01-idea          # 模块唯一 id（建议 stage-名称）
stage: 1             # 归属阶段（生成器据此纳入）
type: phase|gate|cross   # phase=开发阶段 / gate=门禁 / cross=跨阶段（如状态判断）
title: 阶段1 纯想法 → 立项
next: 立项（阶段2）   # 该阶段完成后下一步（仅 phase 需要）
-->
# 正文（给 IDE 的 LLM 看的方法论，要可操作）
```

## 规则

- `type: cross` 的模块（如状态判断）每个阶段都会注入。
- `type: gate` 的模块在该阶段临近门禁时注入。
- `type: phase` 的模块只在当前阶段注入。
- 加模块 = 加文件；删模块 = 删文件。生成器不写死模块列表。

## 当前模块

- `00-state-judgment.md` — cross，5 种项目状态判断
- `00-cost-context.md` — cross，成本与心力管理（每次会话注入）
- `01-idea.md` … `05-constitution.md` — phase，立项门路径
- `gate-立项门.md` — gate，立项门 7 项硬指标
