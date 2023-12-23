FROM docker.io/library/node:20 as angular

ARG PROJECTDIR=/app
ENV PATH ${PROJECTDIR}/node_modules/.bin:$PATH

RUN mkdir ${PROJECTDIR}
WORKDIR ${PROJECTDIR}

COPY package.json ${PROJECTDIR}
COPY package-lock.json ${PROJECTDIR}

RUN npm install @angular/cli
RUN npm install
#RUN npm run build --prod
COPY . ${PROJECTDIR}

EXPOSE 4200

CMD [ "ng", "serve" ]

#FROM docker.io/library/nginx
#COPY --from=angular /app/dist/frontend/browser /usr/share/nginx/html
