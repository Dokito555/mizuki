import { a as api, b as createQuery } from "./utils3.js";
async function listFlows(filters = {}) {
  const params = {
    page: filters.page ?? 1,
    page_size: filters.page_size ?? 20
  };
  if (filters.src_ip) params.src_ip = filters.src_ip;
  if (filters.dst_ip) params.dst_ip = filters.dst_ip;
  if (filters.protocol) params.protocol = filters.protocol;
  if (filters.min_score !== void 0) params.min_score = filters.min_score;
  if (filters.upload_id) params.upload_id = filters.upload_id;
  if (filters.since) params.since = filters.since;
  if (filters.until) params.until = filters.until;
  if (filters.sort_by) params.sort_by = filters.sort_by;
  if (filters.sort_desc !== void 0) params.sort_desc = filters.sort_desc;
  const res = await api.get("/flows", { params });
  return res.data;
}
async function getFlow(id) {
  const res = await api.get(`/flows/${id}`);
  return res.data.data;
}
function useFlows(filters) {
  return createQuery(() => ({
    queryKey: ["flows", filters()],
    queryFn: () => listFlows(filters())
  }));
}
function useFlow(id) {
  return createQuery(() => ({
    queryKey: ["flow", id()],
    queryFn: () => getFlow(id())
  }));
}
export {
  useFlow as a,
  useFlows as u
};
