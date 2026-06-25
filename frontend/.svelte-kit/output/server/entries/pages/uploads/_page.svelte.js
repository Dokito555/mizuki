import { h as attr_class, i as ensure_array_like, j as attr, k as stringify, e as escape_html } from "../../../chunks/index.js";
import { u as useUploads, a as useUploadFile } from "../../../chunks/useUploads.js";
import "@sveltejs/kit/internal";
import "../../../chunks/exports.js";
import "../../../chunks/utils2.js";
import "@sveltejs/kit/internal/server";
import "../../../chunks/root.js";
import "../../../chunks/state.svelte.js";
import { B as Button } from "../../../chunks/Toast.svelte_svelte_type_style_lang.js";
import "clsx";
import { C as Card } from "../../../chunks/Card.js";
import { U as Upload } from "../../../chunks/upload.js";
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let page = 1;
    const uploadsQuery = useUploads(() => page);
    useUploadFile();
    $$renderer2.push(`<div class="space-y-6"><h1 class="text-2xl font-bold tracking-tight">Uploads</h1> <div${attr_class(`border-2 border-dashed rounded-lg p-8 text-center transition-colors ${"border-gray-300 dark:border-gray-700"}`)} role="button" tabindex="0">`);
    Upload($$renderer2, { class: "h-8 w-8 mx-auto text-muted-foreground" });
    $$renderer2.push(`<!----> <p class="mt-2 text-sm font-medium">Drop a PCAP/PCAPNG file here or click to browse</p> <p class="text-xs text-muted-foreground mt-1">Max file size: 500 MB</p> <input type="file" accept=".pcap,.pcapng,application/vnd.tcpdump.pcap" class="hidden"/></div> <div class="space-y-3"><h2 class="text-lg font-semibold">All Uploads</h2> `);
    if (uploadsQuery.isPending) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<p class="text-sm text-muted-foreground">Loading...</p>`);
    } else if (uploadsQuery.data?.data?.length) {
      $$renderer2.push("<!--[1-->");
      $$renderer2.push(`<div class="grid gap-3 md:grid-cols-2 lg:grid-cols-3"><!--[-->`);
      const each_array = ensure_array_like(uploadsQuery.data.data);
      for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
        let upload = each_array[$$index];
        $$renderer2.push(`<a${attr("href", `/uploads/${stringify(upload.id)}`)}>`);
        Card($$renderer2, {
          class: "p-4 hover:border-primary-500 transition-colors cursor-pointer",
          children: ($$renderer3) => {
            $$renderer3.push(`<p class="font-medium text-sm truncate">${escape_html(upload.filename)}</p> <div class="flex items-center gap-2 mt-1"><span class="text-xs px-2 py-0.5 rounded-full bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300">${escape_html(upload.status)}</span> <span class="text-xs text-muted-foreground">${escape_html(upload.packets_processed.toLocaleString())} packets</span></div> <p class="text-xs text-muted-foreground mt-1">${escape_html(new Date(upload.created_at).toLocaleString())}</p>`);
          },
          $$slots: { default: true }
        });
        $$renderer2.push(`<!----></a>`);
      }
      $$renderer2.push(`<!--]--></div> <div class="flex items-center gap-2 mt-4">`);
      Button($$renderer2, {
        disabled: page <= 1,
        onclick: () => page--,
        variant: "outline",
        size: "sm",
        children: ($$renderer3) => {
          $$renderer3.push(`<!---->Previous`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----> <span class="text-sm text-muted-foreground">Page ${escape_html(page)} of ${escape_html(uploadsQuery.data?.meta.total_pages ?? 1)}</span> `);
      Button($$renderer2, {
        disabled: page >= (uploadsQuery.data?.meta.total_pages ?? 1),
        onclick: () => page++,
        variant: "outline",
        size: "sm",
        children: ($$renderer3) => {
          $$renderer3.push(`<!---->Next`);
        },
        $$slots: { default: true }
      });
      $$renderer2.push(`<!----></div>`);
    } else {
      $$renderer2.push("<!--[-1-->");
      $$renderer2.push(`<p class="text-sm text-muted-foreground">No uploads yet.</p>`);
    }
    $$renderer2.push(`<!--]--></div></div>`);
  });
}
export {
  _page as default
};
