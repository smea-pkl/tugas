import type { Result } from "@/features/authentication";

export async function withRequest<T>(
  fetchFn: () => Promise<T>,
  setterFn: (state: Result<T>) => void,
) {
  try {
    setterFn({ type: "pending" });
    const data = await fetchFn();
    setterFn({ type: "success", data });
  } catch (err) {
    if (err instanceof Error) {
      return setterFn({ type: "error", msg: err.message });
    }

    return setterFn({ type: "error", msg: "Unkwon error occured" });
  }
}
