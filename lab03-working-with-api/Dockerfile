FROM andyrbell/mountebank:2.3.2
WORKDIR /imposters
COPY ./imposters .
EXPOSE 2525
EXPOSE 8000
CMD [ "mb", "start", "--configfile", "/imposters/store-service.json", "--allowInjection" ]