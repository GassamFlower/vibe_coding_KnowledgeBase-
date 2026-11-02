# VibeVibe 优质文章索引

> 用途：把 vibevibe.cn「优质文章」41 篇整理成逐篇摘要，供立项、架构、开发、安全、上线与复盘时按主题取用。
> 来源：https://www.vibevibe.cn/Articles/
> 许可：VibeVibe 使用 CC BY-NC-SA 4.0；本文件只写摘要、要点与原文链接，不复制正文。
> 全文缓存：`.audit/crawl/text/`（本机目录，已被 `.gitignore` 忽略，不入库）。
> 整理日期：2026-11-02。

## 怎么用

1. 先按分类找主题，读「核心」判断是否与当前问题相关。
2. 「要点」给出文中明确写出的做法、判断标准或结论；需要上下文时打开原文。
3. 「适用」对应知识库的阶段与 Skill，进入对应阶段时与门禁卡一起使用。
4. 同一篇文章出现在两个分类时，只保留一处完整摘要，另一处标注对应关系。

## 分类

| 分类 | 篇数 | 主要用途 |
|---|---|---|
| [01 核心概念与范式演进](#01-核心概念与范式演进) | 10 | 理解智能体、规范驱动与 Vibe Coding 的边界 |
| [02 技术架构](#02-技术架构) | 7 | 智能体架构、MCP、多智能体与工具取舍 |
| [03 工具链与开发框架](#03-工具链与开发框架) | 7 | 原型工具、Skills、Cline 与 AGENTS.md |
| [04 工程实践](#04-工程实践) | 12 | 代码审查、规范、测试、团队与工作流 |
| [05 安全与合规](#05-安全与合规) | 2 | 智能体威胁模型与安全边界 |
| [06 商业与趋势](#06-商业与趋势) | 3 | 行业变化、社区运营与项目案例 |

---

## 01 核心概念与范式演进

> 分类页：https://www.vibevibe.cn/Articles/01-core-concepts/

### 1.1 智能体工程

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/agentic-engineering.html>
- 核心：把严谨的 AI 协作称为智能体工程：人负责架构、审查与正确性，AI 负责推进实现，工程纪律回到中心位置。
- 要点：
  - 动手前先写设计文档或规范，把工作分成边界清晰的任务再交给 AI。
  - 把测试当成主干，建立稳定验证闭环，AI 才能在明确目标下迭代。
  - AI 产出按团队成员 PR 的标准逐条审查，讲不清的实现不进代码库。
  - 架构、版本控制、持续集成与生产可观测性始终由人承担，责任不转移。
- 适用：阶段 4 Agent 宪法 / writing-plans、requesting-code-review

### 1.2 Coding Agents 入门：真正完成工作的艺术

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/coding-agents-101.html>
- 核心：把 Coding Agents 当成需要明确上下文的搭档，交代做法、给出反馈回路、设置检查点，中大型任务才委托得动。
- 要点：
  - 提示里说明先看哪个仓库、文档和关键组件，减少无效探索。
  - 开放持续集成、类型检查器、代码检查器和测试，让代理依据报错自行修正。
  - 多部分任务按计划、实现、测试、修复、审查设置检查点再推进。
  - 代理开始绕圈就停止对话或手动接管，必要时开新会话重新交代。
- 适用：阶段 11 业务模块开发 / writing-plans、testing-strategy

### 1.3 我们在 FAANG 是怎么做 Vibe Coding 的

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/how-we-vibe-code-at-faang.html>
- 核心：生产环境里 AI 只是加速既有流程：设计文档、设计审查、子系统文档、任务分解、测试先行、双人审查逐项保留。
- 要点：
  - 一切从技术设计文档开始，方向获认可后才补齐架构与依赖关系。
  - 设计通过后按子系统补文档，写清接口、职责划分与交付边界。
  - 采用测试驱动开发，先补测试再让代理实现，任务边界更清楚。
  - 代码经两位开发者批准才合并；原文称周期缩短约 30%。
- 适用：阶段 3 架构设计 / writing-plans、requesting-code-review

### 1.4 规范是新的源代码

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/specs-are-the-new-source-code.html>
- 核心：AI 让实现速度不再是瓶颈，写得足够精确的规范成为团队与模型的共同真源，代码只是规范的下游产物。
- 要点：
  - 先用 v0、Lovable、Replit 快速做出原型，收集真实客户反馈。
  - 拿到反馈之后写出清晰规范，再交给 AI 实施，原型让规范更好。
  - 规范模糊会产出混乱代码，可让 AI 起草工单并审查代码库使其具体。
  - 非技术人员可借 MCP 与 Copilot 从工单走到 PR，仍需工程师守门。
- 适用：阶段 5 实施真源文档 / writing-plans、user-feedback-iteration

### 1.5 工厂模型：Coding Agents 如何改变软件工程

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/the-factory-model-how-coding-agents-changed-software-engineering.html>
- 核心：开发者的工作转为构建稳定产出软件的工厂：编排智能体、写好规范、投资测试，并严格验证每一份输出。
- 要点：
  - 向二十个以上并行智能体下发任务时，规范质量决定产出差异。
  - 红绿 TDD：先写测试并确认失败，再迭代实现直到测试通过。
  - 生成不再是瓶颈，验证才是，人工审查仍承担安全系统角色。
  - 用新工程师能否凭文档和提交历史理解代码来检验投入方向。
- 适用：阶段 5 实施真源文档 / testing-strategy、verification-before-completion

### 1.6 Vibe Coding 不是低质量工作的借口

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/vibe-coding-is-not-an-excuse-for-low-quality-work.html>
- 核心：AI 加速实现不等于质量自动达标，未经审查的生成代码不得进入代码库，人类必须始终留在回路里。
- 要点：
  - 逐条阅读 AI 产出，确认自己能解释结构与意图，再决定是否保留。
  - 把大块生成代码整理为更小模块，补上边缘情况与类型约束。
  - 关键路径必须有测试，AI 生成的测试也要检查是否为表面文章。
  - 企业级系统、关键算法与安全模块不适合完全交给 AI 生成。
- 适用：阶段 13 上线前验收 / requesting-code-review、verification-before-completion

### 1.7 Vibe Coding：革命还是鲁莽放纵？

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/vibe-coding-revolution-or-reckless-abandon.html>
- 核心：梳理 Vibe Coding 从周末项目到生产争议的全貌，结论是原型可以凭感觉推进，进入生产仍须纪律与人工审查。
- 要点：
  - Karpathy 原意是不读代码、只凭运行结果反复迭代的极端方式。
  - 爱好者用它做原型与探索，安全与可维护性风险集中在生产环节。
  - 原文称 Vibe Coding 大约能推进到功能概念的 80%，余下部分靠工程师。
  - 稳妥做法是原型阶段放开，上线之前走完代码审查与测试。
- 适用：阶段 1 立项 / product-validation、requesting-code-review

### 1.8 什么是 AI Agent？为什么它们很重要？

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/what-are-ai-agents-why-do-they-matter.html>
- 核心：Agent 能感知环境、保持上下文、调用工具并依据结果调整行动，与单次问答模型的差别在自主执行与目标导向。
- 要点：
  - Agent 可保持持久上下文，把复杂目标分解为子任务并监控进展。
  - MCP 负责把 Agent 接到工具，A2A 负责 Agent 之间的通信协作。
  - 分层规划分为战略、战术、执行、行动四层，模仿专家处理问题。
  - 幻觉、错误传播与访问控制仍是当前 Agent 的主要限制。
- 适用：阶段 3 架构设计 / api-design、brainstorming

### 1.9 关于 Vibe Coding 你需要知道的一切

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/what-you-need-to-know-about-vibe-coding.html>
- 核心：整理 Vercel 发布的现状报告：超过九成美国开发者使用 AI 编程工具，非开发者大量进入，平台须把安全防护内置进工具。
- 要点：
  - 工具从拖拽组件搭建进入描述需求即生成应用的阶段。
  - 原文称约 63% 探索这类工具的用户是非开发者。
  - 平台需实时扫描并阻断漏洞，在部署阶段拦下高风险问题。
  - 企业可在关键系统保留约束更强的流程，只把 AI 用于可控环节。
- 适用：阶段 1 立项 / product-validation、deployment-and-ops

### 1.10 如何成为世界级的智能体工程师

- 原文：<https://www.vibevibe.cn/Articles/01-core-concepts/world-class-agent-engineer.html>
- 核心：配置保持精简，精确控制上下文，用测试与截图定义任务终点，用独立会话和完成清单约束长时间运行。
- 要点：
  - 把研究与实现分开，包含实现细节的请求能明显减少幻觉。
  - 用中性提示避免模型迎合预设结论，必要时用多方评审交叉验证。
  - 以测试全部通过作为任务终点，禁止代理修改这些测试。
  - 用停止钩子与任务完成清单约束长会话，每份清单开新会话执行。
- 适用：阶段 4 Agent 宪法 / writing-plans、verification-before-completion

---

## 02 技术架构

> 分类页：https://www.vibevibe.cn/Articles/02-technical-architecture/

### 2.1 用文件系统和 Bash 构建智能体

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/build-agents-with-filesystems-and-bash.html>
- 核心：用文件系统工具与 Bash 替换智能体系统里大部分定制工具，让模型凭已有的目录导航与检索能力处理领域数据，成本下降且输出质量提升。
- 要点：
  - 把客户记录、工单、销售通话等数据按目录层级组织成文件，供智能体按需读取。
  - 智能体在沙箱里先用 ls、find 探索，再用 grep、cat 定位，只把必要内容送给模型。
  - grep 返回明确匹配，比向量检索更适合查找具体值，上下文也能保持最小。
  - 销售电话总结智能体单次通话成本从约 1 美元降到约 0.25 美元。
- 适用：阶段 3 架构设计 / brainstorming

### 2.2 揭秘 Claude Code 的工作原理

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/how-claude-code-works.html>
- 核心：Claude Code 的效果来自超长系统提示、设计良好的工具描述、系统化的上下文工程与关键时刻插入的提醒，与模型本身的强弱关系不大。
- 要点：
  - 执行任务前先总结对话、识别主题与项目目标，建立清楚的上下文边界。
  - 用 `<system-reminder>` 标签在关键节点注入提醒，让模型在长上下文里持续关注目标。
  - 安全判断由生成式提示完成，命令前缀提取与注入检测嵌进主循环。
  - 主智能体编排多步骤任务，子智能体处理具体单一目标；工具描述写明禁止项与替代路径。
- 适用：阶段 4 Agent 宪法 / writing-plans

### 2.3 MCP：它是什么，以及为什么重要

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/mcp-what-it-is-and-why-it-matters.html>
- 核心：MCP 是统一 AI 与外部工具连接方式的开放标准，把零散定制集成收拢到协议层，工具实现一次即可被所有支持的客户端复用。
- 要点：
  - 架构含 MCP 服务器、客户端、协议与服务四部分，协议与传输层解耦。
  - 工具按统一规范实现一次即可复用，N 个工具乘 M 个平台的集成矩阵被压缩。
  - 入门先用官方目录中的现成服务器，专有场景再用 SDK 自建并以 MCP Inspector 联调。
  - 限制包括认证标准仍在演进、生态处于早期、多一层通信带来延迟与资源开销。
- 适用：阶段 2 技术选型 / api-design

### 2.4 多智能体系统如何支撑 AI 原生工程

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/multi-agent-systems-ai-native-engineering.html>
- 核心：生产系统的故障排查需要多个专业维度同时介入，只有多智能体系统能并行验证假设；AI 原生工程要让 AI 成为处理工作的主入口。
- 要点：
  - AI 原生工程指工程师通过 AI 接口协调工作，由 AI 组织调查、收集证据、执行动作。
  - 从 LLM 到 LLM 加工具、单一智能体、多智能体逐层演进，每层都有可扩展性上限。
  - 多智能体要处理上下文传播、并行协调、冲突解决、失败重试与持续学习。
  - 领域知识决定智能体边界，AI 工程能力决定多个智能体能否协同。
- 适用：阶段 3 架构设计 / systematic-debugging

### 2.5 理解 AI 的「五层蛋糕」架构

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/nvidia-ai-5-layer-cake.html>
- 核心：NVIDIA 把 AI 体系分为芯片、系统、软件平台、基础模型与 AI 应用五层，上层能力受下层发展水平限制。
- 要点：
  - 芯片层决定计算上限，GPU 的数千个核心支撑训练与推理的并行计算。
  - 系统层把多个 GPU 组装成服务器节点，再用高速网络连接成集群。
  - 软件平台层含 CUDA、cuDNN、TensorRT 等，深度学习框架建立在其上。
  - 应用层创新速度最快，天花板取决于所依赖模型、平台与硬件的水平。
- 适用：阶段 2 技术选型 / brainstorming

### 2.6 自我改进的编码智能体

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/self-improving-coding-agents.html>
- 核心：把开发工作切成小任务，让编码智能体在每轮隔离的循环里实现、验证、提交，并用外部文件传递上下文与经验。
- 要点：
  - 每轮从任务列表取一项，实现后跑测试与类型检查，通过后提交并更新任务状态。
  - 任务要小到一次会话可完成，并有明确的通过或失败标准。
  - 用 AGENTS.md 记录约定、易错点与最近变更，条目要短、准确、可直接执行。
  - 循环只在功能分支或隔离环境中运行，写操作、发布与删除保留人工确认。
  - 设置最大迭代数、最长运行时间等停止条件，避免空转与令牌消耗失控。
- 适用：阶段 4 Agent 宪法 / verification-before-completion、testing-strategy

### 2.7 测试「Bash 就够了」假设：SQL vs Bash 智能体对比实验

- 原文：<https://www.vibevibe.cn/Articles/02-technical-architecture/testing-bash-vs-sql.html>
- 核心：在查询 GitHub issues 的对比实验中，SQL 智能体准确率 100%，Bash 只有 52.7%，两者结合的混合方案靠自我验证稳定达到 100%。
- 要点：
  - Bash 智能体消耗约 7 倍令牌，成本高出 6.5 倍，耗时接近 9 倍。
  - 只提供基础文件工具的智能体准确率 63%，高于开放完整 Shell 的 52.7%。
  - 混合方案先跑 SQL 再回文件系统用 grep 交叉验证，令牌约为纯 SQL 的两倍。
  - schema 清楚的结构化数据用 SQL 最直接，Bash 适合探索、抽样核对与覆盖边缘情况。
- 适用：阶段 2 技术选型 / testing-strategy、database-design

---

## 03 工具链与开发框架

> 分类页：https://www.vibevibe.cn/Articles/03-toolchain-frameworks/

### 3.1 AGENTS.md vs Skills：被动上下文为何胜过主动检索

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/agents-md-vs-skills.html>
- 核心：Vercel 用评估测试对比 Next.js 文档的两种投喂方式，发现始终可用的 AGENTS.md 索引比按需调用的 Skill 更可靠，通过率从 53% 升到 100%。
- 要点：
  - 默认状态下 Skill 在 56% 的评估案例中从未被调用，通过率与基线同为 53%。
  - 在 AGENTS.md 中写明先探索项目再调用 Skill，触发率超过 95%，通过率提升到 79%。
  - 直接嵌入文档索引并提示优先采用检索引导的推理，通过率达到 100%。
  - 索引从 40KB 压缩到 8KB 仍保持 100% 通过率，避免上下文膨胀。
- 适用：阶段 4 Agent 宪法 / verification-before-completion

### 3.2 AI 驱动的原型开发：v0、Bolt 和 Lovable 对比

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/ai-driven-prototyping-v0-bolt-and-lovable-compared.html>
- 核心：Addy Osmani 从工程视角比较 v0、Bolt、Lovable 三款原型工具的能力与取舍，指出引导式工具做出原型后通常需要导出代码回到本地编辑。
- 要点：
  - v0 以 React 与 shadcn/ui 为中心，擅长快速界面原型，并提供 npx 命令把组件装进项目。
  - Bolt 用浏览器内环境做前后端一体的原型，支持文件锁定与提示文件来约束改动范围。
  - Lovable 采用引导式路线，原生接入 Supabase，GitHub 协作功能较强，自定义上限较低。
  - 三者都有令牌用量限制，复杂度越过阈值后需要导出到本地工具继续开发。
- 适用：阶段 1 立项 / product-validation、ui-design-system

### 3.3 Claude Code Swarms：多智能体协作开发

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/claude-code-swarms.html>
- 核心：介绍 Claude Code 的智能体团队功能，主智能体把工作委派给多个独立实例并行推进，通过共享任务列表和收件箱消息协作。
- 要点：
  - 在 settings.json 中把 CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS 设为 1 启用团队。
  - 委派模式按 Shift+Tab 把负责人限制为只做协调，阻止其自己动手写代码。
  - 任务列表带依赖跟踪和自动解锁，认领任务时用文件锁定防止竞态。
  - 适合并行调试竞争假设、多视角代码审查、跨层功能；顺序任务用单个会话更省令牌。
- 适用：阶段 11 业务模块开发 / requesting-code-review

### 3.4 Claude Skills 完整构建指南

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/complete-guide-to-building-skills-for-claude.html>
- 核心：Anthropic 给出构建 Skill 的完整方法，用渐进式披露组织文件夹，用 YAML frontmatter 决定加载时机，再通过测试与迭代改进。
- 要点：
  - Skill 文件夹必须包含 SKILL.md，可选 scripts、references、assets 三个目录，不放 README.md。
  - description 要同时写清做什么与何时使用，少于 1024 个字符，并包含用户会说的触发短语。
  - 测试覆盖三个领域：触发测试、功能测试、与基线比较的性能测试。
  - 指令未被遵循时，把关键校验写成脚本执行检查，代码比语言说明更确定。
- 适用：阶段 4 Agent 宪法 / writing-plans

### 3.5 如何为 AI 智能体写出优秀的规范

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/how-to-write-good-spec-for-ai-agents.html>
- 核心：Addy Osmani 提出写作智能体规范的五条原则，从高层愿景起步、按 PRD 结构组织、把任务分为模块化提示、设置三层边界与自检。
- 要点：
  - 先用计划模式让智能体在只读状态起草方案，架构与约束讲清后再切换到执行。
  - 最有效的规范覆盖六个领域：命令、测试、项目结构、代码风格、Git 工作流、边界。
  - 采用三层边界：始终做、先询问、永远不要，比单层禁令清单更清晰。
  - 让智能体为长规范生成带摘要的扩展目录，作为可长期保留的提纲。
- 适用：阶段 5 实施真源文档 / writing-plans

> 同文提示：本文与 [4.5](#45-如何为-ai-coding-agents-编写高质量规范) 为同一篇文章的两个版本，完整摘要保留在本节。

### 3.6 我们删掉了智能体 80% 的工具

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/we-removed-80-percent-tools.html>
- 核心：Vercel 把 text-to-SQL 智能体的十几个专用工具换成单个 bash 执行工具，成功率从 80% 升到 100%，步骤与令牌都减少。
- 要点：
  - 只保留执行任意 bash 命令与执行 SQL 两个工具，其余专用工具全部删除。
  - 五个对比查询上平均耗时从 274.8 秒降到 77.4 秒，平均令牌减少 37%。
  - 前提是语义层本身清楚：YAML 文件结构一致、命名统一、定义明确。
  - 建议从模型加文件系统加目标的最简结构起步，只在证明必要时增加复杂度。
- 适用：阶段 3 架构设计 / backend-architecture-acceptance

### 3.7 为什么我在 AI 工程中使用 Cline

- 原文：<https://www.vibevibe.cn/Articles/03-toolchain-frameworks/why-i-use-cline-for-ai-engineering.html>
- 核心：Addy Osmani 说明选择 Cline 的理由，模型可自由切换、检查点版本控制、浏览器与终端运行时交互、计划与执行模式分离。
- 要点：
  - 上下文用 @file、@folder、@url、@problems 四种方式按需注入，避免压满窗口。
  - 规划用 DeepSeek-R1 或 Gemini，实施用 Claude Sonnet，成本最高可降低 97%。
  - 检查点保存完整环境状态，独立于 Git 运行，避免实验性改动污染提交历史。
  - 计划模式先设计与审查方案，执行模式直接实现，两类模型设置可全局保存。
- 适用：阶段 2 技术选型 / browser-verification

---

## 04 工程实践

> 分类页：https://www.vibevibe.cn/Articles/04-engineering-practices/

### 4.1 AI 代码审查的实施与最佳实践

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/ai-code-review-implementation.html>
- 核心：把 AI 代码审查接入提交拉取请求的流程，让机器承担例行检查，人工复核高风险判断，并用反馈循环持续收窄误报。
- 要点：
  - 配置仓库钩子与 webhook，在提交拉取请求时自动触发审查，并按严重程度分级。
  - 明确 AI 与人工的分工：风格一致、基础逻辑错误、安全扫描交给 AI。
  - 人工复核 AI 建议并记录接受与拒绝情况，据此调整敏感性设置与忽略模式。
  - 涉及用户输入、身份验证、数据库查询、文件操作的代码需要额外安全检查。
- 适用：阶段 12 安全审查 / requesting-code-review

### 4.2 AI 写代码更快，你的工作是证明它能用

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/ai-writes-code-faster-your-job-is-to-prove-it-works.html>
- 核心：生成代码变快之后，举证责任转到提交方：变更必须附带能工作的证据，代码审查用来判断风险、意图与责任归属。
- 要点：
  - 拉取请求写清四件事：什么与为什么、能工作的证明、风险与 AI 参与部分、审查重点。
  - 要求代理在生成后执行代码或运行单元测试，附上日志、截图或结果。
  - 把代理产出整理成便于阅读的提交，避免一次交付过大的变更。
  - 身份验证、支付、密钥与不受信任输入必须增加人工威胁建模。
- 适用：阶段 13 上线前验收 / verification-before-completion、requesting-code-review

### 4.3 在 AI 时代避免技能退化

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/avoiding-skill-atrophy-in-the-age-of-ai.html>
- 核心：长期依赖 AI 助手会让批判性思维、调试与架构判断逐渐退化，作者给出可执行的练习方式把这些能力保持在可用状态。
- 要点：
  - 微软与卡内基梅隆 2025 年研究发现，越依赖 AI 工具，投入的批判性思维越少。
  - 每周保留一天「无 AI 日」，从头编写代码、完整阅读错误、查阅文档。
  - 向 AI 提问之前先自己调查 15 到 30 分钟，再对照它的答案学习。
  - 记录经常求助 AI 的主题，把它当作需要补齐的知识缺口。
- 适用：阶段 11 业务模块开发 / systematic-debugging、verification-before-completion

### 4.4 超越 70%：最大化 AI 辅助编程中人类 30% 的价值

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/beyond-the-70-maximizing-the-human-30-of-ai-assisted-coding.html>
- 核心：AI 能完成约七成的模式化工作，剩余三成涉及系统设计、边缘情况与质量判断，人类应把精力集中在这个部分。
- 要点：
  - AI 擅长处理偶然复杂性，本质复杂性仍然需要人类承担。
  - 边缘情况、可维护性与稳固架构需要人类的专业判断。
  - 把 AI 产出当作初级开发者的成果，审查与测试通过之后才接受。
  - 高级开发者用经验筛选 AI 建议，初级开发者先打牢基础知识。
- 适用：阶段 10 架构验收 / testing-strategy、backend-architecture-acceptance

### 4.5 如何为 AI Coding Agents 编写高质量规范

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/how-to-write-a-good-spec-for-ai-agents.html>
- 核心：同一篇文章的另一个版本；规范写法、三层边界与摘要目录的完整摘要见 [3.5](#35-如何为-ai-智能体写出优秀的规范)。
- 要点：
  - 两篇标题与译文不同，正文结构与结论一致。
  - 需要引用时优先使用 3.5 的链接与摘要，避免两处口径分叉。
- 适用：阶段 5 实施真源文档 / writing-plans

### 4.6 GenAI 时代领导高效工程团队

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/leading-effective-engineering-teams-in-the-age-of-genai.html>
- 核心：领导者要定义「更好」的含义，指导团队把 AI 当作需要监督的初级成员，同时守住人类判断与核心技能。
- 要点：
  - 采用「信任但验证」，AI 生成的代码经过与人类代码同等或更严的审查。
  - 用无 AI 日、要求解释推理等方式，防止团队成员技能退化。
  - 制定书面使用政策，写明批准的工具、数据处理规则、审查要求与归属。
  - 衡量长期代码质量与知识保留，避免只看交付速度。
- 适用：阶段 4 Agent 宪法 / requesting-code-review

### 4.7 我在 2026 年的 LLM 编码工作流

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/my-llm-coding-workflow-going-into-2026.html>
- 核心：作者以一年多项目实践总结与 LLM 协作的方式：先把规格说明与分步计划写清楚，把任务切小，补足上下文，全程由人审查、测试并频繁提交。
- 要点：
  - 先让模型反过来提问，直到需求、边界条件与关键约束说透，再整理成 spec.md。
  - 一次只推进一个函数、一个缺陷或一个功能，写完测完再进入下一步。
  - 把相关代码与官方文档主动提供，可借助 Context7、gitingest 打包上下文。
  - 模型产出按初级开发者成果对待，逐行审查、运行测试，用小步提交与 worktree 隔离实验。
- 适用：阶段 5 实施真源文档 / writing-plans、verification-before-completion

### 4.8 React 最佳实践：Vercel 十年性能优化经验总结

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/react-best-practices.html>
- 核心：Vercel 把十多年 React 与 Next.js 性能优化经验整理成按影响力排序的规则库，供 AI Agent 审查代码与提出优化时查阅。
- 要点：
  - 优化先从消除异步瀑布流与减小 Bundle 体积入手，再处理重渲染等微观性能项。
  - 共八个类别、四十余条规则，每条标注 CRITICAL 至 LOW 的影响评级并提供代码示例。
  - 规则文件编译成 AGENTS.md，作为智能体整理代码时一致遵循的统一依据。
  - 规则来自生产环境真实优化，如合并循环遍历、并行 await、惰性状态初始化。
- 适用：阶段 6 前端骨架 / requesting-code-review、verification-before-completion

### 4.9 停止为 AGENTS.md 使用 /init

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/stop-using-init-for-agents.html>
- 核心：自动生成的 AGENTS.md 常常重复智能体本可自行发现的仓库信息，增加成本并稀释注意力，只应写入无法从代码推断的约束。
- 要点：
  - 判断标准是代理能否通过阅读代码自行发现，若能则删掉该行。
  - 值得写入的如包管理工具选择、测试必须带某参数、禁改目录与自定义中间件模式。
  - 代理反复出同类错误时，先修正代码结构、lint 规则与测试覆盖率，最后才动上下文文件。
  - 单体文件无法按任务类型加载，可按协议文件、角色与技能文件、维护子代理分层。
- 适用：阶段 4 Agent 宪法 / writing-plans

### 4.10 Vercel 如何构建 Coding Agents 的 AEO 追踪系统

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/vercel-aeo-tracking.html>
- 核心：Vercel 为每个 Coding Agent 单独启动沙箱运行，统一捕获并规范化运行记录，用来追踪模型与智能体如何发现和引用自家内容。
- 要点：
  - 生命周期分六步：创建沙箱、安装 CLI、注入网关凭证、运行智能体、捕获记录、清理沙箱。
  - 所有 LLM 调用经 AI Gateway 路由，获得统一日志、限速与成本追踪。
  - 规范化分四阶段：捕获记录、解析工具名与消息结构、提取元数据、汇总并抽取品牌。
  - 早期采样显示编程智能体约两成提示会执行网络搜索，推荐多以可运行代码形式出现。
- 适用：阶段 14 上线后迭代 / seo-analytics-compliance

### 4.11 使用 AI Gateway 生成视频

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/video-generation-with-ai-gateway.html>
- 核心：AI Gateway 新增视频生成能力，可用 AI SDK 6 编程调用或在 Playground 试用，覆盖四家提供商模型与多种生成方式。
- 要点：
  - 初步支持文本生成视频、图片生成视频、首尾帧、参考生成视频四类，另有视频编辑与音频。
  - 首批模型为 Grok Imagine、Wan、Kling、Veo，共十七个变体，功能处于 Beta 阶段。
  - 视频提示词需描述镜头移动、主体动作与时间变化，各家特有能力通过 providerOptions 传入。
  - 多参考对象用 character1、character2 等角色名对应，Wan 官方建议多参考生成时这样写。
- 适用：阶段 11 业务模块开发 / api-design

### 4.12 你的 AI Coding Agents 需要一个管理者

- 原文：<https://www.vibevibe.cn/Articles/04-engineering-practices/your-ai-coding-agents-need-a-manager.html>
- 核心：并行运行多个 Coding Agents 之后，工作重心转向编排，技术负责人的清晰度、委派、验证循环与异步沟通成为直接可用的能力。
- 要点：
  - 本地高接触会话处理架构决策与模糊需求，后台会话处理边界清晰的有界任务。
  - 简报需包含结果、上下文、约束、非目标、验收标准、集成说明与验证计划。
  - 验证循环要求代理运行测试与 lint 并附上输出，最后提交结构化 PR 包。
  - 一个代理只负责一个 PR，可能触碰同一文件时重新划分任务，用 worktree 隔离并行会话。
- 适用：阶段 11 业务模块开发 / requesting-code-review、verification-before-completion

---

## 05 安全与合规

> 分类页：https://www.vibevibe.cn/Articles/05-security-compliance/

### 5.1 AI 代理已经到来，威胁也随之而来

- 原文：<https://www.vibevibe.cn/Articles/05-security-compliance/ai-agents-threats-and-mitigations.html>
- 核心：研究用 CrewAI 与 AutoGen 构建投资顾问助手并对它们实施九类攻击，说明多数漏洞来自不安全的设计方式、错误配置与工具集成，代理安全需要分层纵深防御。
- 要点：
  - 九类攻击覆盖代理识别、指令提取、内网访问、卷数据外泄、SQL 注入与 BOLA。
  - 提示强化明确禁止提取系统提示与工具模式，内容过滤在运行时拦截注入尝试。
  - 代码解释器须放进限制网络访问与系统调用的最小权限沙箱。
  - 工具侧需要输入清理、严格授权检查，以及定期执行 SAST、DAST、SCA。
- 适用：阶段 3 架构设计、阶段 12 安全审查 / backend-security-review、frontend-security-review

### 5.2 智能体架构中的安全边界

- 原文：<https://www.vibevibe.cn/Articles/05-security-compliance/security-boundaries-in-agentic-architectures.html>
- 核心：文章区分智能体、智能体密钥、生成代码执行与文件系统四类参与者，主张把框架与生成代码放进独立的安全上下文运行，密钥在网络层按规则注入。
- 要点：
  - 所有组件共用同一安全上下文时，智能体可读到 .env、SSH 密钥与数据库凭证。
  - 密钥注入代理拦截出站流量，只在请求发往预期目标时附加凭证。
  - 共享沙箱只能隔离外部环境，生成代码仍可窃取框架持有的凭证。
  - 完整做法是应用沙箱配合密钥注入，注入的请求头覆盖沙箱自行设置的同名请求头。
- 适用：阶段 3 架构设计、阶段 12 安全审查 / backend-security-review

---

## 06 商业与趋势

> 分类页：https://www.vibevibe.cn/Articles/06-business-trends/

### 6.1 OpenEvidence 如何打造让医生信得过的医疗 AI

- 原文：<https://www.vibevibe.cn/Articles/06-business-trends/how-openevidence-built-healthcare-ai-physicians-trust.html>
- 核心：OpenEvidence 采用后端 Python 运行在 Google Cloud、前端 Next.js 部署在 Vercel 的混合架构，让一个很小的前端团队同时保住医院级可靠性与创业公司的迭代速度。
- 要点：
  - 每个概念验证项目独立部署并绑定自定义域名，利益相关方直接点进去体验工作流。
  - 每次提交自动部署，生产环境部署约 5 分钟，每个分支生成预览链接。
  - 启用 Fluid Compute 后无服务器支出下降 90%，业务规模扩大 1000 倍而可靠性未下降。
  - 前端团队规模一直很小，基础设施支出占比不到 5%。
- 适用：阶段 2 技术选型、阶段 13 上线前验收 / deployment-and-ops

### 6.2 用 AI 智能体扩展社区，同时保持人性化

- 原文：<https://www.vibevibe.cn/Articles/06-business-trends/keeping-community-human-while-scaling-with-agents.html>
- 核心：Vercel 社区把帖子分流、分类与后续跟进交给 Guardian 与 c0 两层智能体，让人把精力放回需要专业判断与真诚交流的回应上。
- 要点：
  - Guardian 分析内容、检测重复并按成员余量分配，48 小时无答复即重新分配。
  - 系统每 10 分钟检查一次新动态，两次检查之间休眠，中断后可从断点恢复。
  - c0 检索知识库、文档、GitHub issues 与历史讨论，整理成上下文包供团队使用。
  - 最初 23 天服务 281 位用户，完成 4,716 次初始响应，复活约八分之一被遗忘的话题。
- 适用：阶段 14 上线后迭代 / user-feedback-iteration、writing-plans

### 6.3 软件工程的未来两年

- 原文：<https://www.vibevibe.cn/Articles/06-business-trends/the-next-two-years-of-software-engineering.html>
- 核心：文章用初级岗位、技能结构、角色分工、专才与通才、教育路径五个问题，推演 AI 编程代理给未来两年软件工程带来的两种可能走向与应对方法。
- 要点：
  - 涉及 6200 万劳动者的研究显示，采用生成式 AI 后六个季度内初级开发者就业率下降约 9% 到 10%。
  - 已有 84% 开发者定期使用 AI 辅助工具，能力重心转向提出清晰要求并判断输出。
  - 接近一半工程岗位期待候选人覆盖多个领域，具备深度专长又能跨域协作的人更占优势。
  - 到 2024 年，接近 45% 的公司计划取消至少一部分岗位的学士学位要求。
- 适用：阶段 1 立项 / product-validation、brainstorming
