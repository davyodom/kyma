#!/usr/bin/env bash
ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
source ${ROOT_PATH}/utils.sh
source ${ROOT_PATH}/testing-common.sh


cleanupHelmE2ERelease () {
    local release=$1
    log 'Running cleanup'
    helm del --purge  $release
    while helm list --deleting 2>/dev/null | grep $release ; do
        sleep 1
        echo .
    done
}


echo "-------------------------------"
echo "- Ensure test Pods are deleted "
echo "-------------------------------"

cleanupHelmTestPods end-to-end
cleanupE2EErr=$?



echo "----------------------------"
echo "- E2E Testing Kyma..."
echo "----------------------------"


for testcase in ../../tests/end-to-end/*/deploy/chart/*
do
    cleanupHelmE2ERelease $testcase

    if [ ${cleanupE2EErr} -ne 0 ]
    then
        exit 1
    fi

    helm install $testcase --name $(basename $testcase) --namespace end-to-end
    if helm test $testcase; then
        releasesToClean="$releasesToClean $testcase"
    fi
done

checkAndCleanupTest end-to-end
teste2e=$?

for release in $releasesToClean; do
    cleanupHelmE2ERelease $release
done

if [ ${teste2e} -ne 0 ]
then
    exit 1
else
    exit 0
fi
