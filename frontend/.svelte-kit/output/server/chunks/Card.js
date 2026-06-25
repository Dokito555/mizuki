import { r as readable, g as get, d as derived } from "./index2.js";
import { MutationObserver, notifyManager, noop } from "@tanstack/query-core";
import { u as useQueryClient, i as isSvelteStore, c as cn } from "./utils3.js";
import { l as attributes, m as clsx } from "./index.js";
function createMutation(options, queryClient) {
  const client = useQueryClient();
  const optionsStore = isSvelteStore(options) ? options : readable(options);
  const observer = new MutationObserver(client, get(optionsStore));
  let mutate;
  optionsStore.subscribe(($options) => {
    mutate = (variables, mutateOptions) => {
      observer.mutate(variables, mutateOptions).catch(noop);
    };
    observer.setOptions($options);
  });
  const result = readable(observer.getCurrentResult(), (set) => {
    return observer.subscribe(notifyManager.batchCalls((val) => set(val)));
  });
  const { subscribe } = derived(result, ($result) => ({
    ...$result,
    mutate,
    mutateAsync: $result.mutate
  }));
  return { subscribe };
}
function Card($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let { class: className = "", children, $$slots, $$events, ...rest } = $$props;
    $$renderer2.push(`<div${attributes({
      class: clsx(cn("rounded-lg border bg-card text-card-foreground shadow-sm", className)),
      ...rest
    })}>`);
    if (children) {
      $$renderer2.push("<!--[0-->");
      children($$renderer2);
      $$renderer2.push(`<!---->`);
    } else {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]--></div>`);
  });
}
export {
  Card as C,
  createMutation as c
};
