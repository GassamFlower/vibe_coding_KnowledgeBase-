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
- `01-idea.md` … `05-constitution.md` — phase，立项门路径（stage 1-5）
- `06-impl-breakdown.md` … `10-backend-integration.md` — phase，架构门前半（stage 6-10）
- `11-arch-acceptance.md` — phase，架构验收（stage 11，临近架构门）
- `12-business-dev.md` — phase，业务模块开发（stage 12，业务门）
- `13-security-review.md` — phase，后端安全审查（stage 13，临近业务门）
- `14-prerelease-review.md` — phase，上线前总验收（stage 14，上线门）
- `15-release.md` — phase，上线 + Skill 反哺（stage 15，临近上线门）
- `gate-立项门.md` — gate，立项门 7 项硬指标（stage 5 注入）
- `gate-架构门.md` — gate，架构门 7 项硬指标（stage 11 注入）
- `gate-业务门.md` — gate，业务门 4 项硬指标（stage 13 注入）
- `gate-上线门.md` — gate，上线门 6 项硬指标（stage 15 注入）

> 阶段编号说明：stage 1-5 = 立项门；6-11 = 架构门；12-13 = 业务门；14-15 = 上线门。
> 与 workflows/全流程速查.md 的 14 阶段模型偏移 1（gen N = 14模型 N-1，N>=6），为兼容既有 01-05 模块而沿用生成器自身编号。
