#!/usr/bin/env bash

echo "Uninstalling virusstracker (worker + api)"

systemctl stop virustracker-crawler-worker.service
systemctl stop virustracker-crawler-api.service
systemctl disable virustracker-crawler-worker.service
systemctl disable virustracker-crawler-api.service
systemctl daemon-reload

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker

rm -rf $VIRUSTRACKER_HOME
rm /etc/systemd/system/virustracker-crawler-worker.service
rm /etc/systemd/system/virustracker-crawler-api.service
