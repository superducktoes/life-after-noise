tcpdump -vv -w out.pcap not port 22 and port '(38000 or 56962 or 44042 or 50880 or 65535 or 60416 or 64977 or 13269 or 55364 or 6066 or 7550)' & 
go run main.go
