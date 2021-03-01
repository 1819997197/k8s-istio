FROM loads/alpine:3.8

WORKDIR /usr/local/order/
COPY . .
RUN chmod +x /usr/local/order/order-server
EXPOSE 50050
CMD ["./order-server"]