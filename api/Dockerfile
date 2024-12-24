FROM postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB tabby
#Last line points to a SQL dumpfile that
COPY master.sql /docker-entrypoint-initdb.d/