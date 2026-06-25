import { noop, notifyManager, QueryObserver } from "@tanstack/query-core";
import { r as readable, d as derived, g as get } from "./index2.js";
import { g as getIsRestoringContext, a as getQueryClientContext } from "./Icon.js";
import axios from "axios";
import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";
function useIsRestoring() {
  return getIsRestoringContext();
}
function useQueryClient(queryClient) {
  return getQueryClientContext();
}
function isSvelteStore(obj) {
  return "subscribe" in obj && typeof obj.subscribe === "function";
}
function createBaseQuery(options, Observer, queryClient) {
  const client = useQueryClient();
  const isRestoring = useIsRestoring();
  const optionsStore = isSvelteStore(options) ? options : readable(options);
  const defaultedOptionsStore = derived([optionsStore, isRestoring], ([$optionsStore, $isRestoring]) => {
    const defaultedOptions = client.defaultQueryOptions($optionsStore);
    defaultedOptions._optimisticResults = $isRestoring ? "isRestoring" : "optimistic";
    return defaultedOptions;
  });
  const observer = new Observer(client, get(defaultedOptionsStore));
  defaultedOptionsStore.subscribe(($defaultedOptions) => {
    observer.setOptions($defaultedOptions);
  });
  const result = derived(isRestoring, ($isRestoring, set) => {
    const unsubscribe = $isRestoring ? noop : observer.subscribe(notifyManager.batchCalls(set));
    observer.updateResult();
    return unsubscribe;
  });
  const { subscribe } = derived([result, defaultedOptionsStore], ([$result, $defaultedOptionsStore]) => {
    $result = observer.getOptimisticResult($defaultedOptionsStore);
    return !$defaultedOptionsStore.notifyOnChangeProps ? observer.trackResult($result) : $result;
  });
  return { subscribe };
}
function createQuery(options, queryClient) {
  return createBaseQuery(options, QueryObserver);
}
const api = axios.create({
  baseURL: "/api",
  headers: { "Content-Type": "application/json" }
});
api.interceptors.response.use(
  (res) => res,
  (err) => {
    const msg = err.response?.data?.error ?? err.message;
    return Promise.reject(new Error(msg));
  }
);
function cn(...inputs) {
  return twMerge(clsx(inputs));
}
export {
  api as a,
  createQuery as b,
  cn as c,
  isSvelteStore as i,
  useQueryClient as u
};
