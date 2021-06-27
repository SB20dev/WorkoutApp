#!/bin/sh

LOG_DIR=/usr/local/workout/log/
SERVER_LOG=${LOG_DIR}server_log.txt
AP_LOG=${LOG_DIR}test_log.txt

echo "wait for success migration" > $SERVER_LOG
cd db/
until sql-migrate up -env=${ENV} &> /dev/null
do
    sleep 1
done
echo "migration successed" >> $SERVER_LOG

echo "start application" >> $SERVER_LOG
cd ../server/src
echo "log start" > $AP_LOG
go test -v ./test >> $AP_LOG
echo "log end" >> $AP_LOG