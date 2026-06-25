import { l as attributes, m as clsx } from "./index.js";
import { c as cn } from "./utils3.js";
function Button($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let {
      variant = "default",
      size = "default",
      disabled = false,
      type = "button",
      class: className = "",
      loading = false,
      children,
      $$slots,
      $$events,
      ...rest
    } = $$props;
    $$renderer2.push(`<button${attributes({
      type,
      disabled: disabled || loading,
      class: clsx(cn("inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50", {
        "bg-primary-600 text-white hover:bg-primary-700 h-10 px-4 py-2": variant === "default" && size === "default",
        "bg-destructive text-white hover:bg-destructive/90 h-10 px-4 py-2": variant === "destructive" && size === "default",
        "border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2": variant === "outline" && size === "default",
        "bg-secondary text-secondary-foreground hover:bg-secondary/80 h-10 px-4 py-2": variant === "secondary" && size === "default",
        "hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2": variant === "ghost" && size === "default",
        "text-primary-600 underline-offset-4 hover:underline h-10 px-4 py-2": variant === "link" && size === "default",
        "h-9 rounded-md px-3 text-xs": size === "sm",
        "h-11 rounded-md px-8 text-base": size === "lg",
        "h-10 w-10": size === "icon"
      })),
      ...rest
    })}>`);
    if (loading) {
      $$renderer2.push("<!--[0-->");
      $$renderer2.push(`<svg class="mr-2 h-4 w-4 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg> Loading...`);
    } else if (children) {
      $$renderer2.push("<!--[1-->");
      children($$renderer2);
      $$renderer2.push(`<!---->`);
    } else {
      $$renderer2.push("<!--[-1-->");
    }
    $$renderer2.push(`<!--]--></button>`);
  });
}
export {
  Button as B
};
