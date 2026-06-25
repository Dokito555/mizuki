import { c as sanitize_props, d as spread_props, a as slot, l as attributes, m as clsx, p as attr_style, j as attr, k as stringify, e as escape_html, i as ensure_array_like, n as derived, o as store_get, u as unsubscribe_stores } from "../../../../chunks/index.js";
import { B as Brain, p as page } from "../../../../chunks/stores.js";
import { b as useUpload, c as useAnalyzeUpload, d as useAIAnalyzeUpload } from "../../../../chunks/useUploads.js";
import { u as useFlows } from "../../../../chunks/useFlows.js";
import { B as Button } from "../../../../chunks/Toast.svelte_svelte_type_style_lang.js";
import "clsx";
import { C as Card } from "../../../../chunks/Card.js";
import { c as cn } from "../../../../chunks/utils3.js";
import "@sveltejs/kit/internal";
import "../../../../chunks/exports.js";
import "../../../../chunks/utils2.js";
import "@sveltejs/kit/internal/server";
import "../../../../chunks/root.js";
import "../../../../chunks/state.svelte.js";
import { I as Icon } from "../../../../chunks/Icon.js";
import { T as Triangle_alert } from "../../../../chunks/triangle-alert.js";
function Circle_check_big($$renderer, $$props) {
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
    ["path", { "d": "M21.801 10A10 10 0 1 1 17 3.335" }],
    ["path", { "d": "m9 11 3 3L22 4" }]
  ];
  Icon($$renderer, spread_props([
    { name: "circle-check-big" },
    $$sanitized_props,
    {
      /**
       * @component @name CircleCheckBig
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMjEuODAxIDEwQTEwIDEwIDAgMSAxIDE3IDMuMzM1IiAvPgogIDxwYXRoIGQ9Im05IDExIDMgM0wyMiA0IiAvPgo8L3N2Zz4K) - https://lucide.dev/icons/circle-check-big
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
function Clock($$renderer, $$props) {
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
    ["circle", { "cx": "12", "cy": "12", "r": "10" }],
    ["polyline", { "points": "12 6 12 12 16 14" }]
  ];
  Icon($$renderer, spread_props([
    { name: "clock" },
    $$sanitized_props,
    {
      /**
       * @component @name Clock
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8Y2lyY2xlIGN4PSIxMiIgY3k9IjEyIiByPSIxMCIgLz4KICA8cG9seWxpbmUgcG9pbnRzPSIxMiA2IDEyIDEyIDE2IDE0IiAvPgo8L3N2Zz4K) - https://lucide.dev/icons/clock
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
function Play($$renderer, $$props) {
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
  const iconNode = [["polygon", { "points": "6 3 20 12 6 21 6 3" }]];
  Icon($$renderer, spread_props([
    { name: "play" },
    $$sanitized_props,
    {
      /**
       * @component @name Play
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cG9seWdvbiBwb2ludHM9IjYgMyAyMCAxMiA2IDIxIDYgMyIgLz4KPC9zdmc+Cg==) - https://lucide.dev/icons/play
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
function Progress($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let {
      value = 0,
      max = 100,
      class: className = "",
      $$slots,
      $$events,
      ...rest
    } = $$props;
    $$renderer2.push(`<div${attributes({
      class: clsx(cn("relative h-2 w-full overflow-hidden rounded-full bg-secondary", className)),
      ...rest
    })}><div class="h-full bg-primary-600 transition-all duration-300 ease-out"${attr_style(`width: ${stringify(value / max * 100)}%`)} role="progressbar"${attr("aria-valuenow", value)}${attr("aria-valuemin", 0)}${attr("aria-valuemax", max)}></div></div>`);
  });
}
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    var $$store_subs;
    const uploadId = derived(() => Number(store_get($$store_subs ??= {}, "$page", page).params.id));
    const uploadQuery = useUpload(() => uploadId());
    const flowsQuery = useFlows(() => ({ upload_id: uploadId(), page_size: 50 }));
    const analyzeMutation = useAnalyzeUpload();
    const aiAnalyzeMutation = useAIAnalyzeUpload();
    const meta = derived(() => uploadQuery.data);
    const progress = derived(() => meta() ? meta().progress_pct : 0);
    const status = derived(() => meta() ? meta().status : "loading");
    const StatIcon = derived(() => status() === "done" ? Circle_check_big : status() === "error" ? Triangle_alert : Clock);
    const statColor = derived(() => status() === "done" ? "text-green-500" : status() === "error" ? "text-red-500" : "text-blue-500");
    $$renderer2.push(`<div class="space-y-6">`);
    if (uploadQuery.isPending) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-muted-foreground">Loading...</p>`);
    } else if (uploadQuery.data) {
      $$renderer2.push("<!--[1-->");
      $$renderer2.push(`<div class="flex items-center justify-between"><div><h1 class="text-2xl font-bold tracking-tight">${escape_html(meta().filename)}</h1> <p class="text-sm text-muted-foreground mt-1">Uploaded ${escape_html(new Date(meta().created_at).toLocaleString())}</p></div> <div class="flex items-center gap-2">`);
      Button($$renderer2, {
        onclick: () => analyzeMutation.mutate(uploadId()),
        variant: "outline",
        size: "sm",
        loading: analyzeMutation.isPending,
        children: ($$renderer3) => {
          Play($$renderer3, { class: "h-4 w-4 mr-1" });
          $$renderer3.push(`<!----> Analyze`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Button($$renderer2, {
        onclick: () => aiAnalyzeMutation.mutate(uploadId()),
        size: "sm",
        loading: aiAnalyzeMutation.isPending,
        children: ($$renderer3) => {
          Brain($$renderer3, { class: "h-4 w-4 mr-1" });
          $$renderer3.push(`<!----> AI Analyze`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div></div> <div class="grid gap-4 md:grid-cols-4">`);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-sm text-muted-foreground">Status</p> <div class="flex items-center gap-2 mt-1">`);
          if (StatIcon()) {
            $$renderer3.push("<!--[-->");
            StatIcon()($$renderer3, { class: "h-4 w-4 " + statColor() });
            $$renderer3.push("<!--]-->");
          } else {
            $$renderer3.push("<!--[!-->");
            $$renderer3.push("<!--]-->");
          }
          $$renderer3.push(` <span class="font-medium capitalize">${escape_html(status())}</span></div>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-sm text-muted-foreground">Packets</p> <p class="text-xl font-bold mt-1">${escape_html(meta().packets_processed.toLocaleString())}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-sm text-muted-foreground">Flows</p> <p class="text-xl font-bold mt-1">${escape_html(meta().flow_count)}</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> `);
      Card($$renderer2, {
        class: "p-4",
        children: ($$renderer3) => {
          $$renderer3.push(`<p class="text-sm text-muted-foreground">File Size</p> <p class="text-xl font-bold mt-1">${escape_html((meta().file_size / 1024 / 1024).toFixed(1))} MB</p>`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div> `);
      if (status() === "queued" || status() === "parsing" || status() === "inserting") {
        $$renderer2.push("<!--[0-->");
        Card($$renderer2, {
          class: "p-4",
          children: ($$renderer3) => {
            $$renderer3.push(`<p class="text-sm font-medium mb-2">Processing...</p> `);
            Progress($$renderer3, { value: progress(), max: 100 });
            $$renderer3.push(`<!----> <p class="text-xs text-muted-foreground mt-1">${escape_html(meta().packets_processed.toLocaleString())} packets processed</p>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]--> `);
      if (meta().error) {
        $$renderer2.push("<!--[0-->");
        Card($$renderer2, {
          class: "p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30",
          children: ($$renderer3) => {
            $$renderer3.push(`<p class="text-sm font-medium text-red-600 dark:text-red-400">Error: ${escape_html(meta().error)}</p>`);
          },
          $$slots: { default: true }
        });
      } else {
        $$renderer2.push("<!--[-1-->");
      }
      $$renderer2.push(`<!--]--> <div class="space-y-3"><h2 class="text-lg font-semibold">Flows</h2> `);
      if (flowsQuery.isPending) {
        $$renderer2.push("<!--[0-->");
        $$renderer2.push(`<p class="text-sm text-muted-foreground">Loading flows...</p>`);
      } else if (flowsQuery.data?.data?.length) {
        $$renderer2.push("<!--[1-->");
        $$renderer2.push(`<div class="overflow-x-auto"><table class="w-full text-sm"><thead><tr class="border-b border-gray-200 dark:border-gray-800"><th class="text-left py-2 px-3 font-medium">Src IP:Port</th><th class="text-left py-2 px-3 font-medium">Dst IP:Port</th><th class="text-left py-2 px-3 font-medium">Proto</th><th class="text-right py-2 px-3 font-medium">Score</th><th class="text-left py-2 px-3 font-medium">Threats</th><th class="text-right py-2 px-3 font-medium">Packets</th></tr></thead><tbody><!--[-->`);
        const each_array = ensure_array_like(flowsQuery.data.data);
        for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
          let flow = each_array[$$index];
          $$renderer2.push(`<tr class="border-b border-gray-100 dark:border-gray-900 hover:bg-gray-50 dark:hover:bg-gray-900/50 cursor-pointer"><td class="py-2 px-3 font-mono text-xs">${escape_html(flow.src_ip)}:${escape_html(flow.src_port)}</td><td class="py-2 px-3 font-mono text-xs">${escape_html(flow.dst_ip)}:${escape_html(flow.dst_port)}</td><td class="py-2 px-3">${escape_html(flow.protocol)}</td><td class="py-2 px-3 text-right font-mono">${escape_html(flow.score.toFixed(0))}</td><td class="py-2 px-3 max-w-[200px] truncate">${escape_html(flow.threats?.join(", ") ?? "-")}</td><td class="py-2 px-3 text-right">${escape_html(flow.packet_count)}</td></tr>`);
        }
        $$renderer2.push(`<!--]--></tbody></table></div>`);
      } else {
        $$renderer2.push("<!--[-1-->");
        $$renderer2.push(`<p class="text-sm text-muted-foreground">No flows extracted yet.</p>`);
      }
      $$renderer2.push(`<!--]--></div>`);
    } else {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<p class="text-muted-foreground">Upload not found.</p>`);
    }
    $$renderer2.push(`<!--]--></div>`);
    if ($$store_subs) unsubscribe_stores($$store_subs);
  });
}
export {
  _page as default
};
