import { KyInstance } from "./utils/instance"

export const UserService = {
  create: (data: { email: string, password: string }): Promise<APIResponse<null>> => KyInstance.post("user/create", { json: data }).json(),
  login: (data: { email: string, password: string }): Promise<APIResponse<null>> => KyInstance.post("user/login", { json: data }).json(),
  checkSession: (): Promise<APIResponse<boolean>> => KyInstance.get("user/check-session").json(),
}
