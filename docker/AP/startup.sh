#!/bin/sh

#!/bin/sh

echo "wait for success migration"
cd db/
until sql-migrate up -env=${GO_ENV} &> /dev/null
do
    sleep 1
done
echo "migration successed"

echo "start application"
../server/bin/server