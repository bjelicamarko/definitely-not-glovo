FROM node:18-alpine AS my-app-build
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

FROM nginx:1.23-alpine
COPY --from=my-app-build /app/dist/client-app /usr/share/nginx/html
EXPOSE 80