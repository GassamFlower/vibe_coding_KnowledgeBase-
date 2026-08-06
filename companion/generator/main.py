#!/usr/bin/env python3
# Vibe_Coding companion — 生成器原型（Python）
# 读 modules/ + 当前阶段，渲染 AGENTS.md 与 PROJECT_STATUS.md 到目标项目。
# 标准库，无依赖。Go 版(main.go)逻辑与此一致，用于接入 DC Ops 后端。

import os
import re
import argparse
import datetime

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.dirname(HERE)
MODULES_DIR = os.path.join(ROOT, "modules")
TEMPLATES_DIR = os.path.join(ROOT, "templates")

FRONT_RE = re.compile(r"<!--\s*(.*?)\s*-->", re.S)

# stage -> (阶段标题, 下一步短描述)
STAGE_NEXT = {
    1: ("阶段1 纯想法 → 立项", "完成立项说明，进入阶段2 立项"),
    2: ("阶段2 立项", "产出 docs/立项.md，进入阶段3 选型"),
    3: ("阶段3 选型", "产出 docs/选型.md（含否决理由），进入阶段4 架构设计"),
    4: ("阶段4 架构设计", "产出 docs/架构.md，进入阶段5 宪法"),
    5: ("阶段5 宪法", "产出 CONSTITUTION.md，通过立项门校验"),
}

# 阶段 -> 临近门禁模块 id
GATE_AFTER = {5: "gate-立项门"}


def parse_module(path):
    text = open(path, encoding="utf-8").read()
    meta = {"body": text}
    m = FRONT_RE.search(text)
    if m:
        for line in m.group(1).splitlines():
            if ":" in line:
                k, v = line.split(":", 1)
                meta[k.strip()] = v.strip()
        meta["body"] = text[m.end():].strip()
    return meta


def load_modules():
    mods = {}
    for f in sorted(os.listdir(MODULES_DIR)):
        if f.endswith(".md"):
            meta = parse_module(os.path.join(MODULES_DIR, f))
            mods[meta.get("id", f)] = meta
    return mods


def render_agents(stage, mods):
    stage_title, next_step = STAGE_NEXT[stage]
    phase_body = ""
    for m in mods.values():
        if m.get("stage") == str(stage) and m.get("type") == "phase":
            phase_body = m.get("body", "")
    sj = mods.get("00-state-judgment", {}).get("body", "")
    cost_body = "\n\n".join(
        m.get("body", "") for m in mods.values()
        if m.get("type") == "cross" and m.get("id") != "00-state-judgment")
    gate_status = "未到门禁"
    if stage in GATE_AFTER:
        gate_status = "临近立项门，校验 7 项硬指标（见门禁模块）"
    tpl = open(os.path.join(TEMPLATES_DIR, "AGENTS.md.tmpl"), encoding="utf-8").read()
    return (tpl
            .replace("__STAGE_TITLE__", stage_title)
            .replace("__STAGE_BODY__", phase_body)
            .replace("__NEXT_STEP__", next_step)
            .replace("__STATE_JUDGMENT__", sj)
            .replace("__COST_BODY__", cost_body)
            .replace("__GATE_STATUS__", gate_status)
            .replace("__UPDATED_AT__", datetime.date.today().isoformat()))


def render_status(stage, mods):
    stage_title, next_step = STAGE_NEXT[stage]
    state = "1" if stage < 5 else "1(待立项门校验)"
    gate = GATE_AFTER.get(stage, "无")
    tpl = open(os.path.join(TEMPLATES_DIR, "PROJECT_STATUS.md.tmpl"), encoding="utf-8").read()
    return (tpl
            .replace("__STATE__", state)
            .replace("__STAGE_NUM__", str(stage))
            .replace("__STAGE_TITLE__", stage_title)
            .replace("__NEXT_STEP_SHORT__", next_step)
            .replace("__GATE__", gate)
            .replace("__UPDATED_AT__", datetime.date.today().isoformat()))


def main():
    ap = argparse.ArgumentParser(description="Vibe_Coding companion 生成器（Python 原型）")
    ap.add_argument("--project", required=True, help="目标项目目录")
    ap.add_argument("--stage", type=int, default=1, help="当前阶段 (1-5)")
    args = ap.parse_args()

    if args.stage not in STAGE_NEXT:
        raise SystemExit(f"stage 必须是 {sorted(STAGE_NEXT)} 之一")

    mods = load_modules()
    agents = render_agents(args.stage, mods)
    status = render_status(args.stage, mods)

    os.makedirs(args.project, exist_ok=True)
    with open(os.path.join(args.project, "AGENTS.md"), "w", encoding="utf-8") as f:
        f.write(agents)
    with open(os.path.join(args.project, "PROJECT_STATUS.md"), "w", encoding="utf-8") as f:
        f.write(status)

    print(f"[OK] 已生成 {args.project}/AGENTS.md 与 PROJECT_STATUS.md（阶段 {args.stage}）")


if __name__ == "__main__":
    main()
