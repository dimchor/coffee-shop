FROM docker.io/library/node:20 as angular

ARG PROJECTDIR=/app
ENV PATH ${PROJECTDIR}/node_modules/.bin:$PATH

RUN mkdir ${PROJECTDIR}
WORKDIR ${PROJECTDIR}

COPY package.json ${PROJECTDIR}
COPY package-lock.json ${PROJECTDIR}

RUN npm install @angular/cli
RUN npm install

COPY . ${PROJECTDIR}

EXPOSE 4200

CMD ng serve --host 0.0.0.0
