# VibeVibe 工程手册索引

> 用途：把外部工程手册（VibeVibe）的章节对到本知识库的阶段、门禁和 Skill，供查证时按需加载。
> 本文件只记录章节定位、适用场景和对照关系，不复制原文。
> 来源：https://www.vibevibe.cn/zh/
> 许可：VibeVibe 仓库使用 CC BY-NC-SA 4.0；本知识库使用 Apache 2.0。引用时保留来源链接，不整段搬运原文。

---

## 使用方式

1. 先确定当前阶段（14 阶段模型）或当前门禁。
2. 在下表找到对应章节，打开原站章节阅读。
3. 需要实施方法时，回到本知识库的 Skill 和门禁卡；外部手册只作为查证和补充。
4. 外部手册更新后，本索引只维护对照关系，不复制正文。

---

## 一、进阶篇（Advanced，16 章）对照

16 章的逐章核心、关键做法、81 个子页要点与延伸页见 [VibeVibe进阶篇索引](VibeVibe进阶篇索引.md)。

| 章节 | 标题 | 对应阶段 | 对应门禁 | 对应 Skill / 模板 |
|---|---|---|---|---|
| [01](https://www.vibevibe.cn/zh/Advanced/01-environment-setup/) | 环境搭建与代码运行基础 | 阶段 2 技术选型 | 立项门 | `docs/AI协作全流程操作手册.md`（技术选型与运行环境） |
| [02](https://www.vibevibe.cn/zh/Advanced/02-ai-tuning-guide/) | AI 使用说明书 | 全阶段 | 全门禁 | `skills/brainstorming.md`、`skills/systematic-debugging.md` |
| [03](https://www.vibevibe.cn/zh/Advanced/03-prd-doc-driven/) | 产品思维与文档驱动 | 阶段 1 立项 | 立项门 | `skills/product-validation.md`、`skills/project-breakdown.md`、`templates/立项文档PRD模板.md` |
| [04](https://www.vibevibe.cn/zh/Advanced/04-dev-fundamentals/) | 你必须知道的开发基础 | 阶段 2-3 | 立项门 | `docs/AI协作全流程操作手册.md`（技术选型与架构设计） |
| [05](https://www.vibevibe.cn/zh/Advanced/05-ui-ux/) | 界面(UI)与交互(UX) | 阶段 6 前端骨架 | 架构门 | `skills/ui-design-system.md`、`references/UIUX参考库索引.md` |
| [06](https://www.vibevibe.cn/zh/Advanced/06-data-persistence-database/) | 数据持久化与数据库 | 阶段 7 数据库设计 | 架构门 | `skills/database-design.md` |
| [07](https://www.vibevibe.cn/zh/Advanced/07-backend-api/) | 后端 API 开发 | 阶段 8 后端骨架 | 架构门 | `skills/api-design.md`、`templates/API接口清单模板.md` |
| [08](https://www.vibevibe.cn/zh/Advanced/08-auth-security/) | 安全与用户认证 | 阶段 12 安全审查 | 业务门 | `skills/backend-security-review.md`、`skills/frontend-security-review.md` |
| [09](https://www.vibevibe.cn/zh/Advanced/09-testing-automation/) | 功能测试与自动化 | 阶段 11 业务模块开发 | 业务门 | `skills/testing-strategy.md`、`skills/verification-before-completion.md` |
| [10](https://www.vibevibe.cn/zh/Advanced/10-localhost-public-access/) | Localhost 与公网访问 | 阶段 9 后端联调、阶段 14 上线 | 架构门 / 上线门 | `skills/deployment-and-ops.md` |
| [11](https://www.vibevibe.cn/zh/Advanced/11-git-collaboration/) | Git 版本控制与协作开发 | 全阶段 | 全门禁 | 各门禁的 Git 提交检查 |
| [12](https://www.vibevibe.cn/zh/Advanced/12-serverless-deploy-cicd/) | 无服务器部署与 CI/CD 自动化 | 阶段 14 上线 | 上线门 | `skills/deployment-and-ops.md`、`templates/上线检查清单.md` |
| [13](https://www.vibevibe.cn/zh/Advanced/13-domain-dns/) | 域名、DNS 与网络接入 | 阶段 14 上线 | 上线门 | `skills/deployment-and-ops.md` |
| [14](https://www.vibevibe.cn/zh/Advanced/14-vps-ops-deploy/) | 云服务器运维与项目部署 | 阶段 14 上线 | 上线门 | `skills/deployment-and-ops.md` |
| [15](https://www.vibevibe.cn/zh/Advanced/15-seo-analytics/) | SEO、分享与数据统计 | 阶段 13 上线前总验收 | 上线门 | `skills/seo-analytics-compliance.md` |
| [16](https://www.vibevibe.cn/zh/Advanced/16-user-feedback-iteration/) | 用户反馈与产品迭代 | 阶段 14 上线后观察 | 上线门 | `skills/user-feedback-iteration.md`、`templates/上线后30天迭代计划.md` |

---

## 二、入门篇（Basic）对照

10 个页面的逐章核心、关键做法、36 个子页要点见 [VibeVibe入门篇索引](VibeVibe入门篇索引.md)。

| 章节 | 标题 | 适用场景 |
|---|---|---|
| [00](https://www.vibevibe.cn/zh/Basic/00-preface/) | 开始前：你会做出什么，怎么学，卡住怎么办 | 第一次接触 Vibe Coding 的学习路径 |
| [01](https://www.vibevibe.cn/zh/Basic/01-awakening/) | 第一个版本：2 小时做出个人主页与数字分身 | 建立"先做出一个能跑的东西"的信心 |
| [02](https://www.vibevibe.cn/zh/Basic/02-mindset/) | 带回自己的工作台：从平台到本地 | 阶段 1-2 的心态与工作方式 |
| [03](https://www.vibevibe.cn/zh/Basic/03-technique/) | 做好第一印象：界面、风格与更有效的需求表达 | 阶段 1 需求表达、阶段 6 视觉方向 |
| [04](https://www.vibevibe.cn/zh/Basic/04-practice-0-to-1/) | 让主页更完整：内容、引导与基础存档 | 阶段 11 业务模块的第一轮实践 |
| [05](https://www.vibevibe.cn/zh/Basic/05-advanced/) | 让数字分身更像你：人设、回答与排错 | AI 配置与调试入门 |
| [06](https://www.vibevibe.cn/zh/Basic/06-launch/) | 正式上线：部署、分享与第一轮真实反馈 | 阶段 14 上线与首批反馈 |
| [99](https://www.vibevibe.cn/zh/Basic/99-appendix/) | 附录（提示词汇总等） | Prompt 速查 |
| [100](https://www.vibevibe.cn/zh/Basic/100-epilogue/) | 结语：给未来的你 | 第一次闭环收束与下一步选择 |
| [101](https://www.vibevibe.cn/zh/Basic/101-next-part/) | 下部预告：Vibe Coding 前后端实战教程 | 从入门篇进入进阶篇的能力地图 |

---

## 三、实践案例篇（Practice）对照

5 篇的标题、项目背景、核心做法、要点与适用阶段见 [VibeVibe实践案例索引](VibeVibe实践案例索引.md)。

| 案例 | 适用场景 |
|---|---|
| [AI 画布实践](https://www.vibevibe.cn/zh/Practice/ai-canvas-vibecoding-journey/) | 从想法到上线的完整历程参考 |
| [AI 绘本生成器](https://www.vibevibe.cn/zh/Practice/ai-picture-book-generator/) | 零代码工具类项目参考 |
| [AI 简历优化 SaaS](https://www.vibevibe.cn/zh/Practice/ai-resume-saas/) | 前后端一体 SaaS 的目录、接口、支付与部署参考 |
| [21 天思考习惯小游戏](https://www.vibevibe.cn/zh/Practice/full-stack-21-day-thinking-habit-game/) | 前后端一体小游戏的阶段划分与验收参考 |
| [Vibe 编程经验分享](https://www.vibevibe.cn/zh/Practice/vibe-coding-methodology/) | 从执行者到决策者的协作方式 |

---

## 四、文章篇（Articles，6 类）对照

41 篇的标题、核心论点、要点与适用阶段见 [VibeVibe优质文章索引](VibeVibe优质文章索引.md)。

| 分类 | 标题 | 适用场景 |
|---|---|---|
| [01](https://www.vibevibe.cn/zh/Articles/01-core-concepts/) | 核心概念与范式演进 | 理解 Vibe Coding、智能体工程的基本概念 |
| [02](https://www.vibevibe.cn/zh/Articles/02-technical-architecture/) | 基础技术与架构设计 | 智能体架构、MCP、多智能体系统 |
| [03](https://www.vibevibe.cn/zh/Articles/03-toolchain-frameworks/) | 工具链与开发框架 | 开发工具、Skills、规范写法 |
| [04](https://www.vibevibe.cn/zh/Articles/04-engineering-practices/) | 工程实施与质量保障 | 代码审查、测试证明、团队协作 |
| [05](https://www.vibevibe.cn/zh/Articles/05-security-compliance/) | 安全、合规与局限性 | 阶段 12 安全审查、阶段 13 合规检查 |
| [06](https://www.vibevibe.cn/zh/Articles/06-business-trends/) | 商业应用与行业趋势 | 立项阶段判断市场与商业模式 |

---

## 五、维护规则

- 只维护对照关系，不复制外部章节正文。
- 外部章节改名或移动时，更新本文件链接。
- 新增 Skill 或模板后，在本表补对应行。
- 外部手册的方法与本知识库规则冲突时，以本知识库门禁卡为准，并在复盘报告里记录差异。
