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
var stageNext = map[int][2]string{
	1: {"阶段1 纯想法 → 立项", "完成立项说明，进入阶段2 立项"},
	2: {"阶段2 立项", "产出 docs/立项.md，进入阶段3 选型"},
	3: {"阶段3 选型", "产出 docs/选型.md（含否决理由），进入阶段4 架构设计"},
	4: {"阶段4 架构设计", "产出 docs/架构.md，进入阶段5 宪法"},
	5: {"阶段5 宪法", "产出 CONSTITUTION.md，通过立项门校验"},
}

// gateAfter: 阶段 -> 临近门禁模块 id
var gateAfter = map[int]string{5: "gate-立项门"}

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
	gateStatus := "未到门禁"
	if _, ok := gateAfter[stage]; ok {
		gateStatus = "临近立项门，校验 7 项硬指标（见门禁模块）"
	}
	tpl, _ := os.ReadFile(filepath.Join(templatesDir, "AGENTS.md.tmpl"))
	s := string(tpl)
	s = strings.ReplaceAll(s, "__STAGE_TITLE__", title)
	s = strings.ReplaceAll(s, "__STAGE_BODY__", phaseBody)
	s = strings.ReplaceAll(s, "__NEXT_STEP__", nextStep)
	s = strings.ReplaceAll(s, "__STATE_JUDGMENT__", sj)
	s = strings.ReplaceAll(s, "__GATE_STATUS__", gateStatus)
	s = strings.ReplaceAll(s, "__UPDATED_AT__", time.Now().Format("2006-01-02"))
	return s
}

func renderStatus(stage int, mods map[string]*module) string {
	title, nextStep := stageNext[stage]
	state := "1"
	if stage >= 5 {
		state = "1(待立项门校验)"
	}
	gate := "无"
	if g, ok := gateAfter[stage]; ok {
		gate = g
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
	stage := flag.Int("stage", 1, "当前阶段 (1-5)")
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
