FROM scratch
COPY ./builds/sample-service ./
CMD [ "./sample-service" ]
