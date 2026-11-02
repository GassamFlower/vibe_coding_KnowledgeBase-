# 子能力引用说明

本目录不复制方法论内容，而是引用知识库的 `../../skills/` 目录。

## 引用的 18 个方法论

| Skill 文件 | 何时调用 |
|---|---|
| `../../skills/brainstorming.md` | 立项阶段，挖掘需求 |
| `../../skills/product-validation.md` | 立项阶段，验证需求真实性（用户证据 / Out-of-Scope / 风险） |
| `../../skills/writing-plans.md` | 立项后，写开发计划 |
| `../../skills/project-breakdown.md` | 立项后，梳理功能清单 |
| `../../skills/ui-design-system.md` | 前端骨架阶段，确定设计规范与 token |
| `../../skills/database-design.md` | 数据库阶段 |
| `../../skills/api-design.md` | 后端骨架阶段，写 API 接口清单 |
| `../../skills/backend-architecture-acceptance.md` | 后端骨架验收 |
| `../../skills/backend-security-review.md` | 安全审查 |
| `../../skills/frontend-security-review.md` | 前端安全审查（纯前端 / 小程序 / 前后端分离的前端） |
| `../../skills/requesting-code-review.md` | 代码审查 |
| `../../skills/systematic-debugging.md` | 调试 Bug |
| `../../skills/browser-verification.md` | 前端浏览器验证 |
| `../../skills/testing-strategy.md` | 业务模块完成后、宣布完成前——先测试再验收 |
| `../../skills/verification-before-completion.md` | 任何"完成"声明前（在 testing-strategy 之后） |
| `../../skills/deployment-and-ops.md` | 上线执行阶段，部署 / 冒烟 / 监控 / 回滚 |
| `../../skills/seo-analytics-compliance.md` | 上线前，SEO / 统计 / 合规检查 |
| `../../skills/user-feedback-iteration.md` | 上线后，反馈收集与 30 天迭代 |

## 为什么用引用而不是复制

1. **单一来源**：知识库 skills/ 是唯一来源，避免多处不同步
2. **自动继承**：知识库 skills/ 更新时，agent 自动继承
3. **反哺闭环**：agent 使用中发现新规则，反哺到知识库 skills/，agent 自动继承

## 移植时的处理

如果把 agent/ 移植到其他环境（如 WorkBuddy skills 目录或 DC Ops 平台）：
1. 把 `../../skills/` 的 18 个方法论文件复制到本目录
2. 把 SKILL.md 中的引用路径从 `../../skills/` 改为 `./`
