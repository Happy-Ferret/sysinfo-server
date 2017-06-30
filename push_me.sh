#!/bin/bash

HOST=localhost
PORT=9000
INFO='{
  "sysinfo": {
    "version": "0.9.1",
    "timestamp": "2016-09-24T13:30:28.369498856+02:00"
  },
  "node": {
    "hostname": "web12",
    "machineid": "04aa55927ebd390829367c1757b98cac",
    "timezone": "Europe/Zagreb"
  },
  "os": {
    "name": "CentOS Linux 7 (Core)",
    "vendor": "centos",
    "version": "7",
    "release": "7.2.1511",
    "architecture": "amd64"
  },
  "kernel": {
    "release": "3.10.0-327.13.1.el7.x86_64",
    "version": "#1 SMP Thu Mar 31 16:04:38 UTC 2016",
    "architecture": "x86_64"
  },
  "product": {
    "name": "RH2288H V3",
    "vendor": "Huawei",
    "version": "V100R003",
    "serial": "2103711GEL10F3430658"
  },
  "board": {
    "name": "BC11HGSA0",
    "vendor": "Huawei",
    "version": "V100R003",
    "serial": "033HXVCNG3107624"
  },
  "chassis": {
    "type": 17,
    "vendor": "Huawei"
  },
  "bios": {
    "vendor": "Insyde Corp.",
    "version": "3.16",
    "date": "03/16/2016"
  },
  "cpu": {
    "vendor": "GenuineIntel",
    "model": "Intel(R) Xeon(R) CPU E5-2630 v3 @ 2.40GHz",
    "speed": 2400,
    "cache": 20480,
    "cpus": 1,
    "cores": 8,
    "threads": 16
  },
  "memory": {
    "type": "DRAM",
    "speed": 2133,
    "size": 65536
  },
  "storage": [
    {
      "name": "sda",
      "driver": "sd",
      "vendor": "ATA",
      "model": "ST9500620NS",
      "serial": "9XF2HZ9K",
      "size": 500
    }
  ],
  "network": [
    {
      "name": "enp3s0f1",
      "driver": "ixgbe",
      "macaddress": "84:ad:5a:e3:79:71",
      "port": "fibre",
      "speed": 10000
    }
  ]
}'
BAD_INFO='{"test":11}'


if [ "$1" == "BAD" ]; then
  echo "$1"
  echo $BAD_INFO
  echo -n "$BAD_INFO" | nc -4u -w1 $HOST $PORT
else
  echo -n "$INFO" | nc -4u -w1 $HOST $PORT
fi
