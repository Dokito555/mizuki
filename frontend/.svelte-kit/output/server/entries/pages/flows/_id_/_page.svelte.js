import { c as sanitize_props, d as spread_props, a as slot, e as escape_html, h as attr_class, k as stringify, i as ensure_array_like, n as derived, o as store_get, u as unsubscribe_stores } from "../../../../chunks/index.js";
import { B as Brain, p as page } from "../../../../chunks/stores.js";
import { a as useFlow } from "../../../../chunks/useFlows.js";
import { a as api, b as createQuery } from "../../../../chunks/utils3.js";
import { c as createMutation, C as Card } from "../../../../chunks/Card.js";
import { B as Button } from "../../../../chunks/Toast.svelte_svelte_type_style_lang.js";
import "clsx";
import { I as Icon } from "../../../../chunks/Icon.js";
import { T as Triangle_alert } from "../../../../chunks/triangle-alert.js";
function Activity($$renderer, $$props) {
  const $$sanitized_props = sanitize_props($$props);
  /**
   * @license lucide-svelte v0.510.0 - ISC
   *
   * ISC License
   *
   * Copyright (c) for portions of Lucide are held by Cole Bemis 2013-2022 as part of Feather (MIT). All other copyright (c) for Lucide are held by Lucide Contributors 2022.
   *
   * Permission to use, copy, modify, and/or distribute this software for any
   * purpose with or without fee is hereby granted, provided that the above
   * copyright notice and this permission notice appear in all copies.
   *
   * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
   * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
   * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
   * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
   * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
   * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
   * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
   *
   */
  const iconNode = [
    [
      "path",
      {
        "d": "M22 12h-2.48a2 2 0 0 0-1.93 1.46l-2.35 8.36a.25.25 0 0 1-.48 0L9.24 2.18a.25.25 0 0 0-.48 0l-2.35 8.36A2 2 0 0 1 4.49 12H2"
      }
    ]
  ];
  Icon($$renderer, spread_props([
    { name: "activity" },
    $$sanitized_props,
    {
      /**
       * @component @name Activity
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMjIgMTJoLTIuNDhhMiAyIDAgMCAwLTEuOTMgMS40NmwtMi4zNSA4LjM2YS4yNS4yNSAwIDAgMS0uNDggMEw5LjI0IDIuMThhLjI1LjI1IDAgMCAwLS40OCAwbC0yLjM1IDguMzZBMiAyIDAgMCAxIDQuNDkgMTJIMiIgLz4KPC9zdmc+Cg==) - https://lucide.dev/icons/activity
       * @see https://lucide.dev/guide/packages/lucide-svelte - Documentation
       *
       * @param {Object} props - Lucide icons props and any valid SVG attribute
       * @returns {FunctionalComponent} Svelte component
       *
       */
      iconNode,
      children: ($$renderer2) => {
        $$renderer2.push(`<!--[-->`);
        slot($$renderer2, $$props, "default", {});
        $$renderer2.push(`<!--]-->`);
      },
      $$slots: { default: true }
    }
  ]));
}
function Tag($$renderer, $$props) {
  const $$sanitized_props = sanitize_props($$props);
  /**
   * @license lucide-svelte v0.510.0 - ISC
   *
   * ISC License
   *
   * Copyright (c) for portions of Lucide are held by Cole Bemis 2013-2022 as part of Feather (MIT). All other copyright (c) for Lucide are held by Lucide Contributors 2022.
   *
   * Permission to use, copy, modify, and/or distribute this software for any
   * purpose with or without fee is hereby granted, provided that the above
   * copyright notice and this permission notice appear in all copies.
   *
   * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
   * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
   * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
   * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
   * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
   * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
   * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
   *
   */
  const iconNode = [
    [
      "path",
      {
        "d": "M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"
      }
    ],
    [
      "circle",
      { "cx": "7.5", "cy": "7.5", "r": ".5", "fill": "currentColor" }
    ]
  ];
  Icon($$renderer, spread_props([
    { name: "tag" },
    $$sanitized_props,
    {
      /**
       * @component @name Tag
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMTIuNTg2IDIuNTg2QTIgMiAwIDAgMCAxMS4xNzIgMkg0YTIgMiAwIDAgMC0yIDJ2Ny4xNzJhMiAyIDAgMCAwIC41ODYgMS40MTRsOC43MDQgOC43MDRhMi40MjYgMi40MjYgMCAwIDAgMy40MiAwbDYuNTgtNi41OGEyLjQyNiAyLjQyNiAwIDAgMCAwLTMuNDJ6IiAvPgogIDxjaXJjbGUgY3g9IjcuNSIgY3k9IjcuNSIgcj0iLjUiIGZpbGw9ImN1cnJlbnRDb2xvciIgLz4KPC9zdmc+Cg==) - https://lucide.dev/icons/tag
       * @see https://lucide.dev/guide/packages/lucide-svelte - Documentation
       *
       * @param {Object} props - Lucide icons props and any valid SVG attribute
       * @returns {FunctionalComponent} Svelte component
       *
       */
      iconNode,
      children: ($$renderer2) => {
        $$renderer2.push(`<!--[-->`);
        slot($$renderer2, $$props, "default", {});
        $$renderer2.push(`<!--]-->`);
      },
      $$slots: { default: true }
    }
  ]));
}
function Wifi($$renderer, $$props) {
  const $$sanitized_props = sanitize_props($$props);
  /**
   * @license lucide-svelte v0.510.0 - ISC
   *
   * ISC License
   *
   * Copyright (c) for portions of Lucide are held by Cole Bemis 2013-2022 as part of Feather (MIT). All other copyright (c) for Lucide are held by Lucide Contributors 2022.
   *
   * Permission to use, copy, modify, and/or distribute this software for any
   * purpose with or without fee is hereby granted, provided that the above
   * copyright notice and this permission notice appear in all copies.
   *
   * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
   * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
   * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
   * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
   * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
   * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
   * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
   *
   */
  const iconNode = [
    ["path", { "d": "M12 20h.01" }],
    ["path", { "d": "M2 8.82a15 15 0 0 1 20 0" }],
    ["path", { "d": "M5 12.859a10 10 0 0 1 14 0" }],
    ["path", { "d": "M8.5 16.429a5 5 0 0 1 7 0" }]
  ];
  Icon($$renderer, spread_props([
    { name: "wifi" },
    $$sanitized_props,
    {
      /**
       * @component @name Wifi
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMTIgMjBoLjAxIiAvPgogIDxwYXRoIGQ9Ik0yIDguODJhMTUgMTUgMCAwIDEgMjAgMCIgLz4KICA8cGF0aCBkPSJNNSAxMi44NTlhMTAgMTAgMCAwIDEgMTQgMCIgLz4KICA8cGF0aCBkPSJNOC41IDE2LjQyOWE1IDUgMCAwIDEgNyAwIiAvPgo8L3N2Zz4K) - https://lucide.dev/icons/wifi
       * @see https://lucide.dev/guide/packages/lucide-svelte - Documentation
       *
       * @param {Object} props - Lucide icons props and any valid SVG attribute
       * @returns {FunctionalComponent} Svelte component
       *
       */
      iconNode,
      children: ($$renderer2) => {
        $$renderer2.push(`<!--[-->`);
        slot($$renderer2, $$props, "default", {});
        $$renderer2.push(`<!--]-->`);
      },
      $$slots: { default: true }
    }
  ]));
}
async function getAIAnalysis(flowId) {
  const res = await api.get(`/flows/${flowId}/ai`);
  return res.data;
}
async function analyzeFlow(flowId) {
  const res = await api.post(`/flows/${flowId}/ai-analyze`);
  return res.data;
}
function useAIAnalysis(flowId) {
  return createQuery(() => ({
    queryKey: ["ai-analysis", flowId()],
    queryFn: () => getAIAnalysis(flowId())
  }));
}
function useAnalyzeFlow() {
  return createMutation(() => ({
    mutationFn: (flowId) => analyzeFlow(flowId)
  }));
}
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    var $$store_subs;
    const flowId = derived(() => Number(store_get($$store_subs ??= {}, "$page", page).params.id));
    const flowQuery = useFlow(() => flowId());
    const aiQuery = useAIAnalysis(() => flowId());
    const analyzeFlowMutation = useAnalyzeFlow();
    function scoreColor(score) {
      if (score >= 70) return "text-red-500 font-bold";
      if (score >= 40) return "text-yellow-500 font-medium";
      if (score > 0) return "text-blue-500";
      return "text-muted-foreground";
    }
    async function runAI() {
      await analyzeFlowMutation.mutateAsync(flowId());
      aiQuery.refetch();
    }
    $$renderer2.push(`<div class="space-y-6">`);
    if (flowQuery.isPending) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground">Loading...</p>`);
    } else if (flowQuery.data) {
      $$renderer2.push("<!--[1-->");
      const flow = flowQuery.data;
      $$renderer2.push(`<div class="flex items-center justify-between"><h1 class="text-2xl font-bold tracking-tight">Flow #${escape_html(flow.id)}</h1> <div class="flex items-center gap-2">`);
      Button($$renderer2, {
        onclick: runAI,
        loading: analyzeFlowMutation.isPending,
        children: ($$renderer3) => {
          Brain($$renderer3, { class: "h-4 w-4 mr-1" });
          $$renderer3.push(`<!----> Analyze with AI`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div></div> <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">`);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">`);
          Wifi($$renderer3, { class: "h-4 w-4" });
          $$renderer3.push(`<!----> Source</div> <p class="font-mono text-sm">${escape_html(flow.src_ip)}:${escape_html(flow.src_port)}</p> <p class="text-xs text-muted-foreground mt-1">${escape_html(flow.src_mac ?? "-")}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">`);
          Wifi($$renderer3, { class: "h-4 w-4" });
          $$renderer3.push(`<!----> Destination</div> <p class="font-mono text-sm">${escape_html(flow.dst_ip)}:${escape_html(flow.dst_port)}</p> <p class="text-xs text-muted-foreground mt-1">${escape_html(flow.dst_mac ?? "-")}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">`);
          Activity($$renderer3, { class: "h-4 w-4" });
          $$renderer3.push(`<!----> Protocol</div> <p class="font-mono text-sm">${escape_html(flow.protocol)}</p> <p class="text-xs text-muted-foreground mt-1">${escape_html(flow.app_protocol ?? "-")}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">`);
          Triangle_alert($$renderer3, { class: "h-4 w-4" });
          $$renderer3.push(`<!----> Score</div> <p${attr_class(`font-mono text-2xl font-bold ${stringify(scoreColor(flow.score))}`)}>${escape_html(flow.score.toFixed(0))}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div> <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">`);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-xs text-muted-foreground">Packets</p> <p class="font-mono text-lg font-medium mt-1">${escape_html(flow.packet_count.toLocaleString())}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-xs text-muted-foreground">Bytes</p> <p class="font-mono text-lg font-medium mt-1">${escape_html((flow.byte_count / 1024 / 1024).toFixed(2))} MB</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-xs text-muted-foreground">IAT Avg</p> <p class="font-mono text-lg font-medium mt-1">${escape_html(flow.iat_avg_ms.toFixed(1))} ms</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-xs text-muted-foreground">IAT StdDev</p> <p class="font-mono text-lg font-medium mt-1">${escape_html(flow.iat_std_dev_ms.toFixed(1))} ms</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div> `);
      if (flow.threats?.length) {
        $$renderer2.push("<!--[0-->");
        Card($$renderer2, {
          class: "p-4 border-yellow-300 dark:border-yellow-700 bg-yellow-50 dark:bg-yellow-950/30",
          children: ($$renderer3) => {
            $$renderer3.push(`<div class="flex items-center gap-2 mb-2">`);
            Triangle_alert($$renderer3, { class: "h-5 w-5 text-yellow-500" });
            $$renderer3.push(`<!----> <span class="font-semibold">Detected Threats</span></div> <ul class="space-y-1"><!--[-->`);
            const each_array = ensure_array_like(flow.threats);
            for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
              let threat = each_array[$$index];
              $$renderer3.push(`<li class="flex items-center gap-2 text-sm">`);
              Tag($$renderer3, { class: "h-3 w-3 text-yellow-500" });
              $$renderer3.push(`<!----> ${escape_html(threat)}</li>`);
            }
            $$renderer3.push(`<!--]--></ul>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]--> `);
      if (flow.tls_sni || flow.tls_version) {
        $$renderer2.push("<!--[0-->");
        Card($$renderer2, {
          class: "p-4",
          children: ($$renderer3) => {
            $$renderer3.push(`<h3 class="font-semibold text-sm mb-2">TLS / DNS</h3> <div class="grid grid-cols-2 gap-4 text-sm">`);
            if (flow.tls_version) {
              $$renderer3.push("<!--[0-->");
              $$renderer3.push(`<div><span class="text-muted-foreground">Version:</span> ${escape_html(flow.tls_version)}</div>`);
            } else {
              $$renderer3.push("<!--[-1-->");
            }
            $$renderer3.push(`<!--]--> `);
            if (flow.tls_sni) {
              $$renderer3.push("<!--[0-->");
              $$renderer3.push(`<div><span class="text-muted-foreground">SNI:</span> ${escape_html(flow.tls_sni)}</div>`);
            } else {
              $$renderer3.push("<!--[-1-->");
            }
            $$renderer3.push(`<!--]--> `);
            if (flow.dns_queries?.length) {
              $$renderer3.push("<!--[0-->");
              $$renderer3.push(`<div class="col-span-2"><span class="text-muted-foreground">DNS Queries:</span> <ul class="list-disc list-inside mt-1"><!--[-->`);
              const each_array_1 = ensure_array_like(flow.dns_queries);
              for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
                let q = each_array_1[$$index_1];
                $$renderer3.push(`<li class="font-mono text-xs">${escape_html(q)}</li>`);
              }
              $$renderer3.push(`<!--]--></ul></div>`);
            } else {
              $$renderer3.push("<!--[-1-->");
            }
            $$renderer3.push(`<!--]--></div>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]--> `);
      if (aiQuery.data) {
        $$renderer2.push("<!--[0-->");
        const ai = aiQuery.data;
        if (ai.status === "analyzed" && ai.analysis) {
          $$renderer2.push("<!--[0-->");
          Card($$renderer2, {
            class: "p-4 border-primary-300 dark:border-primary-700",
            children: ($$renderer3) => {
              const isFallback = ai.analysis.is_fallback;
              $$renderer3.push(`<div class="flex items-center justify-between mb-3"><div class="flex items-center gap-2">`);
              Brain($$renderer3, { class: "h-5 w-5 text-primary-600" });
              $$renderer3.push(`<!----> <h3 class="font-semibold">AI Analysis</h3></div> <span class="text-xs text-muted-foreground">Model: ${escape_html(ai.model)} | Confidence: ${escape_html((ai.analysis.confidence * 100).toFixed(0))}%</span></div> <p class="text-sm leading-relaxed">${escape_html(ai.analysis.narrative)}</p> `);
              if (isFallback) {
                $$renderer3.push("<!--[0-->");
                $$renderer3.push(`<p class="text-xs text-yellow-500 mt-2">Heuristic fallback (AI provider unavailable)</p>`);
              } else {
                $$renderer3.push("<!--[-1-->");
              }
              $$renderer3.push(`<!--]--> `);
              if (ai.analysis.mitre_ids?.length) {
                $$renderer3.push("<!--[0-->");
                $$renderer3.push(`<div class="mt-3"><p class="text-xs font-medium text-muted-foreground mb-1">MITRE ATT&amp;CK</p> <div class="flex flex-wrap gap-1"><!--[-->`);
                const each_array_2 = ensure_array_like(ai.analysis.mitre_ids);
                for (let $$index_2 = 0, $$length = each_array_2.length; $$index_2 < $$length; $$index_2++) {
                  let id = each_array_2[$$index_2];
                  $$renderer3.push(`<span class="text-xs px-2 py-0.5 rounded-full bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300 font-mono">${escape_html(id)}</span>`);
                }
                $$renderer3.push(`<!--]--></div></div>`);
              } else {
                $$renderer3.push("<!--[-1-->");
              }
              $$renderer3.push(`<!--]--> `);
              if (ai.analysis.attribution) {
                $$renderer3.push("<!--[0-->");
                $$renderer3.push(`<p class="text-xs text-muted-foreground mt-2"><span class="font-medium">Attribution:</span> ${escape_html(ai.analysis.attribution)}</p>`);
              } else {
                $$renderer3.push("<!--[-1-->");
              }
              $$renderer3.push(`<!--]--> `);
              if (ai.analysis.remediation?.length) {
                $$renderer3.push("<!--[0-->");
                $$renderer3.push(`<div class="mt-3"><p class="text-xs font-medium text-muted-foreground mb-1">Remediation</p> <ul class="space-y-1"><!--[-->`);
                const each_array_3 = ensure_array_like(ai.analysis.remediation);
                for (let $$index_3 = 0, $$length = each_array_3.length; $$index_3 < $$length; $$index_3++) {
                  let r = each_array_3[$$index_3];
                  $$renderer3.push(`<li class="text-sm flex items-start gap-2"><span class="text-green-500 mt-0.5">•</span> ${escape_html(r)}</li>`);
                }
                $$renderer3.push(`<!--]--></ul></div>`);
              } else {
                $$renderer3.push("<!--[-1-->");
              }
              $$renderer3.push(`<!--]--> `);
              if (ai.analysis.correlations?.length) {
                $$renderer3.push("<!--[0-->");
                $$renderer3.push(`<div class="mt-3"><p class="text-xs font-medium text-muted-foreground mb-1">Correlations</p> <!--[-->`);
                const each_array_4 = ensure_array_like(ai.analysis.correlations);
                for (let $$index_4 = 0, $$length = each_array_4.length; $$index_4 < $$length; $$index_4++) {
                  let corr = each_array_4[$$index_4];
                  $$renderer3.push(`<div class="text-sm p-2 rounded bg-gray-50 dark:bg-gray-900 mb-1"><p class="font-medium">${escape_html(corr.pattern)}</p> <p class="text-xs text-muted-foreground">${escape_html(corr.description)}</p></div>`);
                }
                $$renderer3.push(`<!--]--></div>`);
              } else {
                $$renderer3.push("<!--[-1-->");
              }
              $$renderer3.push(`<!--]-->`);
            },
            $$slots: { default: true }
          });
        } else {
          $$renderer2.push("<!--[-1-->");
          Card($$renderer2, {
            class: "p-4",
            children: ($$renderer3) => {
              $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground">`);
              Brain($$renderer3, { class: "h-5 w-5" });
              $$renderer3.push(`<!----> <p class="text-sm">No AI analysis yet. Click "Analyze with AI" to run.</p></div>`);
            },
            $$slots: { default: true }
          });
        }
        $$renderer2.push(`<!--]-->`);
      } else if (!aiQuery.isPending) {
        $$renderer2.push("<!--[1-->");
        Card($$renderer2, {
          class: "p-4",
          children: ($$renderer3) => {
            $$renderer3.push(`<div class="flex items-center gap-2 text-muted-foreground">`);
            Brain($$renderer3, { class: "h-5 w-5" });
            $$renderer3.push(`<!----> <p class="text-sm">No AI analysis yet. Click "Analyze with AI" to run.</p></div>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]--> `);
      if (flow.packet_samples?.length) {
        $$renderer2.push("<!--[0-->");
        Card($$renderer2, {
          class: "p-4",
          children: ($$renderer3) => {
            $$renderer3.push(`<h3 class="font-semibold text-sm mb-3">Packet Samples (${escape_html(flow.packet_samples.length)})</h3> <div class="overflow-x-auto"><table class="w-full text-xs"><thead><tr class="border-b border-gray-200 dark:border-gray-800"><th class="text-left py-1 px-2 font-medium">#</th><th class="text-left py-1 px-2 font-medium">Timestamp</th><th class="text-right py-1 px-2 font-medium">Size</th></tr></thead><tbody><!--[-->`);
            const each_array_5 = ensure_array_like(flow.packet_samples);
            for (let i = 0, $$length = each_array_5.length; i < $$length; i++) {
              let sample = each_array_5[i];
              $$renderer3.push(`<tr class="border-b border-gray-100 dark:border-gray-900"><td class="py-1 px-2">${escape_html(i + 1)}</td><td class="py-1 px-2 font-mono">${escape_html(new Date(sample.ts).toISOString())}</td><td class="py-1 px-2 text-right">${escape_html(sample.size)} B</td></tr>`);
            }
            $$renderer3.push(`<!--]--></tbody></table></div>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]-->`);
    } else {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<p class="text-muted-foreground">Flow not found.</p>`);
    }
    $$renderer2.push(`<!--]--></div>`);
    if ($$store_subs) unsubscribe_stores($$store_subs);
  });
}
export {
  _page as default
};
