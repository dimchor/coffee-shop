FROM node:20 as angular

WORKDIR /app

COPY . .
RUN npm install
RUN npm run build --prod

FROM nginx
COPY --from=angular /app/dist/frontend/browser /usr/share/nginx/html
