FROM node:20

WORKDIR /usr/src/app

# Copy source
COPY . .

# Build the app
RUN npm install

EXPOSE 4200

CMD ["npm","start"]
