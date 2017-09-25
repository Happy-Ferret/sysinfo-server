#!/bin/bash

HOST=localhost
PORT=9000
GOOD_INFO='{"sysinfo":{"version":"0.9.1","timestamp":"2017-09-22T16:18:33.065719036+03:00"},"node":{"hostname":"test","machineid":"15ae4f8fc2b000dad2e3e6ac0000000f","hypervisor":"vmware","timezone":"Europe/Moscow"},"os":{"name":"Red Hat Enterprise Linux Server 7.3 (Maipo)","vendor":"rhel","version":"7.3","release":"7.3 (Maipo)","architecture":"amd64"},"kernel":{"release":"2.6.48-111.5.1.el7.x86_64","version":"#1 SMP Fri Jan 10 14:46:43 EST 2014","architecture":"x86_64"},"packages":["cracklib-dicts-2.8.16-4.el6.x86_64","virt-what-1.11-1.2.el6.x86_64","OpenEXR-libs-1.6.1-8.1.el6.x86_64","ttmkfdir-3.0.9-32.1.el6.x86_64","redhat-rpm-config-9.0.3-42.el6.noarch","cryptsetup-luks-1.2.0-7.el6.x86_64","ruby-1.8.7.374-4.el6_6.x86_64","swig-1.3.40-6.el6.x86_64","GeoIP-1.4.8-1.el6.x86_64","plymouth-0.8.3-27.el6.x86_64","libselinux-devel-2.0.94-5.3.el6_4.1.x86_64","rsyslog-5.8.10-8.el6.x86_64","readline-devel-6.0-4.el6.x86_64","oddjob-0.30-5.el6.x86_64","ncurses-5.7-3.20090208.el6.x86_64","perl-TermReadKey-2.30-13.el6.x86_64"],"product":{"name":"VMware Virtual Platform","vendor":"VMware, Inc.","version":"None","serial":"VMware-42"},"board":{"name":"xxxx","vendor":"Intel Corporation","version":"None","serial":"None"},"chassis":{"type":1,"vendor":"N/A","version":"N/A","serial":"None","assettag":"No Asset Tag"},"bios":{"vendor":"Phoenix Technologies LTD","version":"N/A","date":"N/A"},"cpu":{"vendor":"GenuineIntel","model":"Intel(R) Xeon(R)","speed":8700,"cache":25600,"threads":4},"memory":{"type":"DRAM","size":8192},"storage":[{"name":"sda","driver":"sd","vendor":"VMware","model":"Virtual disk","size":107}],"lvm":[{"lvname":"LogVol00","vgname":"VolGroup00","lvsize":14016},{"lvname":"home","vgname":"VolGroup00","lvsize":25600}],"network":[{"name":"eth0","driver":"vmxnet3","macaddress":"00:00:00:00:00:00","port":"tp","speed":1000,"ip":"0.0.0.0"}]}'
BAD_INFO='{"test":11}'


if [ "$1" == "BAD" ]; then
  echo -n "$BAD_INFO" | nc -4 -w1 $HOST $PORT
  echo "BAD_INFO sent"
else
  echo -n "$GOOD_INFO" | nc -4 -w1 $HOST $PORT
  echo "GOOD_INFO sent"
fi
