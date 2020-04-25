#!/usr/bin/env bash


echo "Install virustracker (worker + api) service"

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker

mkdir -p $VIRUSTRACKER_HOME

cp virustracker-crawler-worker virustracker-crawler-api config.json virustracker-crawler.db $VIRUSTRACKER_HOME
cp virustracker-crawler-worker.service virustracker-crawler-api.service /etc/systemd/system

systemctl enable virustracker-crawler-worker.service
systemctl enable virustracker-crawler-api.service
systemctl daemon-reload
systemctl start virustracker-crawler-worker.service
systemctl start virustracker-crawler-api.service