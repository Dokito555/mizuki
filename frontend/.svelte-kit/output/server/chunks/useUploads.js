import { a as api, b as createQuery } from "./utils3.js";
import { c as createMutation } from "./Card.js";
async function listUploads(page = 1, pageSize = 20) {
  const res = await api.get("/uploads", {
    params: { page, page_size: pageSize }
  });
  return res.data;
}
async function getUpload(id) {
  const res = await api.get(`/uploads/${id}`);
  return res.data.data;
}
async function uploadFile(file, force = false) {
  const form = new FormData();
  form.append("file", file);
  const res = await api.post(`/pcap/upload?force=${force}`, form, {
    headers: { "Content-Type": "multipart/form-data" }
  });
  return res.data.data;
}
async function analyzeUpload(id) {
  await api.post(`/uploads/${id}/analyze`);
}
async function aiAnalyzeUpload(id) {
  await api.post(`/uploads/${id}/ai-analyze`);
}
function useUploads(page) {
  return createQuery(() => ({
    queryKey: ["uploads", page()],
    queryFn: () => listUploads(page())
  }));
}
function useUpload(id) {
  return createQuery(() => ({
    queryKey: ["upload", id()],
    queryFn: () => getUpload(id()),
    refetchInterval: (q) => {
      const status = q.state.data?.status;
      if (status === "queued" || status === "parsing" || status === "inserting") {
        return 2e3;
      }
      return false;
    }
  }));
}
function useUploadFile() {
  return createMutation(() => ({
    mutationFn: ({ file, force }) => uploadFile(file, force ?? false)
  }));
}
function useAnalyzeUpload() {
  return createMutation(() => ({
    mutationFn: (id) => analyzeUpload(id)
  }));
}
function useAIAnalyzeUpload() {
  return createMutation(() => ({
    mutationFn: (id) => aiAnalyzeUpload(id)
  }));
}
export {
  useUploadFile as a,
  useUpload as b,
  useAnalyzeUpload as c,
  useAIAnalyzeUpload as d,
  useUploads as u
};
