// Vibe_Coding companion — 生成器（Go）
// 读 modules/ + 当前阶段，渲染 AGENTS.md 与 PROJECT_STATUS.md 到目标项目。
// 标准库，无第三方依赖。逻辑与 main.py 一致，用于接入 DC Ops 后端。
//
// 用法：
//   go run ./generator --project /path/to/project --stage 1
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	modulesDir   = filepath.Join("..", "modules")
	templatesDir = filepath.Join("..", "templates")
	frontRe      = regexp.MustCompile(`(?s)<!--\s*(.*?)\s*-->`)
)

// STAGE_NEXT: 阶段 -> (标题, 下一步短描述)
// 说明：stage 1-5 对应 立项门路径；6-15 续接架构门/业务门/上线门。
// 与 workflows/全流程速查.md 的 14 阶段模型存在偏移（gen N = 14模型 N-1，N>=6），
// 这是为兼容现有 01-05 模块而采用的生成器自身编号。逻辑与 main.py 保持一致。
var stageNext = map[int][2]string{
	1:  {"阶段1 纯想法 → 立项", "完成立项说明，进入阶段2 立项"},
	2:  {"阶段2 立项", "产出 docs/a-立项文档.md，进入阶段3 选型"},
	3:  {"阶段3 选型", "产出 docs/d-选型.md（含否决理由），进入阶段4 架构设计"},
	4:  {"阶段4 架构设计", "产出 docs/e-架构设计文档.md，进入阶段5 宪法"},
	5:  {"阶段5 宪法", "产出 AGENT_CONSTITUTION.md，通过立项门校验"},
	6:  {"阶段6 实施真源文档拆解", "产出 docs/f-实施真源文档-第N阶段.md，进入阶段7 前端骨架"},
	7:  {"阶段7 前端骨架搭建", "前端能启动、首页能打开，进入阶段8 数据库设计"},
	8:  {"阶段8 数据库设计", "产出 docs/h-数据库设计文档.md + db/schema.sql 且能建表查询，进入阶段9 后端骨架"},
	9:  {"阶段9 后端骨架搭建", "后端四条线跑通，进入阶段10 后端联调"},
	10: {"阶段10 后端联调", "前后端联调通过，进入阶段11 架构验收"},
	11: {"阶段11 架构验收", "后端+前端验收通过，通过架构门校验"},
	12: {"阶段12 业务模块开发", "业务模块逐个交付验收，进入阶段13 安全审查"},
	13: {"阶段13 后端安全审查", "安全 5 关全过，通过业务门校验"},
	14: {"阶段14 上线前总验收", "5 维度审查无严重项，进入阶段15 上线"},
	15: {"阶段15 上线 + Skill 反哺", "已上线 + 复盘 + 反哺，通过上线门校验"},
}

// gateAfter: 阶段 -> 临近门禁模块 id（抵达该 stage 时触发对应 gate 模块注入与校验）
var gateAfter = map[int]string{
	5:  "gate-立项门",
	11: "gate-架构门",
	13: "gate-业务门",
	15: "gate-上线门",
}

func nextGate(stage int) string {
	gates := make([]int, 0, len(gateAfter))
	for g := range gateAfter {
		gates = append(gates, g)
	}
	sort.Ints(gates)
	for _, g := range gates {
		if g >= stage {
			return gateAfter[g]
		}
	}
	return ""
}

func gateStageOf(gateID string) int {
	for g, gid := range gateAfter {
		if gid == gateID {
			return g
		}
	}
	return 0
}

func gateCN(gateID string) string {
	return strings.TrimPrefix(gateID, "gate-")
}

type module struct {
	id    string
	stage string
	typ   string
	body  string
}

func loadModules() (map[string]*module, error) {
	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil, err
	}
	mods := map[string]*module{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(modulesDir, e.Name()))
		if err != nil {
			return nil, err
		}
		text := string(b)
		m := &module{id: e.Name()}
		if sm := frontRe.FindStringSubmatch(text); sm != nil {
			for _, line := range strings.Split(sm[1], "\n") {
				kv := strings.SplitN(line, ":", 2)
				if len(kv) != 2 {
					continue
				}
				k, v := strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1])
				switch k {
				case "id":
					m.id = v
				case "stage":
					m.stage = v
				case "type":
					m.typ = v
				}
			}
			m.body = strings.TrimSpace(text[len(sm[0]):])
		} else {
			m.body = strings.TrimSpace(text)
		}
		mods[m.id] = m
	}
	return mods, nil
}

func renderAgents(stage int, mods map[string]*module) string {
	title, nextStep := stageNext[stage]
	phaseBody := ""
	for _, m := range mods {
		if m.stage == fmt.Sprint(stage) && m.typ == "phase" {
			phaseBody = m.body
		}
	}
	sj := ""
	if m, ok := mods["00-state-judgment"]; ok {
		sj = m.body
	}
	costParts := make([]string, 0, 2)
	for id, m := range mods {
		if m.typ == "cross" && id != "00-state-judgment" {
			costParts = append(costParts, m.body)
		}
	}
	costBody := strings.Join(costParts, "\n\n")
	// 门禁状态：抵达门禁 stage 时注入对应 gate 模块正文；否则提示临近门禁
	gateStatus := ""
	if gateID, ok := gateAfter[stage]; ok {
		if m, ok2 := mods[gateID]; ok2 {
			gateStatus = "**临近门禁：" + m.id + "**\n\n" + m.body
		}
	} else {
		ng := nextGate(stage)
		if ng != "" {
			gateStatus = fmt.Sprintf("未到门禁。下一步临近「%s」（阶段 %d）时校验对应硬指标。", gateCN(ng), gateStageOf(ng))
		} else {
			gateStatus = "已是最后阶段（上线门）。项目收尾 / 健康迭代中。"
		}
	}
	tpl, _ := os.ReadFile(filepath.Join(templatesDir, "AGENTS.md.tmpl"))
	s := string(tpl)
	s = strings.ReplaceAll(s, "__STAGE_TITLE__", title)
	s = strings.ReplaceAll(s, "__STAGE_BODY__", phaseBody)
	s = strings.ReplaceAll(s, "__NEXT_STEP__", nextStep)
	s = strings.ReplaceAll(s, "__STATE_JUDGMENT__", sj)
	s = strings.ReplaceAll(s, "__COST_BODY__", costBody)
	s = strings.ReplaceAll(s, "__GATE_STATUS__", gateStatus)
	s = strings.ReplaceAll(s, "__UPDATED_AT__", time.Now().Format("2006-01-02"))
	return s
}

func renderStatus(stage int, mods map[string]*module) string {
	title, nextStep := stageNext[stage]
	// 可读状态：按所在门禁显示「进行中 / 待校验」
	var state string
	g := nextGate(stage)
	if g == "" {
		state = "已上线（收尾）"
	} else if gateStageOf(g) == stage {
		state = gateCN(g) + "（待校验）"
	} else {
		state = gateCN(g) + "（进行中）"
	}
	gate := g
	if gate == "" {
		gate = "无"
	}
	tpl, _ := os.ReadFile(filepath.Join(templatesDir, "PROJECT_STATUS.md.tmpl"))
	s := string(tpl)
	s = strings.ReplaceAll(s, "__STATE__", state)
	s = strings.ReplaceAll(s, "__STAGE_NUM__", fmt.Sprint(stage))
	s = strings.ReplaceAll(s, "__STAGE_TITLE__", title)
	s = strings.ReplaceAll(s, "__NEXT_STEP_SHORT__", nextStep)
	s = strings.ReplaceAll(s, "__GATE__", gate)
	s = strings.ReplaceAll(s, "__UPDATED_AT__", time.Now().Format("2006-01-02"))
	return s
}

func main() {
	project := flag.String("project", "", "目标项目目录（必填）")
	stage := flag.Int("stage", 1, "当前阶段 (1-15)")
	flag.Parse()

	if *project == "" {
		fmt.Fprintln(os.Stderr, "错误：--project 必填")
		os.Exit(2)
	}
	if _, ok := stageNext[*stage]; !ok {
		keys := make([]int, 0, len(stageNext))
		for k := range stageNext {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		fmt.Fprintf(os.Stderr, "错误：stage 必须是 %v 之一\n", keys)
		os.Exit(2)
	}

	mods, err := loadModules()
	if err != nil {
		fmt.Fprintln(os.Stderr, "加载模块失败:", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(*project, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "创建目录失败:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(*project, "AGENTS.md"), []byte(renderAgents(*stage, mods)), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "写 AGENTS.md 失败:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(*project, "PROJECT_STATUS.md"), []byte(renderStatus(*stage, mods)), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "写 PROJECT_STATUS.md 失败:", err)
		os.Exit(1)
	}
	fmt.Printf("[OK] 已生成 %s/AGENTS.md 与 PROJECT_STATUS.md（阶段 %d）\n", *project, *stage)
}
