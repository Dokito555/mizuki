import { c as sanitize_props, d as spread_props, a as slot, l as attributes, m as clsx, b as bind_props, e as escape_html, i as ensure_array_like, h as attr_class, k as stringify } from "../../../chunks/index.js";
import { u as useFlows } from "../../../chunks/useFlows.js";
import { B as Button } from "../../../chunks/Toast.svelte_svelte_type_style_lang.js";
import { c as cn } from "../../../chunks/utils3.js";
import "clsx";
import "@sveltejs/kit/internal";
import "../../../chunks/exports.js";
import "../../../chunks/utils2.js";
import "@sveltejs/kit/internal/server";
import "../../../chunks/root.js";
import "../../../chunks/state.svelte.js";
import { I as Icon } from "../../../chunks/Icon.js";
function Arrow_up_down($$renderer, $$props) {
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
    ["path", { "d": "m21 16-4 4-4-4" }],
    ["path", { "d": "M17 20V4" }],
    ["path", { "d": "m3 8 4-4 4 4" }],
    ["path", { "d": "M7 4v16" }]
  ];
  Icon($$renderer, spread_props([
    { name: "arrow-up-down" },
    $$sanitized_props,
    {
      /**
       * @component @name ArrowUpDown
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJtMjEgMTYtNCA0LTQtNCIgLz4KICA8cGF0aCBkPSJNMTcgMjBWNCIgLz4KICA8cGF0aCBkPSJtMyA4IDQtNCA0IDQiIC8+CiAgPHBhdGggZD0iTTcgNHYxNiIgLz4KPC9zdmc+Cg==) - https://lucide.dev/icons/arrow-up-down
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
function Input($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let {
      type = "text",
      placeholder = "",
      disabled = false,
      value = "",
      class: className = "",
      $$slots,
      $$events,
      ...rest
    } = $$props;
    $$renderer2.push(`<input${attributes(
      {
        type,
        class: clsx(cn("flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50", className)),
        placeholder,
        disabled,
        value,
        ...rest
      },
      void 0,
      void 0,
      void 0,
      4
    )}/>`);
    bind_props($$props, { value });
  });
}
function Select($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let {
      placeholder = "",
      disabled = false,
      value = void 0,
      options = [],
      class: className = "",
      $$slots,
      $$events,
      ...rest
    } = $$props;
    $$renderer2.push(`<div class="relative">`);
    $$renderer2.select(
      {
        class: cn("flex h-10 w-full appearance-none rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50", className),
        value,
        disabled,
        ...rest
      },
      ($$renderer3) => {
        if (placeholder) {
          $$renderer3.push("<!--[0-->");
          $$renderer3.option({ value: "", disabled: true, selected: true }, ($$renderer4) => {
            $$renderer4.push(`${escape_html(placeholder)}`);
          });
        } else {
          $$renderer3.push("<!--[-1-->");
        }
        $$renderer3.push(`<!--]--><!--[-->`);
        const each_array = ensure_array_like(options);
        for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
          let opt = each_array[$$index];
          $$renderer3.option({ value: opt.value }, ($$renderer4) => {
            $$renderer4.push(`${escape_html(opt.label)}`);
          });
        }
        $$renderer3.push(`<!--]-->`);
      }
    );
    $$renderer2.push(`</div>`);
    bind_props($$props, { value });
  });
}
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let filters = {
      src_ip: "",
      dst_ip: "",
      protocol: "",
      min_score: void 0,
      page: 1,
      page_size: 20,
      sort_by: "score",
      sort_desc: true
    };
    const flowsQuery = useFlows(() => filters);
    const protocolOptions = [
      { value: "", label: "All Protocols" },
      { value: "TCP", label: "TCP" },
      { value: "UDP", label: "UDP" },
      { value: "ICMP", label: "ICMP" }
    ];
    function applyFilter() {
      filters = { ...filters, page: 1 };
    }
    function resetFilters() {
      filters = {
        src_ip: "",
        dst_ip: "",
        protocol: "",
        min_score: void 0,
        page: 1,
        page_size: 20,
        sort_by: "score",
        sort_desc: true
      };
    }
    function scoreColor(score) {
      if (score >= 70) return "text-red-500 font-bold";
      if (score >= 40) return "text-yellow-500 font-medium";
      if (score > 0) return "text-blue-500";
      return "text-muted-foreground";
    }
    let $$settled = true;
    let $$inner_renderer;
    function $$render_inner($$renderer3) {
      $$renderer3.push(`<div class="space-y-6"><h1 class="text-2xl font-bold tracking-tight">Flows</h1> <div class="flex flex-wrap items-end gap-3 p-4 border rounded-lg bg-gray-50 dark:bg-gray-900/50"><div class="space-y-1"><label class="text-xs font-medium" for="src-ip">Source IP</label> `);
      Input($$renderer3, {
        id: "src-ip",
        placeholder: "e.g. 10.0.0.1",
        oninput: applyFilter,
        get value() {
          return filters.src_ip;
        },
        set value($$value) {
          filters.src_ip = $$value;
          $$settled = false;
        }
      });
      $$renderer3.push(`<!----></div> <div class="space-y-1"><label class="text-xs font-medium" for="dst-ip">Dest IP</label> `);
      Input($$renderer3, {
        id: "dst-ip",
        placeholder: "e.g. 192.168.1.1",
        oninput: applyFilter,
        get value() {
          return filters.dst_ip;
        },
        set value($$value) {
          filters.dst_ip = $$value;
          $$settled = false;
        }
      });
      $$renderer3.push(`<!----></div> <div class="space-y-1"><label class="text-xs font-medium" for="protocol">Protocol</label> `);
      Select($$renderer3, {
        options: protocolOptions,
        placeholder: "All Protocols",
        onchange: applyFilter,
        get value() {
          return filters.protocol;
        },
        set value($$value) {
          filters.protocol = $$value;
          $$settled = false;
        }
      });
      $$renderer3.push(`<!----></div> <div class="space-y-1"><label class="text-xs font-medium" for="min-score">Min Score</label> `);
      Input($$renderer3, {
        type: "number",
        id: "min-score",
        placeholder: "0",
        oninput: applyFilter,
        get value() {
          return filters.min_score;
        },
        set value($$value) {
          filters.min_score = $$value;
          $$settled = false;
        }
      });
      $$renderer3.push(`<!----></div> `);
      Button($$renderer3, {
        variant: "outline",
        size: "sm",
        onclick: resetFilters,
        children: ($$renderer4) => {
          $$renderer4.push(`<!---->Reset`);
        },
        $$slots: { default: true }
      });
      $$renderer3.push(`<!----></div> `);
      if (flowsQuery.isPending) {
        $$renderer3.push("<!--[0-->");
        $$renderer3.push(`<p class="text-sm text-muted-foreground">Loading...</p>`);
      } else if (flowsQuery.data?.data?.length) {
        $$renderer3.push("<!--[1-->");
        $$renderer3.push(`<div class="overflow-x-auto border rounded-lg"><table class="w-full text-sm"><thead class="bg-gray-50 dark:bg-gray-900"><tr><th class="text-left py-3 px-3 font-medium cursor-pointer hover:text-primary-600"><div class="flex items-center gap-1">Src IP:Port `);
        Arrow_up_down($$renderer3, { class: "h-3 w-3" });
        $$renderer3.push(`<!----></div></th><th class="text-left py-3 px-3 font-medium">Dst IP:Port</th><th class="text-left py-3 px-3 font-medium">Proto</th><th class="text-right py-3 px-3 font-medium cursor-pointer hover:text-primary-600"><div class="flex items-center justify-end gap-1">Score `);
        Arrow_up_down($$renderer3, { class: "h-3 w-3" });
        $$renderer3.push(`<!----></div></th><th class="text-left py-3 px-3 font-medium">Threats</th><th class="text-right py-3 px-3 font-medium">Packets</th><th class="text-right py-3 px-3 font-medium">Bytes</th></tr></thead><tbody><!--[-->`);
        const each_array = ensure_array_like(flowsQuery.data.data);
        for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
          let flow = each_array[$$index];
          $$renderer3.push(`<tr class="border-t border-gray-100 dark:border-gray-900 hover:bg-gray-50 dark:hover:bg-gray-900/50 cursor-pointer"><td class="py-2.5 px-3 font-mono text-xs">${escape_html(flow.src_ip)}:${escape_html(flow.src_port)}</td><td class="py-2.5 px-3 font-mono text-xs">${escape_html(flow.dst_ip)}:${escape_html(flow.dst_port)}</td><td class="py-2.5 px-3">${escape_html(flow.protocol)}</td><td${attr_class(`py-2.5 px-3 text-right font-mono ${stringify(scoreColor(flow.score))}`)}>${escape_html(flow.score.toFixed(0))}</td><td class="py-2.5 px-3 max-w-[200px] truncate text-xs">${escape_html(flow.threats?.join(", ") ?? "-")}</td><td class="py-2.5 px-3 text-right">${escape_html(flow.packet_count.toLocaleString())}</td><td class="py-2.5 px-3 text-right">${escape_html((flow.byte_count / 1024).toFixed(0))} KB</td></tr>`);
        }
        $$renderer3.push(`<!--]--></tbody></table></div> <div class="flex items-center justify-between mt-4"><span class="text-sm text-muted-foreground">Total: ${escape_html(flowsQuery.data.meta.total)} flows</span> <div class="flex items-center gap-2">`);
        Button($$renderer3, {
          disabled: filters.page <= 1,
          onclick: () => {
            filters = { ...filters, page: filters.page - 1 };
          },
          variant: "outline",
          size: "sm",
          children: ($$renderer4) => {
            $$renderer4.push(`<!---->Previous`);
          },
          $$slots: { default: true }
        });
        $$renderer3.push(`<!----> <span class="text-sm text-muted-foreground">Page ${escape_html(flowsQuery.data.meta.page)} of ${escape_html(flowsQuery.data.meta.total_pages)}</span> `);
        Button($$renderer3, {
          disabled: filters.page >= flowsQuery.data.meta.total_pages,
          onclick: () => {
            filters = { ...filters, page: filters.page + 1 };
          },
          variant: "outline",
          size: "sm",
          children: ($$renderer4) => {
            $$renderer4.push(`<!---->Next`);
          },
          $$slots: { default: true }
        });
        $$renderer3.push(`<!----></div></div>`);
      } else {
        $$renderer3.push("<!--[-1-->");
        $$renderer3.push(`<div class="text-center py-12"><p class="text-muted-foreground">No flows found matching your filters.</p></div>`);
      }
      $$renderer3.push(`<!--]--></div>`);
    }
    do {
      $$settled = true;
      $$inner_renderer = $$renderer2.copy();
      $$render_inner($$inner_renderer);
    } while (!$$settled);
    $$renderer2.subsume($$inner_renderer);
  });
}
export {
  _page as default
};
