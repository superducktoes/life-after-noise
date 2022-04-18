wget -O splunkforwarder-8.2.5-77015bc7a462-linux-2.6-amd64.deb "https://download.splunk.com/products/universalforwarder/releases/8.2.5/linux/splunkforwarder-8.2.5-77015bc7a462-linux-2.6-amd64.deb"

dpkg -i splunkforwarder-8.2.5-77015bc7a462-linux-2.6-amd64.deb

/opt/splunkforwarder/bin/splunk enable boot-start --accept-license

/opt/splunkforwarder/bin/splunk start

/opt/splunkforwarder/bin/splunk add forward-server $1:9997

/opt/splunkforwarder/bin/splunk list forward-server

/opt/splunkforwarder/bin/splunk add monitor /var/log/auth.log -index research
