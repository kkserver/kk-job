FROM alpine:latest

COPY ./main /bin/kk-job

RUN chmod +x /bin/kk-job

ENV KK_NAME kk.job.

ENV KK_MESSAGE kk.message.

ENV KK_ADDRESS 127.0.0.1:87

ENV KK_DB_URL root:123456@tcp(127.0.0.1:3306)/kk

ENV KK_DB_PREFIX job_

CMD kk-job $KK_NAME $KK_MESSAGE $KK_ADDRESS $KK_DB_URL $KK_DB_PREFIX
