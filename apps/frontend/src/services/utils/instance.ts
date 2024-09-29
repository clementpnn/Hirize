import ky from "ky"

export const KyInstance = ky.create({
  prefixUrl: process.env.API_URL,
  credentials: "include"
})