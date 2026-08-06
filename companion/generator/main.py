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
# 说明：stage 1-5 对应 立项门路径；6-15 续接架构门/业务门/上线门。
# 与 workflows/全流程速查.md 的 14 阶段模型存在偏移（gen N = 14模型 N-1，N>=6），
# 这是为兼容现有 01-05 模块而采用的生成器自身编号，避免重编号破坏已有文件。
STAGE_NEXT = {
    1: ("阶段1 纯想法 → 立项", "完成立项说明，进入阶段2 立项"),
    2: ("阶段2 立项", "产出 docs/a-立项文档.md，进入阶段3 选型"),
    3: ("阶段3 选型", "产出 docs/d-选型.md（含否决理由），进入阶段4 架构设计"),
    4: ("阶段4 架构设计", "产出 docs/e-架构设计文档.md，进入阶段5 宪法"),
    5: ("阶段5 宪法", "产出 AGENT_CONSTITUTION.md，通过立项门校验"),
    6: ("阶段6 实施真源文档拆解", "产出 docs/f-实施真源文档-第N阶段.md，进入阶段7 前端骨架"),
    7: ("阶段7 前端骨架搭建", "前端能启动、首页能打开，进入阶段8 数据库设计"),
    8: ("阶段8 数据库设计", "产出 docs/h-数据库设计文档.md + db/schema.sql 且能建表查询，进入阶段9 后端骨架"),
    9: ("阶段9 后端骨架搭建", "后端四条线跑通，进入阶段10 后端联调"),
    10: ("阶段10 后端联调", "前后端联调通过，进入阶段11 架构验收"),
    11: ("阶段11 架构验收", "后端+前端验收通过，通过架构门校验"),
    12: ("阶段12 业务模块开发", "业务模块逐个交付验收，进入阶段13 安全审查"),
    13: ("阶段13 后端安全审查", "安全 5 关全过，通过业务门校验"),
    14: ("阶段14 上线前总验收", "5 维度审查无严重项，进入阶段15 上线"),
    15: ("阶段15 上线 + Skill 反哺", "已上线 + 复盘 + 反哺，通过上线门校验"),
}

# 阶段 -> 临近门禁模块 id（抵达该 stage 时触发对应 gate 模块注入与校验）
GATE_AFTER = {
    5: "gate-立项门",
    11: "gate-架构门",
    13: "gate-业务门",
    15: "gate-上线门",
}


def next_gate(stage):
    """返回 >= stage 的最近门禁模块 id；无则 None。"""
    for g in sorted(GATE_AFTER):
        if g >= stage:
            return GATE_AFTER[g]
    return None


def gate_stage_of(gate_id):
    """门禁模块 id -> 其对应的 stage 编号。"""
    for g, gid in GATE_AFTER.items():
        if gid == gate_id:
            return g
    return None


def gate_cn(gate_id):
    """gate-立项门 -> 立项门"""
    return gate_id.replace("gate-", "")


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
    # 门禁状态：抵达门禁 stage 时注入对应 gate 模块正文；否则提示临近门禁
    gate_id = GATE_AFTER.get(stage)
    if gate_id and gate_id in mods:
        gate_status = "**临近门禁：" + mods[gate_id].get("title", gate_id) + "**\n\n" + mods[gate_id].get("body", "")
    else:
        ng = next_gate(stage)
        if ng:
            gate_status = ("未到门禁。下一步临近「%s」（阶段 %d）时校验对应硬指标。"
                           % (gate_cn(ng), gate_stage_of(ng)))
        else:
            gate_status = "已是最后阶段（上线门）。项目收尾 / 健康迭代中。"
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
    # 可读状态：按所在门禁显示「进行中 / 待校验」
    g = next_gate(stage)
    if g is None:
        state = "已上线（收尾）"
    elif gate_stage_of(g) == stage:
        state = "%s（待校验）" % gate_cn(g)
    else:
        state = "%s（进行中）" % gate_cn(g)
    gate = g or "无"
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
    ap.add_argument("--stage", type=int, default=1, help="当前阶段 (1-15)")
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
