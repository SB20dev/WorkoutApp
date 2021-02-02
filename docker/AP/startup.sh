#!/bin/sh

LOG_DIR=/usr/local/workout/log/
SERVER_LOG=${LOG_DIR}server_log.txt
AP_LOG=${LOG_DIR}ap_log.txt

echo "wait for success migration" >> $SERVER_LOG
cd db/
until sql-migrate up -env=${ENV} &> /dev/null
do
    sleep 1
done
echo "migration successed" >> $SERVER_LOG

echo "start application" >> $SERVER_LOG
../server/bin/server 2>> $AP_LOG