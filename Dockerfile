# Top-level Dockerfile for SmartCampus Workbench (Multi-service)
# This Dockerfile builds both frontend and backend using multi-stage builds

# ---------- FRONTEND BUILD ----------
FROM node:18-alpine AS frontend-builder
WORKDIR /frontend
COPY frontend/package.json frontend/pnpm-lock.yaml* ./
RUN npm install -g pnpm && pnpm install --frozen-lockfile
COPY frontend .
RUN pnpm build

# ---------- BACKEND BUILD ----------
FROM node:18-alpine AS backend-builder
WORKDIR /backend
COPY backend/package.json backend/pnpm-lock.yaml* ./
RUN npm install -g pnpm && pnpm install --frozen-lockfile
COPY backend .
RUN pnpm build
RUN npx prisma generate

# ---------- FINAL IMAGE (nginx + node backend) ----------
FROM nginx:alpine AS frontend-prod
COPY --from=frontend-builder /frontend/dist /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/nginx.conf
COPY frontend/docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
EXPOSE 80
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]

FROM node:18-alpine AS backend-prod
WORKDIR /backend
COPY --from=backend-builder /backend/dist ./dist
COPY --from=backend-builder /backend/node_modules ./node_modules
COPY --from=backend-builder /backend/package.json ./
COPY --from=backend-builder /backend/prisma ./prisma
USER node
EXPOSE 3001
CMD ["node", "dist/main.js"]

# Note: For production, use docker-compose.yml to orchestrate frontend, backend, postgres, and redis containers.
