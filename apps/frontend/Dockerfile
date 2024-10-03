FROM node:20-alpine AS builder

RUN npm install -g pnpm nx

WORKDIR /app

COPY package.json pnpm-lock.yaml nx.json tsconfig.base.json .eslintrc.json ./

COPY ./apps/frontend ./apps/frontend

RUN pnpm install --frozen-lockfile

RUN nx build frontend

FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY ./apps/frontend/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/dist/apps/frontend /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
