# 子能力引用说明

本目录不复制方法论内容，而是引用知识库的 `../../skills/` 目录。

## 引用的 9 个方法论

| Skill 文件 | 何时调用 |
|---|---|
| `../../skills/brainstorming.md` | 立项阶段，挖掘需求 |
| `../../skills/writing-plans.md` | 立项后，写开发计划 |
| `../../skills/project-breakdown.md` | 立项后，拆功能清单 |
| `../../skills/database-design.md` | 数据库阶段 |
| `../../skills/backend-architecture-acceptance.md` | 后端骨架验收 |
| `../../skills/backend-security-review.md` | 安全审查 |
| `../../skills/requesting-code-review.md` | 代码审查 |
| `../../skills/systematic-debugging.md` | 调试 Bug |
| `../../skills/browser-verification.md` | 前端浏览器验证 |
| `../../skills/verification-before-completion.md` | 任何"完成"声明前 |

## 为什么用引用而不是复制

1. **单一来源**：知识库 skills/ 是唯一来源，避免多处不同步
2. **自动继承**：知识库 skills/ 更新时，agent 自动继承
3. **反哺闭环**：agent 使用中发现新规则，反哺到知识库 skills/，agent 自动继承

## 移植时的处理

如果把 agent/ 移植到其他环境（如 WorkBuddy skills 目录或 DC Ops 平台）：
1. 把 `../../skills/` 的 9 个文件复制到本目录
2. 把 SKILL.md 中的引用路径从 `../../skills/` 改为 `./`
