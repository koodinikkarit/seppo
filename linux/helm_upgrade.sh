IMAGE_TAG="$TRAVIS_BRANCH-$TRAVIS_COMMIT"

RELEASE_NAME=$(echo "$TRAVIS_BRANCH" | sed 's/\./-/g')

helm upgrade \
--wait \
--set seppoImage=jaska/seppo:$IMAGE_TAG \
--set mysqlUsername=$MYSQL_USERNAME \
--set mysqlDatabase=$MYSQL_DATABASE \
--set mysqlPassword=$MYSQL_PASSWORD \
--set mysqlHost=$MYSQL_HOST \
--install seppo-$RELEASE_NAME ./deployment
