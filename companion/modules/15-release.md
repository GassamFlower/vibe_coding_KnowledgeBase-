<!--
id: 15-release
stage: 15
type: phase
title: 阶段15 上线 + Skill 反哺
next: 项目结束，进入健康迭代 / 收尾
-->
# 阶段15：上线 + Skill 反哺（上线门 · 门禁校验）

目标：上线部署 + 复盘 + 反哺知识库。

动作清单：
- 执行部署方案，线上冒烟测试（核心流程跑一遍）
- 整理踩坑清单（从 PROJECT_STATUS 踩坑记录提取）
- 反哺 `skills/`：哪条规则没覆盖 → 加；哪条不够强 → 强化
- 反哺 Prompt 模板库
- 写复盘报告（Keep / Change / Add / 踩坑清单 / 下次提醒），每条来源标注
- 更新 PROJECT_STATUS「下次项目提醒」
- 调用 `gate-上线门` 自动验收 → 知识库 git commit + push
- 提交 Git：`git commit -m "chore: 项目上线 + Skill 反哺"`

提示：上线完就结束不反哺 = 踩过的坑下次继续踩。复盘只写成绩不写问题 = 自欺。
