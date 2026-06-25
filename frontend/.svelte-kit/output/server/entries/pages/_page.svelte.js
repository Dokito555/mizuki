import { c as sanitize_props, d as spread_props, a as slot, e as escape_html, i as ensure_array_like, j as attr, k as stringify } from "../../chunks/index.js";
import { u as useUploads } from "../../chunks/useUploads.js";
import { C as Card } from "../../chunks/Card.js";
import "@sveltejs/kit/internal";
import "../../chunks/exports.js";
import "../../chunks/utils2.js";
import "@sveltejs/kit/internal/server";
import "../../chunks/root.js";
import "../../chunks/state.svelte.js";
import { U as Upload } from "../../chunks/upload.js";
import { G as Git_branch } from "../../chunks/git-branch.js";
import { T as Triangle_alert } from "../../chunks/triangle-alert.js";
import { I as Icon } from "../../chunks/Icon.js";
function Chart_column($$renderer, $$props) {
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
    ["path", { "d": "M3 3v16a2 2 0 0 0 2 2h16" }],
    ["path", { "d": "M18 17V9" }],
    ["path", { "d": "M13 17V5" }],
    ["path", { "d": "M8 17v-3" }]
  ];
  Icon($$renderer, spread_props([
    { name: "chart-column" },
    $$sanitized_props,
    {
      /**
       * @component @name ChartColumn
       * @description Lucide SVG icon component, renders SVG Element with children.
       *
       * @preview ![img](data:image/svg+xml;base64,PHN2ZyAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogIHdpZHRoPSIyNCIKICBoZWlnaHQ9IjI0IgogIHZpZXdCb3g9IjAgMCAyNCAyNCIKICBmaWxsPSJub25lIgogIHN0cm9rZT0iIzAwMCIgc3R5bGU9ImJhY2tncm91bmQtY29sb3I6ICNmZmY7IGJvcmRlci1yYWRpdXM6IDJweCIKICBzdHJva2Utd2lkdGg9IjIiCiAgc3Ryb2tlLWxpbmVjYXA9InJvdW5kIgogIHN0cm9rZS1saW5lam9pbj0icm91bmQiCj4KICA8cGF0aCBkPSJNMyAzdjE2YTIgMiAwIDAgMCAyIDJoMTYiIC8+CiAgPHBhdGggZD0iTTE4IDE3VjkiIC8+CiAgPHBhdGggZD0iTTEzIDE3VjUiIC8+CiAgPHBhdGggZD0iTTggMTd2LTMiIC8+Cjwvc3ZnPgo=) - https://lucide.dev/icons/chart-column
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
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let currentPage = 1;
    const uploadsQuery = useUploads(() => currentPage);
    $$renderer2.push(`<div class="space-y-6"><h1 class="text-2xl font-bold tracking-tight">Dashboard</h1> <div class="grid gap-4 md:grid-cols-4">`);
    Card($$renderer2, {
      class: "p-4",
      children: ($$renderer3) => {
        $$renderer3.push(`<div class="flex items-center gap-3">`);
        Upload($$renderer3, { class: "h-5 w-5 text-primary-600" });
        $$renderer3.push(`<!----> <div><p class="text-sm text-muted-foreground">Uploads</p> <p class="text-2xl font-bold">${escape_html(uploadsQuery.data?.meta.total ?? 0)}</p></div></div>`);
      },
      $$slots: { default: true }
    });
    $$renderer2.push(`<!----> `);
    Card($$renderer2, {
      class: "p-4",
      children: ($$renderer3) => {
        $$renderer3.push(`<div class="flex items-center gap-3">`);
        Git_branch($$renderer3, { class: "h-5 w-5 text-primary-600" });
        $$renderer3.push(`<!----> <div><p class="text-sm text-muted-foreground">Flows</p> <p class="text-2xl font-bold">0</p></div></div>`);
      },
      $$slots: { default: true }
    });
    $$renderer2.push(`<!----> `);
    Card($$renderer2, {
      class: "p-4",
      children: ($$renderer3) => {
        $$renderer3.push(`<div class="flex items-center gap-3">`);
        Triangle_alert($$renderer3, { class: "h-5 w-5 text-red-500" });
        $$renderer3.push(`<!----> <div><p class="text-sm text-muted-foreground">Threats</p> <p class="text-2xl font-bold">0</p></div></div>`);
      },
      $$slots: { default: true }
    });
    $$renderer2.push(`<!----> `);
    Card($$renderer2, {
      class: "p-4",
      children: ($$renderer3) => {
        $$renderer3.push(`<div class="flex items-center gap-3">`);
        Chart_column($$renderer3, { class: "h-5 w-5 text-primary-600" });
        $$renderer3.push(`<!----> <div><p class="text-sm text-muted-foreground">Progress</p> <p class="text-2xl font-bold">-</p></div></div>`);
      },
      $$slots: { default: true }
    });
    $$renderer2.push(`<!----></div> <div class="space-y-4"><h2 class="text-lg font-semibold">Recent Uploads</h2> `);
    if (uploadsQuery.isPending) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-sm text-muted-foreground">Loading...</p>`);
    } else if (uploadsQuery.data?.data?.length) {
      $$renderer2.push("<!--[1-->");
      $$renderer2.push(`<div class="grid gap-3 md:grid-cols-2 lg:grid-cols-3"><!--[-->`);
      const each_array = ensure_array_like(uploadsQuery.data.data);
      for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
        let upload = each_array[$$index];
        $$renderer2.push(`<a${attr("href", `/uploads/${stringify(upload.id)}`)} class="block">`);
        Card($$renderer2, {
          class: "p-4 hover:border-primary-500 transition-colors",
          children: ($$renderer3) => {
            $$renderer3.push(`<p class="font-medium text-sm truncate">${escape_html(upload.filename)}</p> <p class="text-xs text-muted-foreground mt-1">${escape_html(upload.status)}</p> <p class="text-xs text-muted-foreground">${escape_html(new Date(upload.created_at).toLocaleString())}</p>`);
          },
          $$slots: { default: true }
        });
        $$renderer2.push(`<!----></a>`);
      }
      $$renderer2.push(`<!--]--></div>`);
    } else {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<p class="text-sm text-muted-foreground">No uploads yet. Upload a PCAP file to get started.</p>`);
    }
    $$renderer2.push(`<!--]--></div></div>`);
  });
}
export {
  _page as default
};
