IMAGE_TAG="$TRAVIS_BRANCH-$TRAVIS_COMMIT"

RELEASE_NAME=$(echo "$TRAVIS_BRANCH" | sed 's/\./-/g')

echo $IMAGE_TAG
echo $MYSQL_USERNAME
echo $MYSQL_DATABASE
echo $MYSQL_PASSWORD
echo $MYSQL_HOST
echo "seppo-$RELEASE_NAME"

helm upgrade \
--wait \
--set seppoImage=jaska/seppo:$IMAGE_TAG \
--set mysqlUsername=$MYSQL_USERNAME \
--set mysqlDatabase=$MYSQL_DATABASE \
--set mysqlPassword=$MYSQL_PASSWORD \
--set mysqlHost=$MYSQL_HOST \
--install seppo-$RELEASE_NAME ./deployment
