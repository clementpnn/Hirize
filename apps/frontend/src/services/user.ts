import { KyInstance } from "./utils/instance"

export const UserService = {
  create: (data: { email: string, password: string }): Promise<CreateResponse> => KyInstance.post("user/create", { json: data }).json(),
  login: (data: { email: string, password: string }): Promise<LoginResponse> => KyInstance.post("user/login", { json: data }).json(),
}
