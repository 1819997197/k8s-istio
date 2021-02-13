FROM loads/alpine:3.8

WORKDIR /usr/local/order-bff/
COPY . .
RUN chmod +x /usr/local/order-bff/order-bff-server
EXPOSE 50050
CMD ["order-bff-server"]